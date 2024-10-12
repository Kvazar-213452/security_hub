#include <windows.h>
#include <string>
#include <sstream>

extern "C" {
    __declspec(dllexport) const char* GetSystemMemory() {
        static std::string result;
        MEMORYSTATUSEX memInfo;
        memInfo.dwLength = sizeof(MEMORYSTATUSEX);
        GlobalMemoryStatusEx(&memInfo);

        std::ostringstream oss;
        oss << "Total RAM: " << memInfo.ullTotalPhys / (1024 * 1024) << " MB\n"
            << "Available RAM: " << memInfo.ullAvailPhys / (1024 * 1024) << " MB\n";
        result = oss.str();
        return result.c_str();
    }

    __declspec(dllexport) const char* GetProcessorInfo() {
        static std::string result;
        SYSTEM_INFO sysInfo;
        GetSystemInfo(&sysInfo);

        std::ostringstream oss;
        oss << "Number of Processors: " << sysInfo.dwNumberOfProcessors << "\n"
            << "Processor Architecture: " << sysInfo.wProcessorArchitecture << "\n";
        result = oss.str();
        return result.c_str();
    }
}
