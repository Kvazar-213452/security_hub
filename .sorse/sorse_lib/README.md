gcc -shared -o cleanup.dll cleanup.c
gcc -o get_ssid get_ssid.c -static -static-libgcc -lwlanapi -mwindows
gcc -o available_wifi available_wifi.c -static -static-libgcc -lwlanapi -mwindows

go build -o FindFreePort.dll -buildmode=c-shared main.go