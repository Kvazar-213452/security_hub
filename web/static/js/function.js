function button_hover(name) {
    $("#" + name).css("background-color", "#2e363f");
    $("#" + name).css("text-decoration", "underline");
    $("#" + name).css("text-decoration-thickness", "2px");
    $("#" + name).css("text-underline-offset", "4px");
}

function clos(name) {
    $('#' + name).hide(); 
}

function openModal(name) {
    $('#' + name).show(); 
}

function fetchLogs() {
    $.post('/get_logs', function(data) {
        const logsArray = data.log.split('\n');

        $('.console').html(logsArray.join('<br>'));
    });
}

function console_open() {
    fetchLogs();
    $('.console').toggle();
}

function button_active(name, mas) {
    mas.forEach(function(item) {
        $("#" + item).css({
            "border": "none",
            "color": "#ffffffd4"
        });
    });

    $("#" + name).css({
        "border": "2px solid #55c959",
        "color": "#55c959"
    });
}

function get_data_wifi_now() {
    $.ajax({
        url: "/get_wifi_now",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function(response) {
            data_wifi_render_now(response)
            checkUnsafeProtocols();
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function get_data_wifi_now_level() {
    $.ajax({
        url: "/get_wifi_now",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function(response) {
            level_wifi_render(response['signal_strength'])
        },
        error: function(xhr, status, error) {
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
        success: function(response) {
            get_network_now(function(ssid) {
                render_all_network_wifi(response, ssid);
            });
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function data_wifi_render_now(response) {
    $('#name_wifi').text(response['ssid_name'] || 'N/A');
    $('#authentication_wifi').text(response['authentication'] || 'N/A');
    $('#description_wifi').text(response['description'] || 'N/A');
    $('#guid_wifi').text(response['guid'] || 'N/A');
    $('#physical_address_wifi').text(response['physical_address'] || 'N/A');
    $('#radio_type_wifi').text(response['radio_type'] || 'N/A');
    $('#signal_strength_wifi').text(response['signal_strength'] || 'N/A');
    $('#state_wifi').text(response['state'] || 'N/A');
    $('#cost_wifi').text(response['cost'] || 'N/A');
    $('#ciphers_wifi').text(response['ciphers'] ? response['ciphers'].join(', ') : 'N/A');
    $('#key_content_wifi').text(response['key_content'] || 'N/A');
    $('#profile_name_wifi').text(response['profile_name'] || 'N/A');
    $('#cost_source_wifi').text(response['cost_source'] || 'N/A');
    $('#approaching_limit_wifi').text(response['approaching_limit'] || 'N/A');
    $('#congested_wifi').text(response['congested'] || 'N/A');
    $('#over_limit_wifi').text(response['over_limit'] || 'N/A');
    $('#roaming_wifi').text(response['roaming'] || 'N/A');
    $('#vendor_extension_wifi').text(response['vendor_extension'] || 'N/A');
    $('#version_wifi').text(response['version'] || 'N/A');
}

function max_wifi() {
    $("#signal_1").removeClass("curveOne2");
    $("#signal_1").addClass("curveOne1");

    $("#signal_2").removeClass("curveTwo2");
    $("#signal_2").addClass("curveTwo1");

    $("#signal_3").removeClass("curveThree2");
    $("#signal_3").addClass("curveThree1");

    $("#signal_4").removeClass("curveFour2");
    $("#signal_4").addClass("curveFour1");
}

function checkUnsafeProtocols() {
    const text = $("#ciphers_wifi").text();

    if (text) {
        for (const protocol of unsafeProtocols) {
            if (text.includes(protocol)) {
                $("#wifi_protection").html(`<p class="wifi_3_div_red">Незахищено</p>`);
                break;
            }
        }
    }

    $("#wifi_protection").html(`<p class="wifi_3_div">Захищено</p>`);
}

function level_wifi_render(level) {
    if (level > 90) {
        max_wifi();

    } else if (level > 75) {
        max_wifi();
        $("#signal_1").removeClass("curveOne1");
        $("#signal_1").addClass("curveOne2");
    } else if (level > 50) {
        max_wifi();
        $("#signal_1").removeClass("curveOne1");
        $("#signal_1").addClass("curveOne2");
        
        $("#signal_2").removeClass("curveTwo1");
        $("#signal_2").addClass("curveTwo2");
    } else if (level > 25) {
        max_wifi();
        $("#signal_1").removeClass("curveOne1");
        $("#signal_1").addClass("curveOne2");
        
        $("#signal_2").removeClass("curveTwo1");
        $("#signal_2").addClass("curveTwo2");
        
        $("#signal_3").removeClass("curveThree1");
        $("#signal_3").addClass("curveThree2");
    } else if (level === 0) {
        max_wifi();
        $("#signal_1").removeClass("curveOne1");
        $("#signal_1").addClass("curveOne2");
        
        $("#signal_2").removeClass("curveTwo1");
        $("#signal_2").addClass("curveTwo2");
        
        $("#signal_3").removeClass("curveThree1");
        $("#signal_3").addClass("curveThree2");

        $("#signal_").removeClass("curveFour1");
        $("#signal_").addClass("curveFour2");
    }
}

function render_all_network_wifi(response, ssid) {
    $('#render_all_wifi').html(null);

    for (let i = 0; i < response.length; i++) {
        let text = `<div class="div_wifi_all">
        <p class="name_wifi_div_all">${response[i]['ssid']}</p>
        <p class="right_left_signal">${response[i]['signal']}</p>
        </div>`;
        $('#render_all_wifi').append(text);
    }

    $('.name_wifi_div_all').each(function() {
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
        success: function(response) {
            if (callback) {
                callback(response['ssid']);
            }
        },
        error: function(xhr, status, error) {
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
        let server = data['server'];

        $('#ssdfredfgettt').val(port);
        $('#bsdcfvbttrfgo').val(server);
        
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
        success: function(response) {

        },
        error: function(xhr, status, error) {
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
        success: function(response) {
            write_os_data(response)
        },
        error: function(xhr, status, error) {
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
        success: function(response) {
            render_data_window_open(response);
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function write_os_data(response) {
    $('#adqwdwxxxx').text(response['computer_name'] || 'N/A');
    $('#qwefsbsfb').text(response['os_version'] || 'N/A');
    $('#efwfwefwegwecc').text(response['processor_info'] || 'N/A');
    $('#fewffefevrewew').text(response['system_memory'] || 'N/A');
    $('#gwghyjyyjy').text(response['system_uptime'] || 'N/A');
    $('#retrevvvvde').text(response['user_name'] || 'N/A');
}

function render_data_window_open(response) {
    $('#sifewfewx').html(null);
    response = response['devices']

    for (let i = 0; i < response.length; i++) {
        let text = `<div class="div_wifi_all">
        <p class="name_wifi_div_all">${response[i]}</p>
        </div>`;
        $('#sifewfewx').append(text);
    }
}

function resource_info() {
    $.ajax({
        url: "/resource_info",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function(response) {
            resource_info_render_data(response);
        },
        error: function(xhr, status, error) {
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

function cleanup() {
    message_window('Очищення компютера');

    $.ajax({
        url: "/cleanup",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function(response) {
            
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function button_active_antivitys(name) {
    $("#" + name).css("text-decoration", "underline");
    $("#" + name).css("text-decoration-thickness", "2px");
    $("#" + name).css("text-underline-offset", "4px");
    $("#" + name).css("color", "#55c959");
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
        success: function(response) {
            antivirus_web_end(response)
        },
        error: function(xhr, status, error) {
            console.log("Error: " + error);
            console.log("Response text:", xhr.responseText);
        }
    });
}

function antivirus_web_end(response) {
    $('#dqdcew336g').show(); 
    $('#dwdefw4f4').text('Завершено');

    if (response['found'] === false) {
        $('#dw93244444').text('Сайт безпечний');
    } else {
        $('#dw93244444').text('Сайт небезпечний');
    }

    const decodedString = atob(response['data']);
    const jsonObject = JSON.parse(decodedString);
    
    $('#dwdw3333rffcc').html(null);

    jsonObject1 = jsonObject['data'];

    let i = 0;
    while (i < jsonObject1.length) {
        if (jsonObject1[i] === "window.location.href") {
            $('#dwdw3333rffcc').append("<li>Переадресацію на другі джерела</li><br>");
        } else if (jsonObject1[i] === "window.open") {
            $('#dwdw3333rffcc').append("<li>Використання pop-up вікон</li><br>");
        } else if (jsonObject1[i] === "requestFullscreen") {
            $('#dwdw3333rffcc').append("<li>Відкриття на повний екран</li><br>");
        }

        i++;
    }
}

function change_menu_antivirus(type, id) {
    $("#dwdc21e12d").css("text-decoration", "none");
    $("#dwdc21e12d").css("text-decoration-thickness", "none");
    $("#dwdc21e12d").css("text-underline-offset", "none");
    $("#dwdc21e12d").css("color", "#ffffffd4");

    $("#dwdc21e12d1").css("text-decoration", "none");
    $("#dwdc21e12d1").css("text-decoration-thickness", "none");
    $("#dwdc21e12d1").css("text-underline-offset", "none");
    $("#dwdc21e12d1").css("color", "#ffffffd4");

    button_active_antivitys(id);

    if (type === 0) {
        $('#frg45th9nd1').hide(); 
        $('#frg45th9nd').show();
    } else if (type === 1) {
        $('#frg45th9nd').hide(); 
        $('#frg45th9nd1').show();
    }
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
            success: function(response) {
                if (response == 0) {
                    $('#we332dvc').html('<span class="f343ffv1">Вірусів незнайдено</span>');
                } else {
                    $('#we332dvc').html('<span class="f343ffv">Обережно вірус</span>');
                }
            },
            error: function(error) {
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
            success: function(response) {
                if (response == 0) {
                    $('#we332dvc').html('<span class="f343ffv1">Вірусів незнайдено</span>');
                } else {
                    $('#we332dvc').html('<span class="f343ffv">Обережно вірус</span>');
                }
            },
            error: function(error) {
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
        success: function(response) {
            message_window('Успішно');
            encryption_file_end(response)

            const link = document.createElement('a');
            link.href = '/static/data/main.enc'; 
            link.download = 'main.enc';
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
        },
        error: function(xhr, status, error) {
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
        success: function(response) {
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
        error: function(xhr, status, error) {
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

function log_change(val) {
    $.ajax({
        url: "/log_change",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({value: val}),
        success: function(response) {
            getConfig();
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function port_change() {
    let value = $('#ssdfredfgettt').val();

    $.ajax({
        url: "/port_change",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({value: value}),
        success: function(response) {
            getConfig();
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function server_change() {
    let value = $('#bsdcfvbttrfgo').val();

    $.ajax({
        url: "/server_change",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({value: value}),
        success: function(response) {
            getConfig();
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}
