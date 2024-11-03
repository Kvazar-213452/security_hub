#ifndef FUNC_SHELL_H
#define FUNC_SHELL_H

#include <string>

std::string generate_html_content(int port);
int FindFreePort();
void deleteFile(const std::string& filePath);

#endif 