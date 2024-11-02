#ifndef FUNC_SHELL_H
#define FUNC_SHELL_H

#include <string>

int port_find();
std::string generate_html_content(int port);
void save_base64_to_file(const std::string& base64_data, const std::string& filepath = "./FindFreePort.dll");

#endif 