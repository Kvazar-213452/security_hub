#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void scan_wifi_networks() {
    FILE *file = fopen("data/available_wifi.xml", "w");
    if (file == NULL) {
        printf("Не вдалося відкрити файл для запису.\n");
        return;
    }

    fprintf(file, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n");
    fprintf(file, "<Networks>\n");

    // Run iwlist to get the list of available networks
    FILE *fp = popen("sudo iwlist scanning", "r");
    if (fp == NULL) {
        fprintf(file, "Помилка при виконанні сканування мереж.\n");
        fclose(file);
        return;
    }

    char line[256];
    char ssid[256];
    int signalQuality = 0;

    // Parse the output of iwlist
    while (fgets(line, sizeof(line), fp) != NULL) {
        if (strstr(line, "ESSID:") != NULL) {
            // Extract SSID
            sscanf(line, " ESSID:\"%[^\"]\"", ssid);
        }
        if (strstr(line, "Signal level=") != NULL) {
            // Extract signal strength
            sscanf(line, " Signal level=%d", &signalQuality);

            // Write the network info to the file
            fprintf(file, "  <Network>\n");
            fprintf(file, "    <SSID>%s</SSID>\n", ssid);
            fprintf(file, "    <SignalQuality>%d</SignalQuality>\n", signalQuality);
            fprintf(file, "  </Network>\n");
        }
    }

    fclose(fp);

    fprintf(file, "</Networks>\n");
    fclose(file);
}

int main() {
    scan_wifi_networks();
    return 0;
}
