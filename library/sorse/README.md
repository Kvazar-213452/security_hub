g++ -shared -o cleanup.dll main.cpp

gcc -o get_ssid main.c -lwlanapi
gcc -o geavailable_wifi_ssid main.c -lwlanapi

go build -o FindFreePort.dll -buildmode=c-shared main.go