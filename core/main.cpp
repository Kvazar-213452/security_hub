#include <iostream>
#include <thread>
#include <atomic>
#include <windows.h>
#include <string>
#include <filesystem>
#include "include/config.hpp"
#include "include/server.hpp"
#include "include/func_core.hpp"
#include "include/shell_NM/run_NM.hpp"

extern std::atomic<bool> g_shellWebRunning;
extern std::atomic<bool> g_serverRunning;

int main(int argc, char* argv[]) {
    Config config;
    
    int port = 0;
    if (argc > 1) {
        try {
            port = std::stoi(argv[1]);
            if (port < 1 || port > 65535) {
                std::cerr << "Помилка: Порт повинен бути в діапазоні 1-65535" << std::endl;
                return 1;
            }
        } catch (const std::exception& e) {
            std::cerr << "Помилка: Невірний формат порту. Використовуйте число." << std::endl;
            return 1;
        }
    }

    std::thread shellWebThread;
    int shell = config.get<int>("shell");
    int visualization = config.get<int>("visualization");

    if (visualization == 1) {
        if (shell == 0) {
            shellWebThread = std::thread([&]() {run_NM1(port);});
        } else if (shell == 1) {
            shellWebThread = std::thread([&]() {run_NM2(port);});
        } else if (shell == 2) {
            shellWebThread = std::thread([&]() {run_NM3(port);});
        }
    }

    write_starter_md("http://localhost:" + std::to_string(port) + "/");

    std::thread serverThread([&]() {start_server(port);});

    if (visualization == 1) {
        shellWebThread.detach();
    }
    serverThread.detach();
    
    while (g_shellWebRunning && g_serverRunning) {
        std::this_thread::sleep_for(std::chrono::milliseconds(100));
    }
    
    if (!g_shellWebRunning) {
        std::exit(0);
    }
    
    return 0;
}