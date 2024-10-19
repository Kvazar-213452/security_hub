#include <windows.h>
#include <shlobj.h>

extern "C" __declspec(dllexport) int runDiskCleanup() {
    ShellExecute(NULL, "open", "cleanmgr.exe", "/sagerun:1", NULL, SW_SHOWNORMAL);
    return 0;
}

extern "C" __declspec(dllexport) int emptyRecycleBin() {
    HRESULT result = SHEmptyRecycleBin(NULL, NULL, SHERB_NOCONFIRMATION | SHERB_NOPROGRESSUI | SHERB_NOSOUND);
    
    if (SUCCEEDED(result)) {
        return 1;
    } else {
        return 0;
    }
}

extern "C" __declspec(dllexport) void cleanup() {
    emptyRecycleBin();
    runDiskCleanup();
}
