#include <windows.h>
#include <string>
#include <vector>
#include <iostream>
#include <atomic>
#include "../../include/shell_NM/run_NM.hpp"
#include "../../include/func_core.hpp"
#include "../../include/shell_NM/shared_vars.hpp"
#include "../../include/config.hpp"

void run_NM2(int port) {
    std::string htmlContent = std::to_string(port);

    int freePort = FindFreePort();
    std::string portStr = std::to_string(freePort);

    std::vector<std::string> args = {
        x_scale,
        y_scale,
        htmlContent,
        name_app,
        portStr
    };
    
    std::string cmdLine = NM2_phat;
    
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
        FALSE,
        0,
        NULL,
        NULL,
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