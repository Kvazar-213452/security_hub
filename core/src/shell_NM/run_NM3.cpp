#include <windows.h>
#include <string>
#include <vector>
#include <iostream>
#include "../../include/shell_NM/shared_vars.hpp"
#include "../../include/shell_NM/run_NM.hpp"

void run_NM3(int port) {
    std::string htmlContent = "http://localhost:" + std::to_string(port);
    
    std::vector<std::string> args = {
        "--module-path", "lib/javafx-sdk-17.0.14/lib",
        "--add-modules", "javafx.controls,javafx.web",
        "-jar", "target/webview-project-1.0-SNAPSHOT.jar",
        htmlContent
    };
    
    std::string cmdLine = "java";
    for (const auto& arg : args) {
        cmdLine += " " + arg;
    }

    STARTUPINFO si = { sizeof(si) };
    PROCESS_INFORMATION pi;
    
    if (!CreateProcess(
        NULL,
        cmdLine.data(),
        NULL,
        NULL,
        FALSE
        CREATE_NEW_CONSOLE,
        NULL,
        "../shell_NM/NM3",
        &si,
        &pi)) {
        std::cerr << "CreateProcess failed (" << GetLastError() << ")" << std::endl;
        g_shellWebRunning = false;
        return;
    }
    
    WaitForSingleObject(pi.hProcess, INFINITE);
    
    CloseHandle(pi.hProcess);
    CloseHandle(pi.hThread);
    
    g_shellWebRunning = false;
    g_serverRunning = false;
}