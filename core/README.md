gendef webview.dll
dlltool -v --dllname webview.dll --output-lib libwebview.a --input-def webview.def


g++ -std=c++17 -o main.exe main.cpp lib/httplib.cc src/server.cpp src/func_shell.cpp src/config.cpp src/html.cpp src/func_app.cpp -static -static-libgcc -static-libstdc++ -lminizip -lz -lbz2 -lws2_32 -lole32 -lversion -lshlwapi -lshell32 -mwindows

g++ -std=c++17 -o main.exe main.cpp lib/httplib.cc src/server.cpp src/func_shell.cpp src/config.cpp src/html.cpp src/func_app.cpp -static -static-libgcc -static-libstdc++ -lz -lbz2 -lws2_32 -lole32 -lversion -lshlwapi -lshell32 -mwindows

g++ -std=c++17 -o main.exe main.cpp lib/httplib.cc src/server.cpp src/func_core.cpp src/config.cpp -lws2_32 -lole32 -lversion -lshlwapi -lshell32

g++ -std=c++17 -o main.exe main.cpp src/server.cpp src/func_core.cpp src/config.cpp lib/httplib.o -lws2_32 -lole32 -lversion -lshlwapi -lshell32

g++ -std=c++17 -o main.exe main.cpp lib/httplib.o src/server.cpp src/api_handler.cpp src/routes.cpp src/static_handler.cpp src/func_core.cpp src/config.cpp -lws2_32 -lole32 -lversion -lshlwapi -lshell32 -lstdc++fs


g++ -std=c++17 -o main.exe main.cpp lib/httplib.o src/server.cpp src/api_handler.cpp src/routes.cpp src/static_handler.cpp src/func_core.cpp src/config.cpp src/shell_NM/run_NM1.cpp src/shell_NM/run_NM2.cpp src/shell_NM/run_NM3.cpp src/shell_NM/shared_vars.cpp -lws2_32 -lole32 -lversion -lshlwapi -lshell32