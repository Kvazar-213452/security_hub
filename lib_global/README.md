g++ -shared -o cleanup.dll main.cpp

go build -o FindFreePort.dll -buildmode=c-shared main.go