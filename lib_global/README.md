go build -buildmode=c-shared -o find_free_port.dll main.go

g++ -shared -o system_info.dll main.cpp
g++ -o usb_info.exe main.cpp -lsetupapi -mwindows
g++ -o resource_info.exe main.cpp -lpsapi