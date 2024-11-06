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

    SYSTEM_INFO si;
    GetSystemInfo(&si);

    switch (si.wProcessorArchitecture) {
        case PROCESSOR_ARCHITECTURE_AMD64:
            fprintf(file, "Архітектура: x64 (AMD or Intel)\n");
            break;
        case PROCESSOR_ARCHITECTURE_ARM:
            fprintf(file, "Архітектура: ARM\n");
            break;
        case PROCESSOR_ARCHITECTURE_IA64:
            fprintf(file, "Архітектура: IA64 (Intel Itanium)\n");
            break;
        default:
            fprintf(file, "Архітектура: невідома\n");
            break;
    }

    fprintf(file, "Кількість процесорів: %u\n", si.dwNumberOfProcessors);

    OSVERSIONINFOEX osvi;
    osvi.dwOSVersionInfoSize = sizeof(OSVERSIONINFOEX);
    if (GetVersionEx((OSVERSIONINFO*)&osvi)) {
        fprintf(file, "Операційна система: ");
        if (osvi.dwMajorVersion == 10) {
            fprintf(file, "Windows 10\n");
        } else if (osvi.dwMajorVersion == 6) {
            if (osvi.dwMinorVersion == 3) {
                fprintf(file, "Windows 8.1\n");
            } else if (osvi.dwMinorVersion == 2) {
                fprintf(file, "Windows 8\n");
            } else if (osvi.dwMinorVersion == 1) {
                fprintf(file, "Windows 7\n");
            } else if (osvi.dwMinorVersion == 0) {
                fprintf(file, "Windows Vista\n");
            }
        } else if (osvi.dwMajorVersion == 5) {
            if (osvi.dwMinorVersion == 1) {
                fprintf(file, "Windows XP\n");
            } else if (osvi.dwMinorVersion == 0) {
                fprintf(file, "Windows 2000\n");
            }
        }
        fprintf(file, "Версія: %d.%d\n", osvi.dwMajorVersion, osvi.dwMinorVersion);
    } else {
        fprintf(file, "Не вдалося отримати версію операційної системи.\n");
    }

    MEMORYSTATUSEX statex;
    statex.dwLength = sizeof(statex);
    if (GlobalMemoryStatusEx(&statex)) {
        fprintf(file, "Вільна пам'ять: %llu MB\n", statex.ullAvailPhys / (1024 * 1024));
        fprintf(file, "Загальна пам'ять: %llu MB\n", statex.ullTotalPhys / (1024 * 1024));
        fprintf(file, "Вільна віртуальна пам'ять: %llu MB\n", statex.ullAvailVirtual / (1024 * 1024));
    } else {
        fprintf(file, "Не вдалося отримати інформацію про пам'ять.\n");
    }

    DWORD uptime = GetTickCount64() / 1000;
    DWORD days = uptime / (24 * 3600);
    uptime %= (24 * 3600);
    DWORD hours = uptime / 3600;
    uptime %= 3600;
    DWORD minutes = uptime / 60;
    uptime %= 60;
    DWORD seconds = uptime;
    fprintf(file, "Час роботи системи: %lu днів, %lu годин, %lu хвилин, %lu секунд\n", days, hours, minutes, seconds);

    ULARGE_INTEGER freeBytes;
    ULARGE_INTEGER totalBytes;
    if (GetDiskFreeSpaceEx("C:\\", &freeBytes, &totalBytes, NULL)) {
        fprintf(file, "Вільне місце на диску C:\\: %llu MB\n", freeBytes.QuadPart / (1024 * 1024));
        fprintf(file, "Загальний обсяг диску C:\\: %llu MB\n", totalBytes.QuadPart / (1024 * 1024));
    } else {
        fprintf(file, "Не вдалося отримати інформацію про диск C:\\.\n");
    }

    IP_ADAPTER_INFO adapterInfo[16];
    DWORD dwBufLen = sizeof(adapterInfo);
    DWORD dwRetVal = GetAdaptersInfo(adapterInfo, &dwBufLen);
    if (dwRetVal == ERROR_SUCCESS) {
        PIP_ADAPTER_INFO pAdapterInfo = adapterInfo;
        while (pAdapterInfo) {
            fprintf(file, "Мережевий адаптер: %s\n", pAdapterInfo->Description);
            fprintf(file, "IP-адреса: %s\n", pAdapterInfo->IpAddressList.IpAddress.String);
            pAdapterInfo = pAdapterInfo->Next;
        }
    } else {
        fprintf(file, "Не вдалося отримати інформацію про мережеві адаптери.\n");
    }

    DWORD processID = GetCurrentProcessId();
    HANDLE hProcess = OpenProcess(PROCESS_QUERY_INFORMATION | PROCESS_VM_READ, FALSE, processID);
    if (hProcess) {
        HMODULE hMods[1024];
        DWORD cbNeeded;
        if (EnumProcessModules(hProcess, hMods, sizeof(hMods), &cbNeeded)) {
            fprintf(file, "Завантажені бібліотеки:\n");
            for (unsigned int i = 0; i < cbNeeded / sizeof(HMODULE); i++) {
                char szModName[MAX_PATH];
                if (GetModuleFileNameEx(hProcess, hMods[i], szModName, sizeof(szModName) / sizeof(char))) {
                    fprintf(file, "%s\n", szModName);
                }
            }
        } else {
            fprintf(file, "Не вдалося отримати список завантажених бібліотек.\n");
        }
        CloseHandle(hProcess);
    }

    fclose(file);
}

int main() {
    getSystemInfo();
    return 0;
}
