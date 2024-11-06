#include <windows.h>
#include <wlanapi.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#pragma comment(lib, "wlanapi.lib")
#pragma comment(lib, "ole32.lib")

__declspec(dllexport) const char* ListWifiNetworks() {
    static char networkList[1024];  // використовуємо static, щоб дані не були втрачені після виконання
    memset(networkList, 0, sizeof(networkList)); // очищаємо масив

    DWORD dwResult = 0;
    HANDLE hClient = NULL;
    PWLAN_INTERFACE_INFO_LIST pIfList = NULL;
    PWLAN_INTERFACE_INFO pIfInfo = NULL;
    PWLAN_AVAILABLE_NETWORK_LIST pBssList = NULL;

    dwResult = WlanOpenHandle(2, NULL, &dwResult, &hClient);
    if (dwResult != ERROR_SUCCESS) {
        return "WlanOpenHandle failed";
    }

    dwResult = WlanEnumInterfaces(hClient, NULL, &pIfList);
    if (dwResult != ERROR_SUCCESS) {
        WlanCloseHandle(hClient, NULL);
        return "WlanEnumInterfaces failed";
    }

    int networkIndex = 0;
    for (int i = 0; i < (int)pIfList->dwNumberOfItems; i++) {
        pIfInfo = &pIfList->InterfaceInfo[i];

        dwResult = WlanScan(hClient, &pIfInfo->InterfaceGuid, NULL, NULL, NULL);
        if (dwResult != ERROR_SUCCESS) {
            continue;
        }

        dwResult = WlanGetAvailableNetworkList(hClient, &pIfInfo->InterfaceGuid, 0, NULL, &pBssList);
        if (dwResult != ERROR_SUCCESS) {
            continue;
        }

        for (int j = 0; j < (int)pBssList->dwNumberOfItems; j++) {
            PWLAN_AVAILABLE_NETWORK pNetwork = &pBssList->Network[j];
            for (int k = 0; k < pNetwork->dot11Ssid.uSSIDLength; k++) {
                if (networkIndex < sizeof(networkList) - 1) {
                    networkList[networkIndex++] = pNetwork->dot11Ssid.ucSSID[k];
                }
            }
            networkList[networkIndex++] = '\n';  // додаємо новий рядок між мережами
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

    return networkList;
}
