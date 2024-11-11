function clos(name) {
    $('#' + name).hide(); 
}

function openModal(name) {
    $('#' + name).show(); 
}

function button_hover(name) {
    $("#" + name).css("opacity", "1");
    $("#" + name).css("background-color", "#565574");
}

function fetchLogs() {
    $.post('/get_logs', function (data) {
        const logsArray = data.log.split('\n');

        $('.console').html(logsArray.join('<br>'));
    });
}

function console_open() {
    fetchLogs();
    $('.console').toggle();
}

function button_active(name, mas) {
    mas.forEach(function (item) {
        $("#" + item).css({
            "border": "none",
            "color": "#ffffffd4"
        });
    });

    $("#" + name).css({
        "color": "#766aff"
    });
}

function get_data_wifi_now() {
    $.ajax({
        url: "/get_wifi_now",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            console.log(response)
            data_wifi_render_now(response)
            checkUnsafeProtocols();
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function get_data_wifi_all() {
    $.ajax({
        url: "/get_wifi",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            get_network_now(function (ssid) {
                render_all_network_wifi(response, ssid);
            });
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function get_data_for_schedule() {
    $.ajax({
        url: "/network_now",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            let wifi_now = response['ssid'];

            const now = new Date();

            const hours = now.getHours();
            const minutes = now.getMinutes();
            const seconds = now.getSeconds();

            let data = `${hours}:${minutes}:${seconds}`;

            get_wifi_info_level(data, wifi_now)
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function get_wifi_info_level(data, wifi_now) {
    $.ajax({
        url: "/get_wifi",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            let i = 0

            while (i < response.length) {
                wifi_now = wifi_now.replace(/^"|"$/g, '');
                if (wifi_now === response[i]['ssid']) {
                    schedule_render(data, response[i]['signal']);
                }
                i++;
            }
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function data_wifi_render_now(response) {
    $('#fweetrfgcvweee').html(null);

    for (let key in response) {

        if (response.hasOwnProperty(key)) {
            let formattedKey = key.charAt(0).toUpperCase() + key.slice(1);

            if (response[key] == "") {
                let text = `
                <div class="div_info_os">
                    <p class="name_o">${formattedKey}</p>
                    <p class="desc_o">N/A</p>
                    <div class="hr_div"></div>
                </div>
                `;

                $('#fweetrfgcvweee').append(text);
            } else {
                let text = `
                <div class="div_info_os">
                    <p class="name_o">${formattedKey}</p>
                    <p class="desc_o">${response[key]}</p>
                    <div class="hr_div"></div>
                </div>
                `;

                $('#fweetrfgcvweee').append(text);
            }
        }
    }
}

function checkUnsafeProtocols() {
    const text = $("#ciphers_wifi").text();

    if (text) {
        for (const protocol of unsafeProtocols) {
            if (text.includes(protocol)) {
                if (lang_global === "uk") {
                    $("#wifi_protection").html(`<p class="wifi_3_div_red">Незахищено</p>`);
                } else if (lang_global === "en") {
                    $("#wifi_protection").html(`<p class="wifi_3_div_red">Not protected</p>`);
                }
                break;
            }
        }
    }

    if (lang_global === "uk") {
        $("#wifi_protection").html(`<p class="wifi_3_div">Захищено</p>`);
    } else if (lang_global === "en") {
        $("#wifi_protection").html(`<p class="wifi_3_div">Protected</p>`);
    }
}

function render_all_network_wifi(response, ssid) {
    $('#render_all_wifi').html(null);

    for (let i = 0; i < response.length; i++) {
        for (let r = i + 1; r < response.length; r++) {
            if (response[i]['ssid'] === response[r]['ssid']) {
                response.splice(r, 1);
            }
        }
    }

    for (let i = 0; i < response.length; i++) {
        let text = `<div class="div_wifi_all">
        <p class="name_wifi_div_all">${response[i]['ssid']}</p>
        <p class="right_left_signal">${response[i]['signal']}</p>
        </div>`;
        $('#render_all_wifi').append(text);
    }

    $('.name_wifi_div_all').each(function () {
        if ($(this).text() === ssid) {
            $(this).closest('.div_wifi_all').addClass("border_active_grran");
        }
    });
}

function get_network_now(callback) {
    $.ajax({
        url: "/network_now",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            if (callback) {
                callback(response['ssid']);
            }
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
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

function get_os_data() {
    $.ajax({
        url: "/get_os_data",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            write_os_data(response)
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function window_open() {
    $.ajax({
        url: "/window_open",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            render_data_window_open(response);
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function write_os_data(response) {
    let data = response['data'];
    let jsonData = JSON.parse(data);

    $('#ddcbnxcew33333').html(null);
    $('#ndwe8rfier').html(null);
    $('#bfgtey65yt').html(null);

    console.log(jsonData);

    let text = `
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Операційна система" : (lang_global === "en" ? "Operating System" : "")}</p>
            <p class="desc_o">${jsonData['OS']['Name']}</p>
            <p class="desc_o_1">${jsonData['OS']['Version']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Архітектура" : (lang_global === "en" ? "Architecture" : "")}</p>
            <p class="desc_o">${jsonData['Architecture']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Диск" : (lang_global === "en" ? "Disk" : "")}</p>
            <p class="desc_o">${jsonData['Disk']['FreeSpace']}</p>
            <p class="desc_o_1">${jsonData['Disk']['TotalSpace']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Память" : (lang_global === "en" ? "Memory" : "")}</p>
            <p class="desc_o">${jsonData['Memory']['FreeMemory']}</p>
            <p class="desc_o_1">${jsonData['Memory']['FreeVirtualMemory']}</p>
            <p class="desc_o_2">${jsonData['Memory']['TotalMemory']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Кількість процесорів" : (lang_global === "en" ? "Processor сount" : "")}</p>
            <p class="desc_o">${jsonData['ProcessorCount']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Час роботи системи" : (lang_global === "en" ? "System uptime" : "")}</p>
            <p class="desc_o">${jsonData['SystemUptime']['Days']}:${jsonData['SystemUptime']['Hours']}:${jsonData['SystemUptime']['Minutes']}:${jsonData['SystemUptime']['Seconds']}</p>
            <div class="hr_div"></div>
        </div>
    `;

    $('#ddcbnxcew33333').append(text);

    for (let i = 1; i < jsonData['LoadedLibraries']['Libraries'].length; i++) {
        let data_text = `<p>${jsonData['LoadedLibraries']['Libraries'][i]}</p>`;

        $('#ndwe8rfier').append(data_text);
    }

    for (let i = 1; i < jsonData['NetworkAdapters']['Adapters'].length; i++) {
        let data_text = `
        <div class="div_info_os">
            <p class="name_o">${jsonData['NetworkAdapters']['Adapters'][i]['Description']}</p>
            <p class="desc_o_3">${jsonData['NetworkAdapters']['Adapters'][i]['IPAddress']}</p>
            <div class="hr_div"></div>
        </div>
        `;

        $('#bfgtey65yt').append(data_text);
    }
}

function render_data_window_open(response) {
    $('#sifewfewx').html(null);
    response = response['devices']

    for (let i = 0; i < response.length; i++) {
        if (i + 1 === response.length) {
            //pass
        } else {
            let text = `<div class="div_wifi_all">
            <p class="name_wifi_div_all">${response[i]}</p>
            </div>`;
            $('#sifewfewx').append(text);
        }
    }
}

function resource_info() {
    $.ajax({
        url: "/resource_info",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            resource_info_render_data(response);
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function resource_info_render_data(response) {
    response = response['data'];

    $('#dwqwdwfcfff44').text(response[0] || 'N/A');
    $('#rggwiovnewcee').text(response[1] || 'N/A');
}

function message_window(content) {
    const block = document.createElement("div");
    block.className = "animatedBlock hide";
    block.textContent = content;
    document.body.appendChild(block);

    setTimeout(() => {
        block.classList.remove("hide");
        block.classList.add("show");
    }, 0);

    setTimeout(() => {
        block.classList.remove("show");
        block.classList.add("hide");

        setTimeout(() => {
            block.remove();
        }, 1000);
    }, 3000);
}

function antivirus_web_start() {
    let inputValue = $('#fkwe9203f').val();
    const dataToSend = {
        url_site: [
            inputValue
        ]
    };

    $('#dwdefw4f4').text('Перевірка.....');

    $.ajax({
        url: '/antivirus_web',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(dataToSend),
        success: function (response) {
            antivirus_web_end(response)
        },
        error: function (xhr, status, error) {
            console.log("Error: " + error);
            console.log("Response text:", xhr.responseText);
        }
    });
}

function antivirus_web_end(response) {
    $('#dqdcew336g').show();

    if (lang_global === "uk") {
        $('#dwdefw4f4').text('Завершено'); 
    } else if (lang_global === "en") {
        $('#dwdefw4f4').text('Completed'); 
    }

    if (response['found'] === false) {
        if (lang_global === "uk") {
            $('#dw93244444').text('Сайт безпечний');
        } else if (lang_global === "en") {
            $('#dw93244444').text('The site is safe');
        }
    } else {
        if (lang_global === "uk") {
            $('#dw93244444').text('Сайт небезпечний');
        } else if (lang_global === "en") {
            $('#dw93244444').text('The site is dangerous');
        }
    }
}

function change_menu_antivirus(id) {
    for (let i = 0; i < dwdc21e12d.length; i++) {
        $("#" + dwdc21e12d[i]).css("color", "#ffffffd4");
    }

    for (let i = 0; i < frg45th9nd.length; i++) {
        $("#" + frg45th9nd[i]).hide();
    }

    $("#" + dwdc21e12d[id]).css("color", "#766aff");
    $('#' + frg45th9nd[id]).show();
}

const fileUpload = () => {
    const $inputFile = $('#upload-files');
    const $inputContainer = $('#upload-container');
    const $filesListContainer = $('#files-list-container');
    const $uploadButton = $('#upload-button');
    const $uploadButton1 = $('#upload-button1');
    let fileList = [];

    $inputFile.on('click dragstart dragover', () => {
        $inputContainer.addClass('active');
    });

    $inputFile.on('dragleave dragend drop change', () => {
        $inputContainer.removeClass('active');
        const files = Array.from($inputFile[0].files);

        fileList = [];

        files.forEach(file => {
            const fileName = file.name;
            const uploadedFiles = {
                name: fileName,
                file: file
            };

            fileList.push(uploadedFiles);

            $filesListContainer.html('');

            const content = `
                <div class="form__files-container">
                    <span class="form__text">${uploadedFiles.name}</span>
                </div>
            `;
            $filesListContainer.append(content);
        });
    });

    $uploadButton.on('click', () => {
        if (fileList.length === 0) {
            message_window('Виберіть файл перед відправкою');
            return;
        }

        $('#we332dvc').html("Обробка");

        const formData = new FormData();
        formData.append('file', fileList[0].file);
        formData.append('value', 0);

        $.ajax({
            url: '/antivirus_bekend',
            method: 'POST',
            data: formData,
            processData: false,
            contentType: false,
            success: function (response) {
                if (response == 0) {
                    $('#we332dvc').html('<span class="f343ffv1">Вірусів незнайдено</span>');
                } else {
                    $('#we332dvc').html('<span class="f343ffv">Обережно вірус</span>');
                }
            },
            error: function (error) {
                console.error('Помилка відправки:', error);
            }
        });
    });

    $uploadButton1.on('click', () => {
        if (fileList.length === 0) {
            message_window('Виберіть файл перед відправкою');
            return;
        }

        $('#we332dvc').html("Обробка");

        const formData = new FormData();
        formData.append('file', fileList[0].file);
        formData.append('value', 1);

        $.ajax({
            url: '/antivirus_bekend',
            method: 'POST',
            data: formData,
            processData: false,
            contentType: false,
            success: function (response) {
                if (response == 0) {
                    $('#we332dvc').html('<span class="f343ffv1">Вірусів незнайдено</span>');
                } else {
                    $('#we332dvc').html('<span class="f343ffv">Обережно вірус</span>');
                }
            },
            error: function (error) {
                console.error('Помилка відправки:', error);
            }
        });
    });
};

function encryption_file_start() {
    const fileInput = document.getElementById('file_dwqdw');
    const files = fileInput.files;

    if (files.length === 0) {
        message_window('Файл невибрано');
        return;
    }

    $('#errewfffee').html("Обробка");

    const file = files[0];
    const formData = new FormData();

    formData.append('file', file);
    formData.append('filename', file.name);

    $.ajax({
        url: '/encryption_file',
        type: 'POST',
        processData: false,
        contentType: false,
        data: formData,
        success: function (response) {
            message_window('Успішно');
            encryption_file_end(response)

            const link = document.createElement('a');
            link.href = '/static/data/main.enc';
            link.download = 'main.enc';
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
        },
        error: function (xhr, status, error) {
            console.log("Error: " + error);
            console.log("Response text:", xhr.responseText);
        }
    });
}

function encryption_file_end(response) {
    $('#errewfffee').html("");

    $('#errewfffee').html(`
        <p class="deferghhhh777">Ванш куюч для розшифрування:</p> 
        <br>
        <span class="dwedffvvvvv">${response}</span>
    `);
}

const fileUpload_1 = () => {
    const $inputFile = $('#file_dwqdw');
    const $inputContainer = $('#upload-container');
    const $filesListContainer = $('#files-list-container');
    let fileList = [];

    $inputFile.on('click dragstart dragover', () => {
        $inputContainer.addClass('active');
    });

    $inputFile.on('dragleave dragend drop change', () => {
        $inputContainer.removeClass('active');
        const files = Array.from($inputFile[0].files);

        fileList = [];

        files.forEach(file => {
            const fileName = file.name;
            const uploadedFiles = {
                name: fileName,
                file: file
            };

            fileList.push(uploadedFiles);

            $filesListContainer.html('');

            const content = `
                <div class="form__files-container">
                    <span class="form__text">${uploadedFiles.name}</span>
                </div>
            `;
            $filesListContainer.append(content);
        });
    });
};

function decipher_file() {
    const fileInput = document.getElementById('file_dwqdw1');
    const files = fileInput.files;

    if (files.length === 0) {
        message_window('Файл невибрано');
        return;
    }

    $('#bbbnsqee343').html("Обробка");

    const file = files[0];
    const formData = new FormData();

    formData.append('file', file);
    formData.append('key', document.getElementById('cwwzevbnnn').value);

    $.ajax({
        url: '/decipher_file',
        type: 'POST',
        processData: false,
        contentType: false,
        data: formData,
        success: function (response) {
            if (response === 0) {
                message_window('Помилка');
            } else {
                message_window('Успішно');
                $('#bbbnsqee343').html("Успішно");

                const link = document.createElement('a');
                link.href = '/static/data/main';
                link.download = 'main';
                document.body.appendChild(link);
                link.click();
                document.body.removeChild(link);
            }
        },
        error: function (xhr, status, error) {
            console.log("Error: " + error);
            console.log("Response text:", xhr.responseText);
        }
    });
}

const fileUpload_2 = () => {
    const $inputFile = $('#file_dwqdw1');
    const $inputContainer = $('#upload-container');
    const $filesListContainer = $('#files-list-container');
    let fileList = [];

    $inputFile.on('click dragstart dragover', () => {
        $inputContainer.addClass('active');
    });

    $inputFile.on('dragleave dragend drop change', () => {
        $inputContainer.removeClass('active');
        const files = Array.from($inputFile[0].files);

        fileList = [];

        files.forEach(file => {
            const fileName = file.name;
            const uploadedFiles = {
                name: fileName,
                file: file
            };

            fileList.push(uploadedFiles);

            $filesListContainer.html('');

            const content = `
                <div class="form__files-container">
                    <span class="form__text">${uploadedFiles.name}</span>
                </div>
            `;
            $filesListContainer.append(content);
        });
    });
};

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

$(".toggle-btn").click(function() {
    let nextDropdown = $(this).next(".dropdown-content");
    
    if (nextDropdown.css("display") === "none") {
        nextDropdown.show();
        $(this).css("color", "#766aff");
    } else if (nextDropdown.css("display") === "block") {
        nextDropdown.hide();
        $(this).css("color", "#fff");
    }
});

function cleanup() {
    message_window('Очищення компютера');

    $.ajax({
        url: "/cleanup",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data_cleaning),
        success: function (response) {
            message_window('Завершено');
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

$(".ump_button").click(function() {
    let data = $(this).data("data");
    let currentColor = $(this).css("background-color");

    if (currentColor === "rgb(50, 52, 77)") {
        $(this).css("background-color", "#565574");
        data_cleaning[data] = 1;
    } else if (currentColor === "rgb(86, 85, 116)") {
        $(this).css("background-color", "#32344d");
        data_cleaning[data] = 0;
    }
});

function open_site() {
    $.ajax({
        url: "/browser_site_app",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            message_window('Сайт відкрито');
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function config_bg() {
    $.ajax({
        url: "/config_global",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            antivirus_flash_drive = response['antivirus']['antivirus_flash_drive'];

            $("#bg_input").val(response['antivirus']['antivirus_flash_drive_cmd']);

            if (antivirus_flash_drive === 0) {
                $("#bg_dqwderfd").css("background-color", "#22223a");
            } else {
                $("#bg_dqwderfd").css("background-color", "#565574");
            }
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function new_val_gb_usb() {
    if (antivirus_flash_drive === 0) {
        antivirus_flash_drive = 1;
    } else {
        antivirus_flash_drive = 0;
    }

    $.ajax({
        url: "/change_val_gb_usb",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: antivirus_flash_drive, data1: $("#bg_input").val()}),
        success: function (response) {
            message_window('Значення встановлено');
            config_bg();
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function change_lang_now() {
    $.ajax({
        url: "/config_global",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            lang_global = response['lang'];
            lang_change_page(lang_global);
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function lang_change_page(lang) {
    if (lang === "en") {
        // wifi
        $('#lang_wefsdeeeeee').html("WIFI management");
        $('#lang_nfefdfdvghyt').html("Get more information about yours WIFI");
        $('#lang_dqwedddcwww').html("Connection:");

        // system
        $('#lang_system_wdeewds').html("Data os");
        $('#lang_system_fesrdfyyyyy').html("library data");
        $('#lang_system_wfr839wefsff').html("NetworkAdapters data");
        $('#lang_system_v00qwdweee').html("Open programs");
        $('#lang_system_vfd8723ed').html("System data");
        $('#lang_system_verdfvcww').html("There is something interesting here");

        // antivirus
        $('#antivirus_errfee2').html("Antivirus");
        $('#antivirus_vef093f').html("Be in information security");
        $('#dwdc21e12d').html("Check the site");
        $('#dwdc21e12d1').html("Check the file");
        $('#dwdc21e12d2').html("On the background");
        $('#antivirus_894534ffvvv').html("Check");
        $('#h3ruiwefer24f').html("Description");
        $('#fewrvw243rgefvcc').html("Background features are enabled automatically at startup");
        $('#vbs612dwes655').html("Monitoring flash drives");
        $('#vb92354gu04ttg').html("This setting activates the USB media monitoring function, which notifies users when new flash drives are connected. After enabling this function, the program constantly scans available USB media and reports each new device that is connected.");
        $('#bbv612ee3dwe').html("Set the cmd command that will be launched after connecting the USB drive");
        $('#bg_dqwderfd').html("Turn on");
        $('#upload-button').html("Quick check");
        $('#upload-button1').html("Detailed check");

        // cleaning
        $('#vc728i32000').html("Cleaning pc");
        $('#bvh9iuweddd').html("Unfortunately, the script does not clean the System32 folder. Sorry");
    } else if (lang === "uk") {
        // wifi
        $('#lang_wefsdeeeeee').html("Вайфай менеджер");
        $('#lang_nfefdfdvghyt').html("Отримайте більше інформації про ваш вайфай");
        $('#lang_dqwedddcwww').html("З'єднання:");

        // system
        $('#lang_system_wdeewds').html("Дані операційної системи");
        $('#lang_system_fesrdfyyyyy').html("Дані бібліотек");
        $('#lang_system_wfr839wefsff').html("Дані мережевих адаптерів");
        $('#lang_system_v00qwdweee').html("Відкриті програми");
        $('#lang_system_vfd8723ed').html("Системні дані");
        $('#lang_system_verdfvcww').html("Тут є дещо цікаве можливо");
        
        // antivirus
        $('#antivirus_errfee2').html("Антивірус");
        $('#antivirus_vef093f').html("Будьте в інформаційній безпеці");
        $('#dwdc21e12d').html("Перевірити сайт");
        $('#dwdc21e12d1').html("Перевірити файл");
        $('#dwdc21e12d2').html("На фоні");
        $('#antivirus_894534ffvvv').html("Перевірити");
        $('#h3ruiwefer24f').html("Опис");
        $('#fewrvw243rgefvcc').html("Функції на фоні вмикаються автоматично при запуску");
        $('#vbs612dwes655').html("Моніторинг флешок");
        $('#vb92354gu04ttg').html("Це налаштування активує функцію моніторингу USB-носіїв, яка попереджає користувача при підключенні нових флешок. Після ввімкнення цієї функції програма постійно сканує доступні USB-носії і повідомляє про кожен новий пристрій, що підключається.");
        $('#bbv612ee3dwe').html("Встановити cmd команду яка буде запускатись після підключення USB-носія");
        $('#bg_dqwderfd').html("Увімкнути");
        $('#upload-button').html("Швидка перевірка");
        $('#upload-button1').html("Детальна перевірка");

        // cleaning
        $('#vc728i32000').html("Очищення ПК");
        $('#bvh9iuweddd').html("На жаль скрипт не очищує папку System32. Вибачте");

    }
}

function get_massage_info() {
    openModal("modal1");

    $.ajax({
        url: "http://localhost:3000/data",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            $('#r4e23efvd_').html(response['message']);
            $('#textfrwefwf_').html(response['desc']);
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

