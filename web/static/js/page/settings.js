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
            console.log(data);
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
            } else {
                button_active('shell_NM1', shell_NM);
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

        },
        error: function (xhr, status, error) {
            console.log("Error: " + error);
            console.log("Response text:", xhr.responseText);
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
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
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
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
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
            window.parent.postMessage("lang_change", "*");
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
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
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
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
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}
