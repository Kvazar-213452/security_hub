#include "../include/api_handler.hpp"
#include "../include/func_core.hpp"
#include <nlohmann/json.hpp>
#include <iostream>
#include <string>

using json = nlohmann::json;

void ApiHandler::get_json_file(const json& request, json& response) {
    json json_obj = json::parse(read_file(request["data"]));

    response["val"] = json_obj;
}

void ApiHandler::get_file(const json& request, json& response) {
    response["val"] = read_file(request["data"]);
}

void ApiHandler::get_module_for_render(const json& request, json& response) {
    std::string module_name = request["data"].get<std::string>();
    
    std::string config_path = "../module/" + module_name + "/config.json";
    std::string config_content = read_file(config_path);

    json json_obj = json::parse(config_content);

    std::string icon_path = "../module/" + module_name + "/icon.png";
    std::string icon_content = read_file(icon_path);
    
    response["data"] = json_obj;
    response["icon"] = base64_encode(icon_content);
}
