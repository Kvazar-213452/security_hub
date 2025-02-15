#include <iostream>
#include <filesystem>
#include <string>
#include <Windows.h>
#include <shlobj.h>
#include <sstream>
#include <fstream>
#include <cstdlib>

namespace fs = std::filesystem;

typedef const char* (*UnzipFunc)(const char*, const char*);

void runCommandInBackground(const char* command) {
    std::system(command);
}

bool fileExists(const std::string& filePath) {
    return fs::exists(filePath);
}

void copyFileToDirectory(const std::string& sourceFile, const std::string& targetDir) {
    try {
        if (!fileExists(sourceFile)) {
            std::cerr << "File " << sourceFile << " does not exist!" << std::endl;
            return;
        }

        if (!fs::exists(targetDir)) {
            fs::create_directories(targetDir);
            std::cout << "Directory created: " << targetDir << std::endl;
        }

        std::string targetFile = targetDir + "\\head.exe";

        fs::copy(sourceFile, targetFile, fs::copy_options::overwrite_existing);
        std::cout << "File copied to: " << targetFile << std::endl;
    }
    catch (const std::exception& e) {
        std::cerr << "Error: " << e.what() << std::endl;
    }
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
