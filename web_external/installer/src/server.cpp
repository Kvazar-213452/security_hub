#include "../lib/httplib.h"
#include "../include/config.h"
#include "../include/html.h"

#include <iostream>

void start_server() {
    httplib::Server svr;

    auto html_content_ptr = std::make_shared<std::string>(html_content);

    svr.Get("/", [html_content_ptr](const httplib::Request& req, httplib::Response& res) {
        if (html_content.empty()) {
            res.status = 500;
            res.set_content("Error: Could not read index.html", "text/plain");
        } else {
            res.set_content(html_content, "text/html");
        }
    });

    svr.listen("127.0.0.1", port);
}