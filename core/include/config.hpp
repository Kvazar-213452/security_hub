#ifndef CONFIG_H
#define CONFIG_H

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
    
    // Явні інстанціювання для основних типів
    int getInt(const std::string& key) const;
    bool getBool(const std::string& key) const;
    std::string getString(const std::string& key) const;
};

extern bool debug;

#endif // CONFIG_H