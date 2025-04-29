#include "../include/static_handler.hpp"
#include "../lib/httplib.h"
#include <fstream>
#include <algorithm>
#include <iostream>
#include <filesystem>

namespace fs = std::filesystem;

std::string StaticHandler::read_file_content(const std::string& file_path) {
    std::ifstream file(file_path, std::ios::binary);
    if (!file.is_open()) {
        std::cerr << "Failed to open file: " << file_path << std::endl;
        return "";
    }
    return std::string((std::istreambuf_iterator<char>(file)), 
                      std::istreambuf_iterator<char>());
}

bool StaticHandler::ends_with(const std::string& str, const std::string& suffix) {
    if (suffix.size() > str.size()) return false;
    return std::equal(suffix.rbegin(), suffix.rend(), str.rbegin());
}

std::string StaticHandler::get_mime_type(const std::string& path) {
    if (ends_with(path, ".html")) return "text/html";
    if (ends_with(path, ".css")) return "text/css";
    if (ends_with(path, ".js")) return "application/javascript";
    if (ends_with(path, ".png")) return "image/png";
    if (ends_with(path, ".jpg") || ends_with(path, ".jpeg")) return "image/jpeg";
    if (ends_with(path, ".gif")) return "image/gif";
    if (ends_with(path, ".svg")) return "image/svg+xml";
    if (ends_with(path, ".json")) return "application/json";
    return "application/octet-stream";
}

void StaticHandler::setup_static_routes(httplib::Server& svr, const std::string& base_dir) {
    std::string static_dir = base_dir + "/static";
    
    svr.Get("/static/(.*)", [static_dir](const httplib::Request& req, httplib::Response& res) {
        std::string path = req.matches[1];
        fs::path file_path = fs::path(static_dir) / path;
        
        std::string normalized_path = file_path.lexically_normal().string();
        std::replace(normalized_path.begin(), normalized_path.end(), '\\', '/');
        
        if (normalized_path.find(static_dir) != 0) {
            res.status = 403;
            return;
        }
        
        if (!fs::exists(file_path) || !fs::is_regular_file(file_path)) {
            res.status = 404;
            return;
        }
        
        std::string content = read_file_content(file_path.string());
        if (content.empty()) {
            res.status = 500;
            return;
        }
        
        res.set_content(content, get_mime_type(file_path.string()));
    });
}