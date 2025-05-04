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

            if (style === 1) {
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
            window.parent.postMessage("reload", "*");
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
            if (lang_global === "uk") {
                message_window('Кеш видалено');
            } else if (lang_global === "en") {
                message_window('Cache deleted');
            }
        }
    });
}

function get_status_reg_settings() {
    $.ajax({
        url: "/get_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: "../data/user.json"}),
        success: function (response) {
            if (response == "") {
                $("#settings_1_btn_page2").hide();
            } else {
                let obj = JSON.parse(response);
                reg_login = obj["acsses"];
            }

            if (reg_login == 0) {
                $("#settings_1_btn_page2").hide();
            }
        }
    });
}

function get_data_reg() {
    $.ajax({
        url: "/get_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: "../data/user.json"}),
        success: function (response) {
            if (response != "") {
                let obj = JSON.parse(response);

                let text = `
                <p>name = <span>${obj['name']}</span></p>
                <p>pasw = <span>${obj['pasw']}</span></p>
                <p>gmail = <span>${obj['gmail']}</span></p>
                <p>acsses = <span>${obj['acsses']}</span></p>
                `;
    
                $("#settings13qwas").html(text);
            }
        }
    });
}

function get_all_render_module() {
    $.ajax({
        url: "/get_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: "../data/config_module.json"}),
        success: function (response) {
            let obj = JSON.parse(response);
            
            for (let i = 0; i < obj["module_install"].length; i++) {
                let text = `
                <button onclick="install_module_app('${obj["module_install"][i]}')">${obj["module_install"][i]}</button>
                `;

                $("#div_install_module").append(text);
            }

            for (let i = 0; i < obj["module_uinstall"].length; i++) {
                let text = `
                <button onclick="uinstall_module_app('${obj["module_uinstall"][i]}')">${obj["module_uinstall"][i]}</button>
                `;

                $("#div_uinstall_module").append(text);
            }

            get_render_module_satrt_now();
        }
    });
}

function get_render_module_satrt_now() {
    $.ajax({
        url: "/get_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: "result.json"}),
        success: function (response) {
            let obj = JSON.parse(response);
            const keys = Object.keys(obj["module_uinstall"]);
            
            for (let i = 0; i < keys.length; i++) {
                let text = `
                <button 
                onclick="info_start_modele_now('${obj["module_uinstall"][keys[i]]["pid"]}', '${obj["module_uinstall"][keys[i]]["port"]}', '${keys[i]}')">
                ${keys[i]}</button>
                `;

                $("#div_start_now_module").append(text);
            }
        }
    });
}

function info_start_modele_now(pid, port, name) {
    openModal("modal4");

    $("#fk2f224444_32r").html(pid);
    $("#fk2f224444_32r1").html(port);
    $("#fk2f224444_32r2").html(name);
}

function install_module_app(name) {
    openModal("modal2");
    $("#url_site_doc").html(url_module[name]);
    module_select = name;
    url_site = url_module[name];
}

function install_module_func() {
    clos("modal2");
    openModal("modal1");
    animation = true;

    $.ajax({
        url: "/install_module_app",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: module_select}),
        success: function (response) {
            setTimeout(function () {
                clos("modal1");
                animation = false;
            }, 2000);
        }
    });
}

function uinstall_module_app(name) {
    openModal("modal3");
    $("#url_site_doc1").html(`https://spx-security-hub.wuaze.com/doc/${name}.php`);
    module_select = name;
    url_site = url_module[name];
}

function uinstall_module_app_func() {
    openModal("modal1");
    animation = true;

    $.ajax({
        url: "/uinstall_module_app",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: module_select}),
        success: function (response) {
            setTimeout(function () {
                clos("modal1");
                animation = false;
            }, 2000);
        }
    });
}

function reload_model() {
    openModal("modal1");
    animation = true;

    $.ajax({
        url: "/reload_model",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            setTimeout(function () {
                clos("modal1");
                animation = false;
            }, 4000);
        }
    });
}
