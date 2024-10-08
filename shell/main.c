#include "webview.h"
#include <stddef.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#ifdef _WIN32
#include <windows.h>
#include <shellscalingapi.h> // For High DPI support

char* read_name_from_file(const char* file_path) {
    FILE* file = fopen(file_path, "r");
    if (!file) {
        perror("Failed to open file");
        return NULL;
    }

    char line[256];
    while (fgets(line, sizeof(line), file)) {
        // Шукаємо рядок, що починається з "name ="
        if (strncmp(line, "name =", 6) == 0) {
            // Видаляємо пробіли на початку і кінці
            char* name = strtok(line + 6, "\r\n");
            fclose(file);
            return strdup(name); // Копіюємо значення в нову пам'ять
        }
    }

    fclose(file);
    return NULL; // Не знайдено
}

char* read_file_html(const char* file_path) {
    FILE* file = fopen(file_path, "r");
    if (!file) {
        perror("Failed to open file");
        return NULL;
    }

    char line[256];
    while (fgets(line, sizeof(line), file)) {
        // Шукаємо рядок, що починається з "html ="
        if (strncmp(line, "html =", 6) == 0) {
            // Видаляємо пробіли на початку і кінці
            char* html = strtok(line + 6, "\r\n");
            fclose(file);
            return strdup(html); // Копіюємо значення в нову пам'ять
        }
    }

    fclose(file);
    return NULL; // Не знайдено
}

int read_window_height(const char* file_path) {
    FILE* file = fopen(file_path, "r");
    if (!file) {
        perror("Failed to open file");
        return -1; // Помилка відкриття файлу
    }

    char line[256];
    while (fgets(line, sizeof(line), file)) {
        // Шукаємо рядок, що починається з "window_h ="
        if (strncmp(line, "window_h =", 10) == 0) {
            // Видаляємо пробіли на початку і кінці
            char* value_str = strtok(line + 10, "\r\n");
            if (value_str) {
                int height = atoi(value_str); // Перетворюємо рядок в ціле число
                fclose(file);
                return height;
            }
        }
    }

    fclose(file);
    return -1; // Не знайдено або помилка
}

int read_window_width(const char* file_path) {
    FILE* file = fopen(file_path, "r");
    if (!file) {
        perror("Failed to open file");
        return -1; // Помилка відкриття файлу
    }

    char line[256];
    while (fgets(line, sizeof(line), file)) {
        // Шукаємо рядок, що починається з "window_w ="
        if (strncmp(line, "window_w =", 10) == 0) {
            // Видаляємо пробіли на початку і кінці
            char* value_str = strtok(line + 10, "\r\n");
            if (value_str) {
                int width = atoi(value_str); // Перетворюємо рядок в ціле число
                fclose(file);
                return width;
            }
        }
    }

    fclose(file);
    return -1; // Не знайдено або помилка
}

// Function to set the icon
void SetWindowIcon(HWND hwnd, LPCWSTR iconPath) {
    HICON hIcon = (HICON)LoadImageW(NULL, iconPath, IMAGE_ICON, 0, 0, LR_LOADFROMFILE | LR_DEFAULTSIZE);
    if (hIcon) {
        SendMessage(hwnd, WM_SETICON, ICON_BIG, (LPARAM)hIcon);
        SendMessage(hwnd, WM_SETICON, ICON_SMALL, (LPARAM)hIcon);
    }
}
#endif

void SetupWebview(webview_t w, const char* title, int height, int width, const char* html) {
    webview_set_title(w, title);
    webview_set_size(w, width, height, WEBVIEW_HINT_NONE);
    webview_set_html(w, html);
}

#ifdef _WIN32
int WINAPI WinMain(HINSTANCE hInst, HINSTANCE hPrevInst, LPSTR lpCmdLine, int nCmdShow) {
    (void)hInst;
    (void)hPrevInst;
    (void)lpCmdLine;
    (void)nCmdShow;

    // Read the name, HTML, height, and width from the file
    char* title = read_name_from_file("start_conf.log");
    if (!title) {
        fprintf(stderr, "Failed to read name from file\n");
        return -1;
    }

    char* html = read_file_html("start_conf.log");
    if (!html) {
        fprintf(stderr, "Failed to read HTML from file\n");
        free(title);
        return -1;
    }

    int height = read_window_height("start_conf.log");
    if (height == -1) {
        fprintf(stderr, "Failed to read height from file\n");
        free(title);
        free(html);
        return -1;
    }

    int width = read_window_width("start_conf.log");
    if (width == -1) {
        fprintf(stderr, "Failed to read width from file\n");
        free(title);
        free(html);
        return -1;
    }

    // Initialize the webview
    webview_t w = webview_create(0, NULL);
    if (!w) {
        free(title);
        free(html);
        return -1;
    }

    SetupWebview(w, title, height, width, html);

    // Get the window handle and set the icon
    HWND hwnd = (HWND)webview_get_window(w);
    SetWindowIcon(hwnd, L"icon.ico"); // Update this path to your icon file

    // Run the webview
    webview_run(w);
    webview_destroy(w);
    free(title); // Clean up
    free(html); // Clean up
    return 0;
}
#else
int main(void) {
    // Read the name, HTML, height, and width from the file
    char* title = read_name_from_file("start_conf.log");
    if (!title) {
        fprintf(stderr, "Failed to read name from file\n");
        return -1;
    }

    char* html = read_file_html("start_conf.log");
    if (!html) {
        fprintf(stderr, "Failed to read HTML from file\n");
        free(title);
        return -1;
    }

    int height = read_window_height("start_conf.log");
    if (height == -1) {
        fprintf(stderr, "Failed to read height from file\n");
        free(title);
        free(html);
        return -1;
    }

    int width = read_window_width("start_conf.log");
    if (width == -1) {
        fprintf(stderr, "Failed to read width from file\n");
        free(title);
        free(html);
        return -1;
    }

    // Initialize the webview
    webview_t w = webview_create(0, NULL);
    if (!w) {
        free(title);
        free(html);
        return -1;
    }

    SetupWebview(w, title, height, width, html);

    // Run the webview
    webview_run(w);
    webview_destroy(w);
    free(title); // Clean up
    free(html); // Clean up
    return 0;
}
#endif