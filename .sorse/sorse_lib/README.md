gcc -o get_ssid get_ssid.c -lm
gcc -o available_wifi available_wifi.c -static -static-libgcc
gcc -o system_info system_info.c -static -L/usr/lib -lpcap

go build -o FindFreePort.dll -buildmode=c-shared main.go