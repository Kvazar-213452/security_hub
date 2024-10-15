#include <windows.h>
#include <setupapi.h>
#include <iostream>
#include <string>

#pragma comment(lib, "setupapi.lib")

extern "C" {
    __declspec(dllexport) const char* GetConnectedDevices() {
        static std::string result;
        result.clear();

        HDEVINFO deviceInfoSet = SetupDiGetClassDevs(NULL, "USB", NULL, DIGCF_PRESENT | DIGCF_ALLCLASSES);

        if (deviceInfoSet == INVALID_HANDLE_VALUE) {
            result = "Error: Unable to get device information.";
            return result.c_str();
        }

        SP_DEVINFO_DATA deviceInfoData;
        deviceInfoData.cbSize = sizeof(SP_DEVINFO_DATA);

        for (DWORD i = 0; SetupDiEnumDeviceInfo(deviceInfoSet, i, &deviceInfoData); i++) {
            char deviceName[256];
            if (SetupDiGetDeviceRegistryPropertyA(deviceInfoSet, &deviceInfoData, SPDRP_DEVICEDESC, NULL, (PBYTE)deviceName, sizeof(deviceName), NULL)) {
                result += "Device " + std::to_string(i + 1) + ": " + std::string(deviceName) + "\n";
            }
        }

        SetupDiDestroyDeviceInfoList(deviceInfoSet);
        return result.c_str();
    }
}
