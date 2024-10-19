#include <windows.h>
#include <psapi.h>
#include <iostream>
#include <fstream>
#include <string>

void GetCPUUsage(std::ofstream& logFile) {
    FILETIME idleTime, kernelTime, userTime;
    if (GetSystemTimes(&idleTime, &kernelTime, &userTime)) {
        ULARGE_INTEGER idle, kernel, user;
        idle.HighPart = idleTime.dwHighDateTime;
        idle.LowPart = idleTime.dwLowDateTime;
        kernel.HighPart = kernelTime.dwHighDateTime;
        kernel.LowPart = kernelTime.dwLowDateTime;
        user.HighPart = userTime.dwHighDateTime;
        user.LowPart = userTime.dwLowDateTime;

        ULONGLONG totalTime = (kernel.QuadPart + user.QuadPart);
        ULONGLONG totalIdleTime = idle.QuadPart;

        double cpuUsage = (1.0 - (static_cast<double>(totalIdleTime) / totalTime)) * 100.0;
        logFile << cpuUsage << "%" << std::endl;
    } else {
        logFile << "Failed to get system times: " << GetLastError() << std::endl;
    }
}

void GetMemoryUsage(std::ofstream& logFile) {
    PROCESS_MEMORY_COUNTERS pmc;
    if (GetProcessMemoryInfo(GetCurrentProcess(), &pmc, sizeof(pmc))) {
        double memoryUsage = static_cast<double>(pmc.WorkingSetSize) / (1024 * 1024);
        logFile << memoryUsage << " MB" << std::endl;
    } else {
        logFile << "Failed to get memory info: " << GetLastError() << std::endl;
    }
}

int main() {
    std::ofstream logFile("resource_info.log");
    if (!logFile.is_open()) {
        std::cerr << "Failed to open resource_info.log for writing." << std::endl;
        return 1;
    }

    GetCPUUsage(logFile);
    GetMemoryUsage(logFile);

    logFile.close();
    return 0;
}
