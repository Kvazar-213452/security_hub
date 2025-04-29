#include "../include/routes.hpp"
#include "../lib/httplib.h"
#include "../include/static_handler.hpp"
#include "../include/api_handler.hpp"
#include <nlohmann/json.hpp>

void Routes::setup_routes(httplib::Server& svr, const std::string& base_dir) {
    svr.Get("/", [base_dir](const httplib::Request&, httplib::Response& res) {
        std::string html_content = StaticHandler::read_file_content(base_dir + "/index.html");
        if (html_content.empty()) {
            res.status = 500;
            res.set_content("Error: Could not read index.html", "text/plain");
        } else {
            res.set_content(html_content, "text/html");
        }
    });

    svr.Get("/none_module", [base_dir](const httplib::Request&, httplib::Response& res) {
        std::string html_content = StaticHandler::read_file_content(base_dir + "/none_module.html");
        if (html_content.empty()) {
            res.status = 500;
            res.set_content("Error: Could not read index.html", "text/plain");
        } else {
            res.set_content(html_content, "text/html");
        }
    });

    // api // api // api // api // api // api // api // api // api // api // api // api // api // api // api // api // api // api 
    // api // api // api // api // api // api // api // api // api // api // api // api // api // api // api // api // api // api 
    // api // api // api // api // api // api // api // api // api // api // api // api // api // api // api // api // api // api 

    svr.Post("/api/get_json_file", [](const httplib::Request& req, httplib::Response& res) {
        nlohmann::json request = nlohmann::json::parse(req.body);
        nlohmann::json response;
        
        ApiHandler::get_json_file(request, response);
        
        res.set_content(response.dump(), "application/json");
    });

    svr.Post("/api/get_file", [](const httplib::Request& req, httplib::Response& res) {
        nlohmann::json request = nlohmann::json::parse(req.body);
        nlohmann::json response;
        
        ApiHandler::get_file(request, response);
        
        res.set_content(response.dump(), "application/json");
    });

    svr.Post("/api/get_module_for_render", [](const httplib::Request& req, httplib::Response& res) {
        nlohmann::json request = nlohmann::json::parse(req.body);
        nlohmann::json response;
        
        ApiHandler::get_module_for_render(request, response);
        
        res.set_content(response.dump(), "application/json");
    });



    StaticHandler::setup_static_routes(svr, base_dir);
}