#include <windows.h>
#include <stdio.h>
#include <time.h>
#include <sysinfoapi.h>
#include <psapi.h>
#include <iphlpapi.h>
#include <tchar.h>

void getSystemInfo() {
    SYSTEM_INFO si;
    GetSystemInfo(&si);

    switch (si.wProcessorArchitecture) {
        case PROCESSOR_ARCHITECTURE_AMD64:
            printf("Архітектура: x64 (AMD or Intel)\n");
            break;
        case PROCESSOR_ARCHITECTURE_ARM:
            printf("Архітектура: ARM\n");
            break;
        case PROCESSOR_ARCHITECTURE_IA64:
            printf("Архітектура: IA64 (Intel Itanium)\n");
            break;
        default:
            printf("Архітектура: невідома\n");
            break;
    }

    printf("Кількість процесорів: %u\n", si.dwNumberOfProcessors);

    OSVERSIONINFOEX osvi;
    osvi.dwOSVersionInfoSize = sizeof(OSVERSIONINFOEX);
    if (GetVersionEx((OSVERSIONINFO*)&osvi)) {
        printf("Операційна система: ");
        if (osvi.dwMajorVersion == 10) {
            printf("Windows 10\n");
        } else if (osvi.dwMajorVersion == 6) {
            if (osvi.dwMinorVersion == 3) {
                printf("Windows 8.1\n");
            } else if (osvi.dwMinorVersion == 2) {
                printf("Windows 8\n");
            } else if (osvi.dwMinorVersion == 1) {
                printf("Windows 7\n");
            } else if (osvi.dwMinorVersion == 0) {
                printf("Windows Vista\n");
            }
        } else if (osvi.dwMajorVersion == 5) {
            if (osvi.dwMinorVersion == 1) {
                printf("Windows XP\n");
            } else if (osvi.dwMinorVersion == 0) {
                printf("Windows 2000\n");
            }
        }
        printf("Версія: %d.%d\n", osvi.dwMajorVersion, osvi.dwMinorVersion);
    } else {
        printf("Не вдалося отримати версію операційної системи.\n");
    }

    MEMORYSTATUSEX statex;
    statex.dwLength = sizeof(statex);
    if (GlobalMemoryStatusEx(&statex)) {
        printf("Вільна пам'ять: %llu MB\n", statex.ullAvailPhys / (1024 * 1024));
        printf("Загальна пам'ять: %llu MB\n", statex.ullTotalPhys / (1024 * 1024));
        printf("Вільна віртуальна пам'ять: %llu MB\n", statex.ullAvailVirtual / (1024 * 1024));
    } else {
        printf("Не вдалося отримати інформацію про пам'ять.\n");
    }

    ULONGLONG uptime = GetTickCount64();
    ULONGLONG seconds = uptime / 1000;
    ULONGLONG minutes = seconds / 60;
    ULONGLONG hours = minutes / 60;
    ULONGLONG days = hours / 24;

    printf("Час роботи системи: %llu днів, %llu годин, %llu хвилин, %llu секунд\n", days, hours % 24, minutes % 60, seconds % 60);

    char drive[] = "C:\\";
    ULARGE_INTEGER freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes;
    if (GetDiskFreeSpaceEx(drive, &freeBytesAvailable, &totalNumberOfBytes, &totalNumberOfFreeBytes)) {
        printf("Вільне місце на диску %s: %llu MB\n", drive, freeBytesAvailable.QuadPart / (1024 * 1024));
        printf("Загальний обсяг диску %s: %llu MB\n", drive, totalNumberOfBytes.QuadPart / (1024 * 1024));
    } else {
        printf("Не вдалося отримати інформацію про диск.\n");
    }

    ULONG ulSize = 0;
    PIP_ADAPTER_INFO pAdapterInfo = NULL;
    if (GetAdaptersInfo(pAdapterInfo, &ulSize) == ERROR_BUFFER_OVERFLOW) {
        pAdapterInfo = (PIP_ADAPTER_INFO) malloc(ulSize);
        if (GetAdaptersInfo(pAdapterInfo, &ulSize) == ERROR_SUCCESS) {
            PIP_ADAPTER_INFO pAdapter = pAdapterInfo;
            while (pAdapter) {
                printf("Мережевий адаптер: %s\n", pAdapter->Description);
                printf("IP-адреса: %s\n", pAdapter->IpAddressList.IpAddress.String);
                pAdapter = pAdapter->Next;
            }
        }
        free(pAdapterInfo);
    } else {
        printf("Не вдалося отримати інформацію про мережеві адаптери.\n");
    }

    DWORD cbNeeded;
    DWORD dwProcessId = GetCurrentProcessId();
    HMODULE hMods[1024];
    if (EnumProcessModules(GetCurrentProcess(), hMods, sizeof(hMods), &cbNeeded)) {
        printf("Завантажені бібліотеки:\n");
        for (unsigned int i = 0; i < (cbNeeded / sizeof(HMODULE)); i++) {
            TCHAR szModName[MAX_PATH];
            if (GetModuleFileNameEx(GetCurrentProcess(), hMods[i], szModName, sizeof(szModName) / sizeof(TCHAR))) {
                printf("\t%s\n", szModName);
            }
        }
    } else {
        printf("Не вдалося отримати інформацію про бібліотеки.\n");
    }
}

int main() {
    getSystemInfo();
    return 0;
}
