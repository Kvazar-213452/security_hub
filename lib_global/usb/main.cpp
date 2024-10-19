#include <windows.h>
#include <iostream>
#include <fstream>
#include <setupapi.h>
#include <devguid.h>
#include <regstr.h>
#include <string>

#pragma comment(lib, "setupapi.lib")

void ListDevices() {
    HDEVINFO deviceInfoSet = SetupDiGetClassDevs(NULL, NULL, NULL, DIGCF_ALLCLASSES | DIGCF_PRESENT);
    if (deviceInfoSet == INVALID_HANDLE_VALUE) {
        return;
    }

    SP_DEVINFO_DATA deviceInfoData;
    deviceInfoData.cbSize = sizeof(SP_DEVINFO_DATA);
    DWORD index = 0;

    std::wofstream logFile("devices.log", std::ios::trunc);
    if (!logFile.is_open()) {
        SetupDiDestroyDeviceInfoList(deviceInfoSet);
        return;
    }

    while (SetupDiEnumDeviceInfo(deviceInfoSet, index, &deviceInfoData)) {
        index++;

        wchar_t deviceName[256]; 
        if (SetupDiGetDeviceRegistryPropertyW(deviceInfoSet, &deviceInfoData, SPDRP_DEVICEDESC, NULL, (PBYTE)deviceName, sizeof(deviceName), NULL)) {
            std::wstring deviceDescription = std::wstring(deviceName);
            logFile << deviceDescription << std::endl; 
        }
    }

    logFile.close();
    SetupDiDestroyDeviceInfoList(deviceInfoSet);
}

int main() {
    ListDevices();
    return 0;
}
