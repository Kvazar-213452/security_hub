#include "../include/server.hpp"
#include "../include/routes.hpp"
#include "../lib/httplib.h"
#include <iostream>

void start_server(int port) {
    httplib::Server svr;
    std::string base_dir = "web";
    
    Routes::setup_routes(svr, base_dir);

    std::cout << "Server started at http://localhost:" << port << std::endl;
    svr.listen("127.0.0.1", port);
}