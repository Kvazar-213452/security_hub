#include "../lib/httplib.h"
#include "../include/config.h"
#include "../include/html.h"

#include <iostream>
#include <filesystem>
#include <string>
#include <Windows.h>
#include <shlobj.h>
#include <minizip/unzip.h>
#include <zlib.h>
#include <bzlib.h>

void unzip(const std::string& zipFilePath, const std::string& destDir) {
    unzFile zipfile = unzOpen(zipFilePath.c_str());
    if (!zipfile) {
        std::cerr << "Не вдалося відкрити ZIP-файл: " << zipFilePath << std::endl;
        return;
    }

    std::filesystem::create_directories(destDir);

    char filename[256];
    unz_global_info globalInfo;
    unzGetGlobalInfo(zipfile, &globalInfo);

    for (uLong i = 0; i < globalInfo.number_entry; ++i) {
        if (unzGetCurrentFileInfo(zipfile, nullptr, filename, sizeof(filename), nullptr, 0, nullptr, 0) != UNZ_OK) {
            std::cerr << "Не вдалося отримати інформацію про файл." << std::endl;
            break;
        }

        std::string outputPath = destDir + "\\" + filename;

        if (filename[strlen(filename) - 1] == '/') {
            std::filesystem::create_directories(outputPath);
        } else {
            std::filesystem::create_directories(std::filesystem::path(outputPath).parent_path());

            if (unzOpenCurrentFile(zipfile) != UNZ_OK) {
                std::cerr << "Не вдалося відкрити файл у ZIP: " << filename << std::endl;
                break;
            }

            FILE* outFile = fopen(outputPath.c_str(), "wb");
            if (!outFile) {
                std::cerr << "Не вдалося створити файл: " << outputPath << std::endl;
                unzCloseCurrentFile(zipfile);
                break;
            }

            char buffer[8192];
            int bytesRead;
            while ((bytesRead = unzReadCurrentFile(zipfile, buffer, sizeof(buffer))) > 0) {
                fwrite(buffer, sizeof(char), bytesRead, outFile);
            }

            fclose(outFile);
            unzCloseCurrentFile(zipfile);
        }

        if ((i + 1) < globalInfo.number_entry) {
            unzGoToNextFile(zipfile);
        }
    }

    unzClose(zipfile);
    std::cout << "Розпаковка завершена у " << destDir << std::endl;
}

const IID IID_IShellLinkW = {0x000214F9, 0x0000, 0x0000, {0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}};
const IID IID_IPersistFile = {0x0000010B, 0x0000, 0x0000, {0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}};

void CreateShortcut(const std::wstring& shortcutPath, const std::wstring& targetPath, const std::wstring& workingDir) {
    CoInitialize(NULL);

    IShellLinkW* psl;
    if (SUCCEEDED(CoCreateInstance(CLSID_ShellLink, NULL, CLSCTX_INPROC_SERVER, IID_IShellLinkW, (LPVOID*)&psl))) {
        psl->SetPath(targetPath.c_str());
        psl->SetWorkingDirectory(workingDir.c_str());

        IPersistFile* ppf;
        if (SUCCEEDED(psl->QueryInterface(IID_IPersistFile, (LPVOID*)&ppf))) {
            ppf->Save(shortcutPath.c_str(), TRUE);
            ppf->Release();
        }
        psl->Release();
    }

    CoUninitialize();
}

void deleteFile(const std::string& filePath) {
    try {
        if (std::filesystem::exists(filePath)) {
            std::filesystem::remove(filePath);
            std::cout << "Файл '" << filePath << "' було успішно видалено." << std::endl;
        } else {
            std::cout << "Файл '" << filePath << "' не знайдено." << std::endl;
        }
    } catch (const std::filesystem::filesystem_error& e) {
        std::cerr << "Помилка: " << e.what() << std::endl;
    }
}

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

            const char* command = "curl -L -o main.zip https://raw.githubusercontent.com/Kvazar-213452/security_hub/refs/heads/main/web_external/data/main.zip";

            int result = system(command);

            if (result == 0) {
                std::cout << "Файл main.zip завантажено успішно." << std::endl;
            } else {
                std::cerr << "Помилка при завантаженні файлу main.zip." << std::endl;
            }

            std::string outputDir = "C:\\security_hub";
            std::filesystem::create_directories(outputDir);

            unzip("main.zip", outputDir);

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
        }
    });

    svr.listen("127.0.0.1", port);
}
