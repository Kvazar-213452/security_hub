#include <windows.h>
#include <stdio.h>
#include <psapi.h>
#include <iphlpapi.h>
#include <time.h>

#pragma comment(lib, "iphlpapi.lib")
#pragma comment(lib, "psapi.lib")

void getSystemInfo() {
    FILE *file = fopen("data/file_2.txt", "w");
    if (file == NULL) {
        printf("Не вдалося відкрити файл для запису.\n");
        return;
    }

    fprintf(file, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n");
    fprintf(file, "<SystemInfo>\n");

    SYSTEM_INFO si;
    GetSystemInfo(&si);

    fprintf(file, "  <Architecture>");
    switch (si.wProcessorArchitecture) {
        case PROCESSOR_ARCHITECTURE_AMD64:
            fprintf(file, "x64 (AMD or Intel)");
            break;
        case PROCESSOR_ARCHITECTURE_ARM:
            fprintf(file, "ARM");
            break;
        case PROCESSOR_ARCHITECTURE_IA64:
            fprintf(file, "IA64 (Intel Itanium)");
            break;
        default:
            fprintf(file, "невідома");
            break;
    }
    fprintf(file, "</Architecture>\n");

    fprintf(file, "  <ProcessorCount>%u</ProcessorCount>\n", si.dwNumberOfProcessors);

    OSVERSIONINFOEX osvi;
    osvi.dwOSVersionInfoSize = sizeof(OSVERSIONINFOEX);
    if (GetVersionEx((OSVERSIONINFO*)&osvi)) {
        fprintf(file, "  <OS>\n");
        fprintf(file, "    <Name>");
        if (osvi.dwMajorVersion == 10) {
            fprintf(file, "Windows 10");
        } else if (osvi.dwMajorVersion == 6) {
            if (osvi.dwMinorVersion == 3) {
                fprintf(file, "Windows 8.1");
            } else if (osvi.dwMinorVersion == 2) {
                fprintf(file, "Windows 8");
            } else if (osvi.dwMinorVersion == 1) {
                fprintf(file, "Windows 7");
            } else if (osvi.dwMinorVersion == 0) {
                fprintf(file, "Windows Vista");
            }
        } else if (osvi.dwMajorVersion == 5) {
            if (osvi.dwMinorVersion == 1) {
                fprintf(file, "Windows XP");
            } else if (osvi.dwMinorVersion == 0) {
                fprintf(file, "Windows 2000");
            }
        }
        fprintf(file, "</Name>\n");
        fprintf(file, "    <Version>%d.%d</Version>\n", osvi.dwMajorVersion, osvi.dwMinorVersion);
        fprintf(file, "  </OS>\n");
    } else {
        fprintf(file, "  <OS>\n    <Error>Не вдалося отримати версію операційної системи.</Error>\n  </OS>\n");
    }

    MEMORYSTATUSEX statex;
    statex.dwLength = sizeof(statex);
    if (GlobalMemoryStatusEx(&statex)) {
        fprintf(file, "  <Memory>\n");
        fprintf(file, "    <FreeMemory>%llu</FreeMemory>\n", statex.ullAvailPhys / (1024 * 1024));
        fprintf(file, "    <TotalMemory>%llu</TotalMemory>\n", statex.ullTotalPhys / (1024 * 1024));
        fprintf(file, "    <FreeVirtualMemory>%llu</FreeVirtualMemory>\n", statex.ullAvailVirtual / (1024 * 1024));
        fprintf(file, "  </Memory>\n");
    } else {
        fprintf(file, "  <Memory>\n    <Error>Не вдалося отримати інформацію про пам'ять.</Error>\n  </Memory>\n");
    }

    DWORD uptime = GetTickCount64() / 1000;
    DWORD days = uptime / (24 * 3600);
    uptime %= (24 * 3600);
    DWORD hours = uptime / 3600;
    uptime %= 3600;
    DWORD minutes = uptime / 60;
    uptime %= 60;
    DWORD seconds = uptime;
    fprintf(file, "  <SystemUptime>\n    <Days>%lu</Days>\n    <Hours>%lu</Hours>\n    <Minutes>%lu</Minutes>\n    <Seconds>%lu</Seconds>\n  </SystemUptime>\n", days, hours, minutes, seconds);

    ULARGE_INTEGER freeBytes;
    ULARGE_INTEGER totalBytes;
    if (GetDiskFreeSpaceEx("C:\\", &freeBytes, &totalBytes, NULL)) {
        fprintf(file, "  <Disk>\n");
        fprintf(file, "    <FreeSpace>%llu</FreeSpace>\n", freeBytes.QuadPart / (1024 * 1024));
        fprintf(file, "    <TotalSpace>%llu</TotalSpace>\n", totalBytes.QuadPart / (1024 * 1024));
        fprintf(file, "  </Disk>\n");
    } else {
        fprintf(file, "  <Disk>\n    <Error>Не вдалося отримати інформацію про диск C:\\.</Error>\n  </Disk>\n");
    }

    IP_ADAPTER_INFO adapterInfo[16];
    DWORD dwBufLen = sizeof(adapterInfo);
    DWORD dwRetVal = GetAdaptersInfo(adapterInfo, &dwBufLen);
    if (dwRetVal == ERROR_SUCCESS) {
        fprintf(file, "  <NetworkAdapters>\n");
        PIP_ADAPTER_INFO pAdapterInfo = adapterInfo;
        while (pAdapterInfo) {
            fprintf(file, "    <Adapter>\n");
            fprintf(file, "      <Description>%s</Description>\n", pAdapterInfo->Description);
            fprintf(file, "      <IPAddress>%s</IPAddress>\n", pAdapterInfo->IpAddressList.IpAddress.String);
            fprintf(file, "    </Adapter>\n");
            pAdapterInfo = pAdapterInfo->Next;
        }
        fprintf(file, "  </NetworkAdapters>\n");
    } else {
        fprintf(file, "  <NetworkAdapters>\n    <Error>Не вдалося отримати інформацію про мережеві адаптери.</Error>\n  </NetworkAdapters>\n");
    }

    DWORD processID = GetCurrentProcessId();
    HANDLE hProcess = OpenProcess(PROCESS_QUERY_INFORMATION | PROCESS_VM_READ, FALSE, processID);
    if (hProcess) {
        HMODULE hMods[1024];
        DWORD cbNeeded;
        if (EnumProcessModules(hProcess, hMods, sizeof(hMods), &cbNeeded)) {
            fprintf(file, "  <LoadedLibraries>\n");
            for (unsigned int i = 0; i < cbNeeded / sizeof(HMODULE); i++) {
                char szModName[MAX_PATH];
                if (GetModuleFileNameEx(hProcess, hMods[i], szModName, sizeof(szModName) / sizeof(char))) {
                    fprintf(file, "    <Library>%s</Library>\n", szModName);
                }
            }
            fprintf(file, "  </LoadedLibraries>\n");
        } else {
            fprintf(file, "  <LoadedLibraries>\n    <Error>Не вдалося отримати список завантажених бібліотек.</Error>\n  </LoadedLibraries>\n");
        }
        CloseHandle(hProcess);
    }

    fprintf(file, "</SystemInfo>\n");

    fclose(file);
}

int main() {
    getSystemInfo();
    return 0;
}
