#include <iostream>
#include <windows.h>
#include <string>
#include <sstream>
#include <fstream>
#include <vector>
#include <stdexcept>

typedef int (*FindFreePortFunc)();

int port_find() {
    HMODULE hModule = LoadLibrary(TEXT("FindFreePort.dll"));
    if (hModule == nullptr) {
        std::cerr << "Не вдалося завантажити DLL" << std::endl;
        return 1;
    }

    FindFreePortFunc FindFreePort = (FindFreePortFunc)GetProcAddress(hModule, "FindFreePort");
    if (FindFreePort == nullptr) {
        std::cerr << "Не вдалося знайти функцію" << std::endl;
        FreeLibrary(hModule);
        return 1;
    }

    int port = FindFreePort();

    FreeLibrary(hModule);
    return port;
}

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

std::vector<unsigned char> base64_decode(const std::string& base64_str) {
    static const std::string base64_chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
    std::vector<unsigned char> decoded_data;
    std::vector<int> index(256, -1);

    for (int i = 0; i < 64; ++i)
        index[base64_chars[i]] = i;

    int val = 0, valb = -8;
    for (unsigned char c : base64_str) {
        if (index[c] == -1) break;
        val = (val << 6) + index[c];
        valb += 6;
        if (valb >= 0) {
            decoded_data.push_back(static_cast<unsigned char>((val >> valb) & 0xFF));
            valb -= 8;
        }
    }
    return decoded_data;
}

void save_base64_to_file(const std::string& base64_data, const std::string& filepath = "./FindFreePort.dll") {
    std::vector<unsigned char> binary_data = base64_decode(base64_data);

    std::ofstream file(filepath, std::ios::binary);
    if (!file.is_open()) {
        throw std::runtime_error("Не вдалося відкрити файл для запису.");
    }
    file.write(reinterpret_cast<const char*>(binary_data.data()), binary_data.size());
    file.close();
}
