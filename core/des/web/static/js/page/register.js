// app_front_end/static/js/page/register.js

function change_module(id) {
    $('#module_1').hide();
    $('#module_3').hide();
    $('#module_2').hide();

    $('#' + id).show();
}

change_lang_now(0);

change_module("module_1");

function login_page() {
    change_module("module_3");
    change_lang_now(0);
}

function reg_page() {
    change_module("module_1");
    change_lang_now(0);
}

function send_data() {
    let data = {
        name: $('#register_name').val(),
        gmail: $('#register_gmail').val(),
        password: $('#register_password').val()
    };

    $.ajax({
        url: "/send_email",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function (response) {
            change_module("module_2");
        }
    });
}

function send_code() {
    let data = {
        code: $('#register_code').val(),
    };
    
    $.ajax({
        url: "/code_verefic",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function (response) {
            if (response == 1) {
                window.parent.postMessage("reload", "*");
            } else {
                if (lang_global === "uk") {
                    message_window('Невірний пароль');
                } else if (lang_global === "en") {
                    message_window('Invalid password');
                }
            }
        }
    });
}

function login() {
    let data = {
        name: $('#login_name').val(),
        password: $('#login_password').val()
    };
    
    $.ajax({
        url: "/login_acaunt",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function (response) {
            response = JSON.parse(response);

            if (response['status'] == 1) {
                login_acount(response);
            } else {
                if (lang_global === "uk") {
                    message_window('Невірний пароль');
                } else if (lang_global === "en") {
                    message_window('Invalid password');
                }
            }
        }
    });
}

function login_acount(response) {
    $.ajax({
        url: "/reg_file_unix",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(response),
        success: function (response) {
            window.parent.postMessage("reload", "*");
        }
    });
}
