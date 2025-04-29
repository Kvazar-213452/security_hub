#ifndef FUNC_CORE_H
#define FUNC_CORE_H

#include <string>

std::string read_file(const std::string& relative_path);
std::string base64_encode(const std::string &input);
int FindFreePort();

#endif