g++ -shared -o cleanup.dll main.cpp
g++ -shared -o get_ssid.dll main.cpp -lwlanapi

go build -o FindFreePort.dll -buildmode=c-shared main.go