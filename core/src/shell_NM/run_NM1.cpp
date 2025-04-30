#include "../../include/shell_NM/run_NM.hpp"
#include "../../include/shell_NM/shared_vars.hpp"
#include "../../include/config.hpp"

void run_NM1(int port) {
    std::string htmlContent = std::to_string(port);
    
    std::vector<std::string> args = {name_app, y_scale, x_scale, htmlContent};
    std::string cmdLine = NM1_phat;
    
    for (const auto& arg : args) {
        cmdLine += " " + arg;
    }
    
    STARTUPINFO si = { sizeof(si) };
    PROCESS_INFORMATION pi;
    
    if (!CreateProcess(
        NULL,
        &cmdLine[0],
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