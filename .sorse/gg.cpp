#include <iostream>
#include <windows.h>
#include <shlobj.h> // Для SHEmptyRecycleBin

void emptyRecycleBin() {
    // Очищення Кошика
    HRESULT result = SHEmptyRecycleBin(NULL, NULL, SHERB_NOCONFIRMATION | SHERB_NOPROGRESSUI | SHERB_NOSOUND);
    
    if (SUCCEEDED(result)) {
        std::cout << "Кошик очищено успішно!" << std::endl;
    } else {
        std::cout << "Помилка при очищенні Кошика." << std::endl;
    }
}

int main() {
    emptyRecycleBin();
    return 0;
}
