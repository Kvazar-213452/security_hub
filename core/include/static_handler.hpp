#ifndef STATIC_HANDLER_H
#define STATIC_HANDLER_H

#include <string>

namespace httplib {
    class Server;
}

namespace StaticHandler {
    std::string read_file_content(const std::string& file_path);
    bool ends_with(const std::string& str, const std::string& suffix);
    std::string get_mime_type(const std::string& path);
    void setup_static_routes(httplib::Server& svr, const std::string& base_dir);
}

#endif