package page_func_spec

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	getSystemTimesProc   = kernel32.NewProc("GetSystemTimes")
	psapi                = syscall.NewLazyDLL("psapi.dll")
	getProcessMemoryInfo = psapi.NewProc("GetProcessMemoryInfo")
)

type Filetime struct {
	LowDateTime  uint32
	HighDateTime uint32
}

func getCPUUsage() string {
	var idleTime, kernelTime, userTime Filetime

	ret, _, err := getSystemTimesProc.Call(
		uintptr(unsafe.Pointer(&idleTime)),
		uintptr(unsafe.Pointer(&kernelTime)),
		uintptr(unsafe.Pointer(&userTime)),
	)
	if ret == 0 {
		return fmt.Sprintf("Failed to get system times: %v", err)
	}

	idle := uint64(idleTime.HighDateTime)<<32 | uint64(idleTime.LowDateTime)
	kernel := uint64(kernelTime.HighDateTime)<<32 | uint64(kernelTime.LowDateTime)
	user := uint64(userTime.HighDateTime)<<32 | uint64(userTime.LowDateTime)

	totalTime := kernel + user
	cpuUsage := (1.0 - float64(idle)/float64(totalTime)) * 100.0

	return fmt.Sprintf("%.2f%%", cpuUsage)
}

func getMemoryUsage() string {
	var memCounters struct {
		cb                         uint32
		PageFaultCount             uint32
		PeakWorkingSetSize         uintptr
		WorkingSetSize             uintptr
		QuotaPeakPagedPoolUsage    uintptr
		QuotaPagedPoolUsage        uintptr
		QuotaPeakNonPagedPoolUsage uintptr
		QuotaNonPagedPoolUsage     uintptr
		PagefileUsage              uintptr
		PeakPagefileUsage          uintptr
	}

	handle, _ := syscall.GetCurrentProcess()

	ret, _, err := getProcessMemoryInfo.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(&memCounters)),
		uintptr(unsafe.Sizeof(memCounters)),
	)
	if ret == 0 {
		return fmt.Sprintf("Failed to get memory info: %v", err)
	}

	memoryUsageMB := float64(memCounters.WorkingSetSize) / (1024 * 1024)
	return fmt.Sprintf("%.2f MB", memoryUsageMB)
}

func Get_all_data_now() string {
	cpuInfo := getCPUUsage()
	memoryInfo := getMemoryUsage()

	return fmt.Sprintf("%s\n%s", cpuInfo, memoryInfo)
}
