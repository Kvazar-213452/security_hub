#include "../include/api_handler.hpp"
#include "../include/func_core.hpp"
#include <nlohmann/json.hpp>
#include <iostream>

using json = nlohmann::json;

void ApiHandler::get_json_file(const json& request, json& response) {
    json json_obj = json::parse(read_file(request["data"]));

    response["val"] = json_obj;
}
