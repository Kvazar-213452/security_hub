#include <iostream>
#include <thread>
#include <chrono>
#include <atomic>
#include "lib/webview.h"
#include "include/httplib.h"
#include "include/html.h"
#include "include/config.h"
#include "include/server.h"
#include <string>
#include <sstream>

std::atomic<bool> webview_closed(false);

void start_webview() {
    try {
        webview::webview w(false, nullptr);
        w.set_title(name_app);
        w.set_size(window_h, window_w, WEBVIEW_HINT_NONE);
        w.set_html(html_content_core);
        w.run();
        webview_closed.store(true); // Set the flag when the webview is closed
    } catch (const webview::exception &e) {
        std::cerr << e.what() << std::endl;
        exit(1);
    }
}

void monitor_webview() {
    while (!webview_closed.load()) {
        std::this_thread::sleep_for(std::chrono::milliseconds(100));
    }
    // Terminate the program or perform any cleanup
    std::exit(0); // Exits the entire process
}

#ifdef _WIN32
int WINAPI WinMain(HINSTANCE /*hInst*/, HINSTANCE /*hPrevInst*/,
                   LPSTR /*lpCmdLine*/, int /*nCmdShow*/) {
    std::thread server_thread(start_server);
    std::this_thread::sleep_for(std::chrono::seconds(1));

    std::thread monitor_thread(monitor_webview);
    start_webview();
    monitor_thread.join(); // Ensure the monitor thread completes

    server_thread.join();
    return 0;
}
#else
int main() {
    std::thread server_thread(start_server);
    std::this_thread::sleep_for(std::chrono::seconds(1));

    std::thread monitor_thread(monitor_webview);
    start_webview();
    monitor_thread.join(); // Ensure the monitor thread completes

    server_thread.join();
    return 0;
}
#endif


