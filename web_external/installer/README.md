gendef webview.dll
dlltool -v --dllname webview.dll --output-lib libwebview.a --input-def webview.def


g++ -std=c++17 -o main.exe main.cpp lib/httplib.cc src/server.cpp src/func_shell.cpp src/config.cpp src/html.cpp -static -static-libgcc -static-libstdc++ -lz -lminizip -lws2_32 -lole32 -lversion -lshlwapi -lshell32

