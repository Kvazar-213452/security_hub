#include "include/config.hpp"
#include "include/server.hpp"
#include "include/func_core.hpp"
#include <iostream>
#include <thread>

// core/main.cpp

int main() {
    Config config;

    int port = config.get<int>("port");
    if (port == 0) {
        port = FindFreePort();
    }
    
    start_server(port);

    return 0;
}
