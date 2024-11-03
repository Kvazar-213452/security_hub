#include "lib/webview.h"
#include "lib/httplib.h"
#include "include/html.h"
#include "include/config.h"
#include "include/server.h"
#include "include/func_shell.h"

#include <iostream>
#include <thread>
#include <chrono>
#include <atomic>
#include <string>
#include <sstream>
#include <winsock2.h>
#include <ws2tcpip.h>

#pragma comment(lib, "ws2_32.lib")

int FindFreePort() {
    // Ініціалізація бібліотеки WinSock
    WSADATA wsaData;
    if (WSAStartup(MAKEWORD(2, 2), &wsaData) != 0) {
        std::cerr << "Помилка ініціалізації WinSock" << std::endl;
        return 0;
    }

    // Створення сокета
    SOCKET sock = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
    if (sock == INVALID_SOCKET) {
        std::cerr << "Помилка створення сокета" << std::endl;
        WSACleanup();
        return 0;
    }

    // Налаштування адреси для прив'язки (порт 0 дозволяє вибрати вільний порт)
    sockaddr_in addr;
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = inet_addr("127.0.0.1");
    addr.sin_port = htons(0);

    // Прив'язка сокета
    if (bind(sock, (sockaddr*)&addr, sizeof(addr)) == SOCKET_ERROR) {
        std::cerr << "Помилка прив'язки сокета" << std::endl;
        closesocket(sock);
        WSACleanup();
        return 0;
    }

    // Отримання вибраного порту
    int addrLen = sizeof(addr);
    if (getsockname(sock, (sockaddr*)&addr, &addrLen) == SOCKET_ERROR) {
        std::cerr << "Помилка отримання інформації про сокет" << std::endl;
        closesocket(sock);
        WSACleanup();
        return 0;
    }

    int freePort = ntohs(addr.sin_port);

    // Закриття сокета і очищення WinSock
    closesocket(sock);
    WSACleanup();

    return freePort;
}


std::atomic<bool> webview_closed(false);

void start_webview(int port) {
    try {
        std::string html_content_core = generate_html_content(port);

        webview::webview w(false, nullptr);
        w.set_title(name_app);
        w.set_size(window_h, window_w, WEBVIEW_HINT_NONE);
        w.set_html(html_content_core);
        w.run();
        webview_closed.store(true);
    } catch (const webview::exception &e) {
        std::cerr << e.what() << std::endl;
        exit(1);
    }
}

void monitor_webview() {
    while (!webview_closed.load()) {
        std::this_thread::sleep_for(std::chrono::milliseconds(100));
    }
    std::exit(0);
}

int WINAPI WinMain(HINSTANCE, HINSTANCE, LPSTR, int) {
    int port = FindFreePort();
    std::cerr << port << std::endl;

    std::thread server_thread(std::bind(start_server, port));

    std::this_thread::sleep_for(std::chrono::seconds(1));

    std::thread monitor_thread(monitor_webview);
    start_webview(port);
    monitor_thread.join(); 

    server_thread.join();
    return 0;
}