#include <iostream>
#include <windows.h>
#include <string>
#include <sstream>
#include <fstream>
#include <vector>
#include <stdexcept>

typedef int (*FindFreePortFunc)();

std::string generate_html_content(int port) {
    std::ostringstream url;
    url << "http://127.0.0.1:" << port << "/";

    std::string html_content_core = R"(
        <style>
            iframe{
                position: fixed;
                height: 100%;
                width: 100%;
                top: 0%;
                left: 0%;
            }
        </style>
        <iframe src=")" + url.str() + R"(" frameborder="0"></iframe>
    )";

    return html_content_core;
}
