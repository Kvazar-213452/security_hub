package page_func

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

type ProcessInfo struct {
	Name    string  `json:"name"`
	PID     int32   `json:"pid"`
	CPU     float64 `json:"cpu"`
	Memory  float64 `json:"memory"`
	NetSent uint64  `json:"net_sent"`
	NetRecv uint64  `json:"net_recv"`
}

type Summary struct {
	Processes    []ProcessInfo `json:"processes"`
	TotalCPU     float64       `json:"total_cpu"`
	TotalMem     float64       `json:"total_memory"`
	TotalNetSent uint64        `json:"total_net_sent"`
	TotalNetRecv uint64        `json:"total_net_recv"`
}

func Get_process_info() string {
	var summary Summary

	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Error getting process list:", err)
		return ""
	}

	netIO, err := net.IOCounters(false)
	if err != nil {
		fmt.Println("Error getting network info:", err)
		return ""
	}

	for _, proc := range processes {
		name, err := proc.Name()
		if err != nil {
			continue
		}

		if strings.HasSuffix(strings.ToLower(name), ".exe") {
			cpuPercent, err := proc.CPUPercent()
			if err != nil || cpuPercent <= 0.1 {
				continue
			}

			memInfo, err := proc.MemoryInfo()
			var memUsage float64
			if err == nil {
				memUsage = float64(memInfo.RSS) / 1024 / 1024
			}

			procInfo := ProcessInfo{
				Name:    name,
				PID:     proc.Pid,
				CPU:     cpuPercent,
				Memory:  memUsage,
				NetSent: netIO[0].BytesSent,
				NetRecv: netIO[0].BytesRecv,
			}

			summary.Processes = append(summary.Processes, procInfo)
			summary.TotalCPU += cpuPercent
			summary.TotalMem += memUsage
			summary.TotalNetSent += netIO[0].BytesSent
			summary.TotalNetRecv += netIO[0].BytesRecv
		}
	}

	data, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return ""
	}

	return string(data)
}
