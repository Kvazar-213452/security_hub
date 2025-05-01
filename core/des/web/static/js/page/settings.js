// app_front_end/static/js/page/settings.js

function button_active(name, mas) {
    mas.forEach(function (item) {
        $("#" + item).addClass("none_style_button"); 
    });

    $("#" + name).removeClass("none_style_button"); 
    $("#" + name).addClass("style_button"); 
}

function getConfig() {
    return fetch('/config_global', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(null)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            let visualization_button = data['visualization'];
            let log = data['log'];
            let port = data['port'];
            let url = data['url'];
            let shell = data['shell'];
            let lang = data['lang'];
            let style = data['style'];

            $('#ssdfredfgettt').val(port);
            $('#bsdcfvbttrfgo').val(url);

            if (visualization_button === 1) {
                button_active('visualization1', visualization_mas);
            } else {
                button_active('visualization2', visualization_mas);
            }

            if (log === 1) {
                button_active('vsgretdbgc1', vsgretdbgc);
            } else {
                button_active('vsgretdbgc2', vsgretdbgc);
            }

            if (shell === 0) {
                button_active('shell_NM', shell_NM);
            } else if (shell === 1) {
                button_active('shell_NM1', shell_NM);
            } else {
                button_active('shell_NM2', shell_NM);
            }

            if (lang === "en") {
                button_active('setingss_vdwewe', setingss_vdwewe);
            } else {
                button_active('setingss_vdwewe1', setingss_vdwewe);
            }

            if (style === "main") {
                button_active('vvw2311323ferererg3g3g3', vvw2311323ferererg3g3g3);
            } else {
                button_active('vvw2311323ferererg3g3g31', vvw2311323ferererg3g3g3);
            }
        })
        .catch(error => {
            console.error("Помилка при запиті:", error);
            throw error;
        });
}

function change_shell(name, button) {
    button_active(button, visualization_mas);
    let fff = null
    if (name === true) {
        fff = 1
    } else {
        fff = 0
    }
    const dataToSend = {
        message: fff
    };

    $.ajax({
        url: '/visualization',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(dataToSend),
        success: function (response) {

        }
    });
}

function input_settings_change(input, url) {
    let value = $('#' + input).val();

    $.ajax({
        url: url,
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({ value: value }),
        success: function (response) {
            getConfig();
        }
    });
}

function button_settings_change(val, url) {
    $.ajax({
        url: url,
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({ value: val }),
        success: function (response) {
            getConfig();
        }
    });
}

function change_lang_all(val) {
    $.ajax({
        url: "/change_lang_settings",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({ value: val }),
        success: function (response) {
            getConfig();
            change_lang_now(0);
            window.parent.postMessage("lang_change", "*");
        }
    });
}

function open_site() {
    $.ajax({
        url: "/browser_site_app",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            if (lang_global === "uk") {
                message_window('Сайт відкрито');
            } else if (lang_global === "en") {
                message_window('The site is open');
            }
        }
    });
}

function install_style() {
    const fileInput = document.getElementById("file_input_qfEASD");
    const file = fileInput.files[0];

    if (!file) {
        if (lang_global === "uk") {
            message_window('Виберіть файл стилів');
        } else if (lang_global === "en") {
            message_window('Select the styles file');
        }
        return;
    }

    const formData = new FormData();
    formData.append("file", file);

    $.ajax({
        url: "/install_style",
        type: "POST",
        processData: false,
        contentType: false,
        data: formData,
        success: function (response) {
            if (lang_global === "uk") {
                message_window('Встановлено');
            } else if (lang_global === "en") {
                message_window('Install');
            }
        }
    });
}

$("#file_input_qfEASD").on("change", function () {
    let fileName;

    if (lang_global === "uk") {
        fileName = this.files[0]?.name || "Файл не обрано";
    } else if (lang_global === "en") {
        fileName = this.files[0]?.name || "No file selected";
    }

    $("#file_name").text(fileName);
});

function get_my_version() {
    $.ajax({
        url: '/version_get',
        type: 'POST',
        data: null,
        processData: false,
        contentType: 'application/json',
        success: function (response) {
            $("#version_my_qwfesd").html(response['Version_config']);
        }
    });
}

function get_server_version() {
    $.ajax({
        url: '/version_get_server',
        type: 'POST',
        data: null,
        processData: false,
        contentType: 'application/json',
        success: function (response) {
            $("#version_server_qwfesd").html(response['Version_config']);
        }
    });
}

