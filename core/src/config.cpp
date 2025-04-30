#include "../include/config.hpp"
#include <string>
#include <iostream>
#include <fstream>
#include <filesystem>
#include <algorithm>

bool debug = true;

// shell_NM
std::string name_app = "security_hub";
std::string x_scale = "1000";
std::string y_scale = "800";

std::string NM1_phat = "../shell_NM/NM1/shell_web.exe";
std::string NM2_phat = "../shell_NM/NM2/main.exe";
std::string NM3_phat = "../shell_NM/NM1/shell_web.exe";

Config::Config() {
    try {
        std::ifstream config_file(config_path);
        if (!config_file.is_open()) {
            throw std::runtime_error("Could not open config file: " + config_path);
        }
        config_file >> data;
    } catch (const std::exception& e) {
        std::cerr << "Config error: " << e.what() << std::endl;
        data = {};
    }
}

template<typename T>
T Config::get(const std::string& key) const {
    if (!data.contains(key)) {
        throw std::runtime_error("Config key not found: " + key);
    }
    return data[key].get<T>();
}

bool Config::has(const std::string& key) const {
    return data.contains(key);
}

int Config::getInt(const std::string& key) const {
    return get<int>(key);
}

bool Config::getBool(const std::string& key) const {
    return get<bool>(key);
}

std::string Config::getString(const std::string& key) const {
    return get<std::string>(key);
}