#include <windows.h>
#include <wlanapi.h>
#include <string>
#include <fstream>

#pragma comment(lib, "wlanapi.lib")

extern "C" __declspec(dllexport) void GetConnectedSSIDAndWriteToFile() {
    HANDLE wlanHandle = NULL;
    DWORD dwMaxClient = 2;
    DWORD dwCurVersion = 0;
    DWORD dwResult = WlanOpenHandle(dwMaxClient, NULL, &dwCurVersion, &wlanHandle);

    if (dwResult != ERROR_SUCCESS) {
        return;
    }

    PWLAN_INTERFACE_INFO_LIST pIfList = NULL;
    dwResult = WlanEnumInterfaces(wlanHandle, NULL, &pIfList);
    if (dwResult != ERROR_SUCCESS) {
        WlanCloseHandle(wlanHandle, NULL);
        return;
    }

    std::string ssid = "";

    for (int i = 0; i < (int)pIfList->dwNumberOfItems; i++) {
        PWLAN_INTERFACE_INFO pIfInfo = &pIfList->InterfaceInfo[i];

        PWLAN_CONNECTION_ATTRIBUTES pConnectInfo = NULL;
        DWORD dwDataSize = sizeof(WLAN_CONNECTION_ATTRIBUTES);
        dwResult = WlanQueryInterface(wlanHandle, &pIfInfo->InterfaceGuid, wlan_intf_opcode_current_connection, NULL, &dwDataSize, (PVOID*)&pConnectInfo, NULL);

        if (dwResult == ERROR_SUCCESS) {
            ssid = std::string((char*)pConnectInfo->wlanAssociationAttributes.dot11Ssid.ucSSID, pConnectInfo->wlanAssociationAttributes.dot11Ssid.uSSIDLength);
            break;
        }
    }

    if (pIfList) {
        WlanFreeMemory(pIfList);
    }
    WlanCloseHandle(wlanHandle, NULL);

    if (!ssid.empty()) {
        std::ofstream outFile("data.txt");
        outFile << "Connected SSID: " << ssid << std::endl;
        outFile.close();
    }
}
