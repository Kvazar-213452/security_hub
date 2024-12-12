#include <winsock2.h>
#include <windows.h>

#include <iostream>
#include <string>
#include <sstream>
#include <fstream>
#include <vector>
#include <stdexcept>
#include <filesystem>

#pragma comment(lib, "ws2_32.lib")

typedef int (*FindFreePortFunc)();

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

int FindFreePort() {
    WSADATA wsaData;
    if (WSAStartup(MAKEWORD(2, 2), &wsaData) != 0) {
        std::cerr << "Помилка ініціалізації WinSock" << std::endl;
        return 0;
    }

    SOCKET sock = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
    if (sock == INVALID_SOCKET) {
        std::cerr << "Помилка створення сокета" << std::endl;
        WSACleanup();
        return 0;
    }

    sockaddr_in addr;
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = inet_addr("127.0.0.1");
    addr.sin_port = htons(0);

    if (bind(sock, (sockaddr*)&addr, sizeof(addr)) == SOCKET_ERROR) {
        std::cerr << "Помилка прив'язки сокета" << std::endl;
        closesocket(sock);
        WSACleanup();
        return 0;
    }

    int addrLen = sizeof(addr);
    if (getsockname(sock, (sockaddr*)&addr, &addrLen) == SOCKET_ERROR) {
        std::cerr << "Помилка отримання інформації про сокет" << std::endl;
        closesocket(sock);
        WSACleanup();
        return 0;
    }

    int freePort = ntohs(addr.sin_port);

    closesocket(sock);
    WSACleanup();

    return freePort;
}

void deleteFile(const std::string& filePath) {
    try {
        if (std::filesystem::exists(filePath)) {
            std::filesystem::remove(filePath);
            std::cout << "Файл '" << filePath << "' було успішно видалено." << std::endl;
        } else {
            std::cout << "Файл '" << filePath << "' не знайдено." << std::endl;
        }
    } catch (const std::filesystem::filesystem_error& e) {
        std::cerr << "Помилка: " << e.what() << std::endl;
    }
}
