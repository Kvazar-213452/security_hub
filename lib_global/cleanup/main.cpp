#include <windows.h> // Для WinAPI
#include <shlobj.h> // Для SHEmptyRecycleBin

extern "C" __declspec(dllexport) int runDiskCleanup() {
    // Запускаємо Disk Cleanup з параметрами
    ShellExecute(NULL, "open", "cleanmgr.exe", "/sagerun:1", NULL, SW_SHOWNORMAL);
    return 0; // Повертаємо 0 для успішного виконання
}

extern "C" __declspec(dllexport) int emptyRecycleBin() {
    // Очищення Кошика
    HRESULT result = SHEmptyRecycleBin(NULL, NULL, SHERB_NOCONFIRMATION | SHERB_NOPROGRESSUI | SHERB_NOSOUND);
    
    if (SUCCEEDED(result)) {
        return 1; // Повертаємо 1 для успішного очищення Кошика
    } else {
        return 0; // Повертаємо 0 для помилки
    }
}

extern "C" __declspec(dllexport) void cleanup() {
    emptyRecycleBin(); // Очищуємо Кошик
    runDiskCleanup(); // Запускаємо очищення диска
}
