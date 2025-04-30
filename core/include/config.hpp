#pragma once

#include <nlohmann/json.hpp>
#include <string>
#include <fstream>
#include <stdexcept>

using json = nlohmann::json;

class Config {
private:
    json data;
    std::string config_path = "data/config.json";

public:
    Config();
    
    template<typename T>
    T get(const std::string& key) const;
    
    bool has(const std::string& key) const;
    
    int getInt(const std::string& key) const;
    bool getBool(const std::string& key) const;
    std::string getString(const std::string& key) const;
};

extern bool debug;

// shell_NM
extern std::string name_app;
extern std::string x_scale;
extern std::string y_scale;

extern std::string NM1_phat;
extern std::string NM2_phat;
