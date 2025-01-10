#include <windows.h>
#include <wlanapi.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#pragma comment(lib, "wlanapi.lib")
#pragma comment(lib, "ole32.lib")

// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⣷⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⣀⣶⣿⣿⣿⣿⣿⣿⣷⣶⣶⣶⣦⣀⡀⠀⢀⣴⣇⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀
// ⠀⠀⠀⠀⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀
// ⠀⠀⠀⣴⣿⣿⣿⣿⠛⣿⣿⣿⣿⣿⣿⣿⣿⡿⣿⣿⣿⣿⣿⣿⣿⣿⣄⠀⠀⠀
// ⠀⠀⣾⣿⣿⣿⣿⣿⣶⣿⣯⣭⣬⣉⣽⣿⣿⣄⣼⣿⣿⣿⣿⣿⣿⣿⣿⣷⡀⠀
// ⠀⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄
// ⢸⣿⣿⣿⣿⠟⠋⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠁⣿⣿⣿⣿⡿⠛⠉⠉⠉⠉⠁
// ⠘⠛⠛⠛⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠛⠛⠛⠃

void ListWifiNetworks() {
    DWORD dwResult = 0;
    HANDLE hClient = NULL;
    PWLAN_INTERFACE_INFO_LIST pIfList = NULL;
    PWLAN_INTERFACE_INFO pIfInfo = NULL;
    PWLAN_AVAILABLE_NETWORK_LIST pBssList = NULL;

    FILE *file = fopen("data/get_ssid.xml", "w");
    if (file == NULL) {
        return;
    }

    fprintf(file, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n");
    fprintf(file, "<Networks>\n");

    dwResult = WlanOpenHandle(2, NULL, &dwResult, &hClient);
    if (dwResult != ERROR_SUCCESS) {
        fclose(file);
        return;
    }

    dwResult = WlanEnumInterfaces(hClient, NULL, &pIfList);
    if (dwResult != ERROR_SUCCESS) {
        WlanCloseHandle(hClient, NULL);
        fclose(file);
        return;
    }

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

            fprintf(file, "  <SSID>");
            for (int k = 0; k < pNetwork->dot11Ssid.uSSIDLength; k++) {
                fprintf(file, "%c", pNetwork->dot11Ssid.ucSSID[k]);
            }
            fprintf(file, "</SSID>\n");

            j = 1000;
        }

        if (pBssList != NULL) {
            WlanFreeMemory(pBssList);
            pBssList = NULL;
        }
    }

    fprintf(file, "</Networks>\n");

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
