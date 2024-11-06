gcc -shared -o cleanup.dll main.c
gcc -o get_ssid main.c -lwlanapi -mwindows
gcc -o available_wifi main.c -lwlanapi -mwindows

go build -o FindFreePort.dll -buildmode=c-shared main.go