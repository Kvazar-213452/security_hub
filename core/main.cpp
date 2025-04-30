#include <iostream>
#include <thread>
#include <atomic>
#include "include/config.hpp"
#include "include/server.hpp"
#include "include/func_core.hpp"
#include "include/shell_NM/run_NM.hpp"

extern std::atomic<bool> g_shellWebRunning;
extern std::atomic<bool> g_serverRunning;

int main() {
    Config config;

    std::thread shellWebThread;
    int shell = config.get<int>("shell");
    int port = config.get<int>("port");

    if (port == 0) {
        port = FindFreePort();
    }

    if (shell > 0) {
        if (shell == 1) {
            shellWebThread = std::thread([&]() {run_NM1(port);});
        } else if (shell == 2) {
            shellWebThread = std::thread([&]() {run_NM2(port);});
        } else if (shell == 3) {
            shellWebThread = std::thread([&]() {run_NM3(port);});
        }
    } else {
        write_starter_md("http://localhost:" + std::to_string(port) + "/");
    }

    std::thread serverThread([&]() {start_server(port);});

    if (shell > 0 && (shell == 1 || shell == 2)) {
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