#include "webview.h"
#include "src/func_whrite.h"
#include <stddef.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#ifdef _WIN32
#include <windows.h>
#include <shellscalingapi.h>

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

    webview_t w = webview_create(0, NULL);
    if (!w) {
        free(title);
        free(html);
        return -1;
    }

    SetupWebview(w, title, height, width, html);

    HWND hwnd = (HWND)webview_get_window(w);
    SetWindowIcon(hwnd, L"icon.ico");

    webview_run(w);
    webview_destroy(w);
    free(title);
    free(html);
    return 0;
}
#else
int main(void) {
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

    webview_t w = webview_create(0, NULL);
    if (!w) {
        free(title);
        free(html);
        return -1;
    }

    SetupWebview(w, title, height, width, html);

    webview_run(w);
    webview_destroy(w);
    free(title);
    free(html);
    return 0;
}
#endif