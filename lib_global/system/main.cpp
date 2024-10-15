#include <windows.h>
#include <string>
#include <sstream>
#include <sysinfoapi.h>
#include <Lmcons.h>

extern "C" {
    __declspec(dllexport) const char* GetSystemMemory() {
        static std::string result;
        MEMORYSTATUSEX memInfo;
        memInfo.dwLength = sizeof(MEMORYSTATUSEX);
        GlobalMemoryStatusEx(&memInfo);

        std::ostringstream oss;
        oss << memInfo.ullTotalPhys / (1024 * 1024) << " MB\n"
            << memInfo.ullAvailPhys / (1024 * 1024) << " MB\n";
        result = oss.str();
        return result.c_str();
    }

    __declspec(dllexport) const char* GetProcessorInfo() {
        static std::string result;
        SYSTEM_INFO sysInfo;
        GetSystemInfo(&sysInfo);

        std::ostringstream oss;
        oss << sysInfo.dwNumberOfProcessors << "\n"
            << sysInfo.wProcessorArchitecture << "\n";
        result = oss.str();
        return result.c_str();
    }

    __declspec(dllexport) const char* GetOSVersion() {
        static std::string result;
        OSVERSIONINFO osvi;
        ZeroMemory(&osvi, sizeof(OSVERSIONINFO));
        osvi.dwOSVersionInfoSize = sizeof(OSVERSIONINFO);
        GetVersionEx(&osvi);

        std::ostringstream oss;
        oss << osvi.dwMajorVersion << "." << osvi.dwMinorVersion << "\n"
            << osvi.dwBuildNumber << "\n";
        result = oss.str();
        return result.c_str();
    }

    __declspec(dllexport) const char* GetComputerNameCustom() {
        static std::string result;
        char computerName[MAX_COMPUTERNAME_LENGTH + 1];
        DWORD size = sizeof(computerName);
        GetComputerNameA(computerName, &size);

        result = std::string(computerName) + "\n";
        return result.c_str();
    }

    __declspec(dllexport) const char* GetUserNameCustom() {
        static std::string result;
        char userName[UNLEN + 1];
        DWORD size = sizeof(userName);
        GetUserNameA(userName, &size);

        result = std::string(userName) + "\n";
        return result.c_str();
    }

    __declspec(dllexport) const char* GetSystemUptime() {
        static std::string result;
        DWORD uptime = GetTickCount64() / 1000;
        DWORD days = uptime / (24 * 3600);
        uptime %= 24 * 3600;
        DWORD hours = uptime / 3600;
        uptime %= 3600;
        DWORD minutes = uptime / 60;
        DWORD seconds = uptime % 60;

        std::ostringstream oss;
        oss << days << " days, " << hours << " hours, " 
            << minutes << " minutes, " << seconds << " seconds\n";
        result = oss.str();
        return result.c_str();
    }
}
