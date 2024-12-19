#include "lib/webview.h"
#include "lib/httplib.h"
#include "include/html.h"
#include "include/config.h"
#include "include/server.h"
#include "include/func_shell.h"

#include <iostream>
#include <thread>
#include <chrono>
#include <atomic>
#include <string>
#include <sstream>

std::atomic<bool> webview_closed(false);

void start_webview(int port) {
    try {
        std::string html_content_core = generate_html_content(port);

        webview::webview w(false, nullptr);
        w.set_title(name_app);
        w.set_size(window_h, window_w, WEBVIEW_HINT_FIXED);
        w.set_html(html_content_core);
        w.run();
        webview_closed.store(true);
    } catch (const webview::exception &e) {
        std::cerr << e.what() << std::endl;
        exit(1);
    }
}

void monitor_webview() {
    while (!webview_closed.load()) {
        std::this_thread::sleep_for(std::chrono::milliseconds(100));
    }
    std::exit(0);
}

int WINAPI WinMain(HINSTANCE, HINSTANCE, LPSTR, int) {
    int port = FindFreePort();
    std::cerr << port << std::endl;

    std::thread server_thread(std::bind(start_server, port));

    std::this_thread::sleep_for(std::chrono::seconds(1));

    std::thread monitor_thread(monitor_webview);
    start_webview(port);
    monitor_thread.join(); 

    server_thread.join();
    return 0;
}