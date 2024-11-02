#include <iostream>
#include <windows.h>
#include <string>
#include <sstream>

typedef int (*FindFreePortFunc)();

int port_find() {
    HMODULE hModule = LoadLibrary(TEXT("FindFreePort.dll"));
    if (hModule == nullptr) {
        std::cerr << "Не вдалося завантажити DLL" << std::endl;
        return 1;
    }

    FindFreePortFunc FindFreePort = (FindFreePortFunc)GetProcAddress(hModule, "FindFreePort");
    if (FindFreePort == nullptr) {
        std::cerr << "Не вдалося знайти функцію" << std::endl;
        FreeLibrary(hModule);
        return 1;
    }

    int port = FindFreePort();
    std::cerr << port << std::endl;
    FreeLibrary(hModule);
    return port;
}

std::string generate_html_content(int port) {
    std::ostringstream url;
    url << "http://127.0.0.1:" << port << "/";

    std::string html_content_core = R"(
        <style>
            iframe{
                position: fixed;
                height: 100%;
                width: 100%;
                top: 0%;
                left: 0%;
            }
        </style>
        <iframe src=")" + url.str() + R"(" frameborder="0"></iframe>
    )";

    return html_content_core;
}
