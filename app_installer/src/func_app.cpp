#include <iostream>
#include <filesystem>
#include <string>
#include <Windows.h>
#include <shlobj.h>
#include <minizip/unzip.h>
#include <zlib.h>
#include <bzlib.h>

typedef const char* (*UnzipFunc)(const char*, const char*);

void unzip(const std::string& zipFilePath, const std::string& destDir) {
    unzFile zipfile = unzOpen(zipFilePath.c_str());
    if (!zipfile) {
        std::cerr << "Не вдалося відкрити ZIP-файл: " << zipFilePath << std::endl;
        return;
    }

    std::filesystem::create_directories(destDir);

    char filename[256];
    unz_global_info globalInfo;
    if (unzGetGlobalInfo(zipfile, &globalInfo) != UNZ_OK) {
        std::cerr << "Не вдалося отримати глобальну інформацію ZIP-файлу." << std::endl;
        unzClose(zipfile);
        return;
    }

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

            const size_t bufferSize = 285536;
            char buffer[bufferSize];
            int bytesRead;
            while ((bytesRead = unzReadCurrentFile(zipfile, buffer, sizeof(buffer))) > 0) {
                if (fwrite(buffer, sizeof(char), bytesRead, outFile) != bytesRead) {
                    std::cerr << "Помилка запису до файлу: " << outputPath << std::endl;
                    fclose(outFile);
                    unzCloseCurrentFile(zipfile);
                    return;
                }
            }

            if (bytesRead < 0) {
                std::cerr << "Помилка читання файлу з ZIP: " << filename << std::endl;
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

void runCommandInBackground(const char* command) {
    STARTUPINFO si;
    PROCESS_INFORMATION pi;
    ZeroMemory(&si, sizeof(si));
    si.cb = sizeof(si);
    si.dwFlags |= STARTF_USESHOWWINDOW;
    si.wShowWindow = SW_HIDE;

    ZeroMemory(&pi, sizeof(pi));

    if (!CreateProcess(NULL, const_cast<char*>(command), NULL, NULL, FALSE, CREATE_NO_WINDOW, NULL, NULL, &si, &pi)) {
        std::cerr << "Помилка при виконанні команди." << std::endl;
        return;
    }

    CloseHandle(pi.hProcess);
    CloseHandle(pi.hThread);
}
