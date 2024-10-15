go build -buildmode=c-shared -o find_free_port.dll main.go

g++ -shared -o system_info.dll main.cpp
g++ -shared -o usb_info.dll main.cpp -lsetupapi   