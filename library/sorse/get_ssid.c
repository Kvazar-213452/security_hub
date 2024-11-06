#include <windows.h>
#include <wlanapi.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#pragma comment(lib, "wlanapi.lib")
#pragma comment(lib, "ole32.lib")

void ListWifiNetworks() {
    DWORD dwResult = 0;
    HANDLE hClient = NULL;
    PWLAN_INTERFACE_INFO_LIST pIfList = NULL;
    PWLAN_INTERFACE_INFO pIfInfo = NULL;
    PWLAN_AVAILABLE_NETWORK_LIST pBssList = NULL;

    FILE *file = fopen("data/file.txt", "w");
    if (file == NULL) {
        fprintf(stderr, "Не вдалося відкрити файл для запису\n");
        return;
    }

    dwResult = WlanOpenHandle(2, NULL, &dwResult, &hClient);
    if (dwResult != ERROR_SUCCESS) {
        fprintf(stderr, "WlanOpenHandle failed with error: %u\n", dwResult);
        fclose(file);
        return;
    }

    dwResult = WlanEnumInterfaces(hClient, NULL, &pIfList);
    if (dwResult != ERROR_SUCCESS) {
        fprintf(stderr, "WlanEnumInterfaces failed with error: %u\n", dwResult);
        WlanCloseHandle(hClient, NULL);
        fclose(file);
        return;
    }

    for (int i = 0; i < (int)pIfList->dwNumberOfItems; i++) {
        pIfInfo = &pIfList->InterfaceInfo[i];

        dwResult = WlanScan(hClient, &pIfInfo->InterfaceGuid, NULL, NULL, NULL);
        if (dwResult != ERROR_SUCCESS) {
            fprintf(stderr, "WlanScan failed with error: %u\n", dwResult);
            continue;
        }

        dwResult = WlanGetAvailableNetworkList(hClient, &pIfInfo->InterfaceGuid, 0, NULL, &pBssList);
        if (dwResult != ERROR_SUCCESS) {
            fprintf(stderr, "WlanGetAvailableNetworkList failed with error: %u\n", dwResult);
            continue;
        }

        for (int j = 0; j < (int)pBssList->dwNumberOfItems; j++) {
            PWLAN_AVAILABLE_NETWORK pNetwork = &pBssList->Network[j];

            for (int k = 0; k < pNetwork->dot11Ssid.uSSIDLength; k++) {
                fprintf(file, "%c", pNetwork->dot11Ssid.ucSSID[k]);
            }

            j = 100;
        }

        if (pBssList != NULL) {
            WlanFreeMemory(pBssList);
            pBssList = NULL;
        }
    }

    if (pIfList != NULL) {
        WlanFreeMemory(pIfList);
        pIfList = NULL;
    }

    WlanCloseHandle(hClient, NULL);
    fclose(file);
}

int main() {
    ListWifiNetworks();
    return 0;
}
