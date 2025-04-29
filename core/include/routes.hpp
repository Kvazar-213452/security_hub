#ifndef ROUTES_H
#define ROUTES_H

#include <string>

namespace httplib {
    class Server;
}

namespace Routes {
    void setup_routes(httplib::Server& svr, const std::string& base_dir);
}

#endif