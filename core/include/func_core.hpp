#pragma once

#include <string>

std::string read_file(const std::string& relative_path);
std::string base64_encode(const std::string &input);
int FindFreePort();
void write_starter_md(const std::string& content, const std::string& filePath = "starter.md");
