#pragma once

#include <string>
#include <nlohmann/json.hpp>

using json = nlohmann::json;

namespace ApiHandler {
    void get_json_file(const json& request, json& response);
    void get_file(const json& request, json& response);
    void get_module_for_render(const json& request, json& response);
}