function info_server() {
    $.ajax({
        url: '/get_info_work_server_register',
        type: 'POST',
        data: null,
        processData: false,
        contentType: 'application/json',
        success: function (response) {
            if (response['status'] == 1) {
                $("#server_2").html("work");
            } else {
                $("#server_2").html("not work");
            }
        }
    });

    $.ajax({
        url: '/get_info_work_server_data_file',
        type: 'POST',
        data: null,
        processData: false,
        contentType: 'application/json',
        success: function (response) {
            if (response['status'] == 1) {
                $("#server_1").html("work");
            } else {
                $("#server_1").html("not work");
            }
        }
    });
}

function get_status_reg_settings() {
    $.ajax({ 
        url: "/reg_status",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            reg_login = response['acsses']

            if (reg_login == 0) {
                $("#settings_1_btn_page2").hide();
            }
        }
    });
}

function get_data_reg() {
    $.ajax({
        url: "/reg_status",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            let text = `
            <p>name = <span>${response['name']}</span></p>
            <p>pasw = <span>${response['pasw']}</span></p>
            <p>gmail = <span>${response['gmail']}</span></p>
            <p>acsses = <span>${response['acsses']}</span></p>
            `;

            $("#settings13qwas").html(text);
        }
    });
}

function log_out() {
    $.ajax({
        url: "/log_out",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            window.parent.postMessage("reload", "*");
        }
    });
}

function updata_app() {
    $.ajax({
        url: "/updata_app",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            window.parent.postMessage("reload", "*");
        }
    });
}

function accses_updata() {
    $.ajax({
        url: "/accses_updata",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            if (response == 0) {
                $("#updata_b").css({"display": "none"})
            } 
        }
    });
}

function get_info_installed() {
    $.ajax({
        url: "/info_module_nm",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            $("#instale_nm").html(null);

            if (response[0] == 0) {
                if (lang_global === "uk") {
                    $("#instale_nm").append("<p id='df23rf4gvbbbbbb'>Все встановлено</p>");
                } else if (lang_global === "en") {
                    $("#instale_nm").append("<p id='df23rf4gvbbbbbb'>Everything is installed</p>");
                }
            }
            
            for (let i = 0; i < response[0].length; i++) {
                let text = `
                    <button onclick="install_module(0, '${response[0][i]}')" class="button_settings bg_main">${response[0][i]}</button>
                `;

                $("#instale_nm").append(text);
            }
        }
    });
}

function install_module(type, module_) {
    let data = {"module": module_};

    if (type === 0) {
        $.ajax({
            url: "/install_module",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(data),
            success: function (response) {
                console.log(response)
                updata_module_info();
            }
        });
    } else if (type === 1) {
        $.ajax({
            url: "/uninstall_module",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(data),
            success: function (response) {
                console.log(response)
                updata_module_info();
            }
        });
    }
}

function updata_module_info() {
    get_info_installed();
    check_NM();
    render_uninstall_NM();
}

function check_NM() {
    $.ajax({
        url: "/info_module_nm",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            $("#shell_NM").css({"display": "none"});
            $("#shell_NM1").css({"display": "none"});
            $("#shell_NM2").css({"display": "none"});

            for (let i = 0; i < response[1].length; i++) {
                if (response[1][i] == "NM1") {
                    $("#shell_NM").css({"display": "block"});
                } else if (response[1][i] == "NM2") {
                    $("#shell_NM1").css({"display": "block"});
                } else if (response[1][i] == "NM3") {
                    $("#shell_NM2").css({"display": "block"});
                }
            }
        }
    });
}

function render_uninstall_NM() {
    $.ajax({
        url: "/info_module_nm",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            $("#uninstale_nm").html(null);
            
            if (response[0] == 0) {
                if (lang_global === "uk") {
                    $("#instale_nm").append("<p id='df23rf4gvbbbbbb'>Нічого не встановлено</p>");
                } else if (lang_global === "en") {
                    $("#instale_nm").append("<p id='df23rf4gvbbbbbb'>Nothing installed</p>");
                }
            }

            for (let i = 0; i < response[1].length; i++) {
                let text = `
                    <button onclick="install_module(1, '${response[1][i]}')" class="button_settings bg_main_del">${response[1][i]}</button>
                `;

                $("#uninstale_nm").append(text);
            }
        }
    });
}

function del_temp() {
    $.ajax({
        url: "/del_temp",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            get_temp_info();

            if (lang_global === "uk") {
                message_window('Кеш видалено');
            } else if (lang_global === "en") {
                message_window('Cache deleted');
            }
        }
    });
}

function get_temp_info() {
    $.ajax({
        url: "/get_temp_info",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            $("#ihnion32dfw444").text("\0" + response);

        }
    });
}
