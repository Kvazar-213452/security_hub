#include <stdio.h>
#include <windows.h>
#include <wlanapi.h>
#include <objbase.h>
#include <initguid.h>

#pragma comment(lib, "wlanapi.lib")
#pragma comment(lib, "ole32.lib")

int main() {
    HANDLE hClient = NULL;
    DWORD dwVersion = 0;
    DWORD dwResult = 0;
    
    // Відкриваємо файл для запису XML
    FILE *file = fopen("data/available_wifi.xml", "w");
    if (file == NULL) {
        printf("Не вдалося відкрити файл для запису.\n");
        return 1;
    }

    // Початок XML документа
    fprintf(file, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n");
    fprintf(file, "<Networks>\n");

    // Отримуємо доступ до клієнта WLAN
    if (WlanOpenHandle(2, NULL, &dwVersion, &hClient) != ERROR_SUCCESS) {
        fprintf(file, "Помилка при відкритті ручки до клієнта WLAN.\n");
        fclose(file);
        return 1;
    }

    PWLAN_INTERFACE_INFO_LIST pIfList = NULL;
    if (WlanEnumInterfaces(hClient, NULL, &pIfList) != ERROR_SUCCESS) {
        fprintf(file, "Помилка при отриманні списку інтерфейсів.\n");
        fclose(file);
        return 1;
    }

    // Перебираємо інтерфейси
    for (DWORD i = 0; i < pIfList->dwNumberOfItems; i++) {
        PWLAN_AVAILABLE_NETWORK_LIST pNetworkList = NULL;

        // Отримуємо список доступних мереж
        if (WlanGetAvailableNetworkList(hClient, &pIfList->InterfaceInfo[i].InterfaceGuid, 0, NULL, &pNetworkList) != ERROR_SUCCESS) {
            fprintf(file, "Помилка при отриманні списку доступних мереж.\n");
            continue;
        }

        // Перебираємо кожну мережу
        for (DWORD j = 0; j < pNetworkList->dwNumberOfItems; j++) {
            WLAN_AVAILABLE_NETWORK network = pNetworkList->Network[j];

            // Записуємо SSID у форматі XML
            fprintf(file, "  <Network>\n");
            fprintf(file, "    <SSID>");
            for (DWORD k = 0; k < network.dot11Ssid.uSSIDLength; k++) {
                fprintf(file, "%c", network.dot11Ssid.ucSSID[k]);
            }
            fprintf(file, "</SSID>\n");

            // Записуємо рівень сигналу
            fprintf(file, "    <SignalQuality>%d</SignalQuality>\n", network.wlanSignalQuality);
            fprintf(file, "  </Network>\n");
        }

        // Очищаємо пам'ять
        if (pNetworkList != NULL) {
            WlanFreeMemory(pNetworkList);
        }
    }

    // Очищаємо пам'ять та закриваємо клієнта WLAN
    if (pIfList != NULL) {
        WlanFreeMemory(pIfList);
    }

    if (hClient != NULL) {
        WlanCloseHandle(hClient, NULL);
    }

    // Закриваємо файл і завершуємо XML документ
    fprintf(file, "</Networks>\n");
    fclose(file);

    return 0;
}
