#include <winsock2.h>
#include <windows.h>
#include <iostream>
#include <string>
#include <sstream>
#include <fstream>
#include <vector>
#include <stdexcept>
#include <fstream>
#include <system_error>
#include <filesystem>

namespace fs = std::filesystem;

#pragma comment(lib, "ws2_32.lib")

typedef int (*FindFreePortFunc)();

int FindFreePort() {
    WSADATA wsaData;
    if (WSAStartup(MAKEWORD(2, 2), &wsaData) != 0) {
        std::cerr << "error" << std::endl;
        return 0;
    }

    SOCKET sock = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
    if (sock == INVALID_SOCKET) {
        std::cerr << "error" << std::endl;
        WSACleanup();
        return 0;
    }

    sockaddr_in addr;
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = inet_addr("127.0.0.1");
    addr.sin_port = htons(0);

    if (bind(sock, (sockaddr*)&addr, sizeof(addr)) == SOCKET_ERROR) {
        std::cerr << "error" << std::endl;
        closesocket(sock);
        WSACleanup();
        return 0;
    }

    int addrLen = sizeof(addr);
    if (getsockname(sock, (sockaddr*)&addr, &addrLen) == SOCKET_ERROR) {
        std::cerr << "error" << std::endl;
        closesocket(sock);
        WSACleanup();
        return 0;
    }

    int freePort = ntohs(addr.sin_port);

    closesocket(sock);
    WSACleanup();

    return freePort;
}

std::string read_file(const std::string& relative_path) {
    try {
        std::string full_path = relative_path;
        
        #if defined(__has_include) && (__has_include(<filesystem>) || __has_include(<experimental/filesystem>))
        fs::path file_path = fs::absolute(fs::path(relative_path));
        full_path = file_path.string();
        #endif
        
        std::ifstream file(full_path, std::ios::binary);
        if (!file.is_open()) {
            throw std::runtime_error("Failed to open file: " + full_path);
        }
        
        return std::string(
            (std::istreambuf_iterator<char>(file)),
            std::istreambuf_iterator<char>()
        );
    } catch (const std::exception& e) {
        std::cerr << "Error reading file: " << e.what() << std::endl;
        return "";
    }
}