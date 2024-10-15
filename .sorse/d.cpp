#include <iostream>
#include <windows.h> // Для WinAPI

void runDiskCleanup() {
    std::cout << "Запуск очищення диска з усіма опціями..." << std::endl;
    // Запускаємо Disk Cleanup з параметрами
    ShellExecute(NULL, "open", "cleanmgr.exe", "/sagerun:1", NULL, SW_SHOWNORMAL);
}

int main() {
    runDiskCleanup();
    return 0;
}
