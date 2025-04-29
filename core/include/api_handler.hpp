#ifndef API_HANDLER_H
#define API_HANDLER_H

#include <string>
#include <nlohmann/json.hpp>

using json = nlohmann::json;

namespace ApiHandler {
    void get_json_file(const json& request, json& response);
}

#endif