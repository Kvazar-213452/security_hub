#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/sysinfo.h>
#include <unistd.h>
#include <sys/statvfs.h>
#include <time.h>
#include <arpa/inet.h>
#include <ifaddrs.h>
#include <sys/utsname.h>

void getSystemInfo() {
    FILE *file = fopen("data/system_info.xml", "w");
    if (file == NULL) {
        printf("Не вдалося відкрити файл для запису.\n");
        return;
    }

    fprintf(file, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n");
    fprintf(file, "<SystemInfo>\n");

    fprintf(file, "  <Architecture>x86_64 (AMD or Intel)</Architecture>\n");

    fprintf(file, "  <ProcessorCount>%ld</ProcessorCount>\n", sysconf(_SC_NPROCESSORS_ONLN));  // Замінено на %ld

    struct utsname buffer;
    if (uname(&buffer) == 0) {
        fprintf(file, "  <OS>\n");
        fprintf(file, "    <Name>%s</Name>\n", buffer.sysname);
        fprintf(file, "    <Version>%s</Version>\n", buffer.release);
        fprintf(file, "  </OS>\n");
    } else {
        fprintf(file, "  <OS>\n    <Error>Не вдалося отримати версію операційної системи.</Error>\n  </OS>\n");
    }

    struct sysinfo info;
    if (sysinfo(&info) == 0) {
        fprintf(file, "  <Memory>\n");
        fprintf(file, "    <FreeMemory>%ld</FreeMemory>\n", info.freeram / (1024 * 1024));
        fprintf(file, "    <TotalMemory>%ld</TotalMemory>\n", info.totalram / (1024 * 1024));
        fprintf(file, "    <FreeVirtualMemory>%ld</FreeVirtualMemory>\n", info.freeswap / (1024 * 1024));
        fprintf(file, "  </Memory>\n");
    } else {
        fprintf(file, "  <Memory>\n    <Error>Не вдалося отримати інформацію про пам'ять.</Error>\n  </Memory>\n");
    }

    FILE *uptimeFile = fopen("/proc/uptime", "r");
    if (uptimeFile) {
        double uptime;
        fscanf(uptimeFile, "%lf", &uptime);
        fclose(uptimeFile);

        int days = uptime / 86400;
        uptime -= days * 86400;
        int hours = uptime / 3600;
        uptime -= hours * 3600;
        int minutes = uptime / 60;
        uptime -= minutes * 60;
        int seconds = uptime;

        fprintf(file, "  <SystemUptime>\n");
        fprintf(file, "    <Days>%d</Days>\n", days);
        fprintf(file, "    <Hours>%d</Hours>\n", hours);
        fprintf(file, "    <Minutes>%d</Minutes>\n", minutes);
        fprintf(file, "    <Seconds>%d</Seconds>\n", seconds);
        fprintf(file, "  </SystemUptime>\n");
    }

    struct statvfs stat;
    if (statvfs("/", &stat) == 0) {
        fprintf(file, "  <Disk>\n");
        fprintf(file, "    <FreeSpace>%lu</FreeSpace>\n", stat.f_bfree * stat.f_frsize / (1024 * 1024));
        fprintf(file, "    <TotalSpace>%lu</TotalSpace>\n", stat.f_blocks * stat.f_frsize / (1024 * 1024));
        fprintf(file, "  </Disk>\n");
    } else {
        fprintf(file, "  <Disk>\n    <Error>Не вдалося отримати інформацію про диск.</Error>\n  </Disk>\n");
    }

    struct ifaddrs *ifaddr;
    if (getifaddrs(&ifaddr) == 0) {
        fprintf(file, "  <NetworkAdapters>\n");
        struct ifaddrs *ifa = ifaddr;
        while (ifa != NULL) {
            if (ifa->ifa_addr && ifa->ifa_addr->sa_family == AF_INET) {
                char ip[INET_ADDRSTRLEN];
                void *tmpAddrPtr = &((struct sockaddr_in *)ifa->ifa_addr)->sin_addr;
                inet_ntop(AF_INET, tmpAddrPtr, ip, INET_ADDRSTRLEN);
                fprintf(file, "    <Adapter>\n");
                fprintf(file, "      <Description>%s</Description>\n", ifa->ifa_name);
                fprintf(file, "      <IPAddress>%s</IPAddress>\n", ip);
                fprintf(file, "    </Adapter>\n");
            }
            ifa = ifa->ifa_next;
        }
        fprintf(file, "  </NetworkAdapters>\n");
        freeifaddrs(ifaddr);
    } else {
        fprintf(file, "  <NetworkAdapters>\n    <Error>Не вдалося отримати інформацію про мережеві адаптери.</Error>\n  </NetworkAdapters>\n");
    }

    fprintf(file, "  <LoadedLibraries>\n");
    FILE *procFile = fopen("/proc/self/maps", "r");
    if (procFile) {
        char line[256];
        while (fgets(line, sizeof(line), procFile)) {
            fprintf(file, "    <Library>%s</Library>\n", line);
        }
        fclose(procFile);
    }
    fprintf(file, "  </LoadedLibraries>\n");

    fprintf(file, "</SystemInfo>\n");

    fclose(file);
}

int main() {
    getSystemInfo();
    return 0;
}
