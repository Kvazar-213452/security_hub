#include "webview.h"
#include <stddef.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <windows.h>
#include <shellscalingapi.h>

void SetupWebview(webview_t w, const char* title, int height, int width, const char* html) {
    webview_set_title(w, title);
    webview_set_size(w, width, height, WEBVIEW_HINT_NONE);
    webview_set_html(w, html);
}

int main(int argc, char *argv[]) {
    if (argc < 5) {
        printf("use: %s <title> <height> <width> <html>\n", argv[0]);
        return 1;
    }

    const char* title = argv[1];
    int height = atoi(argv[2]); 
    int width = atoi(argv[3]);
    const char* html = argv[4];

    webview_t w = webview_create(0, NULL);
    if (!w) {
        return -1;
    }

    SetupWebview(w, title, height, width, html);

    webview_run(w);
    webview_destroy(w);

    return 0;
}
