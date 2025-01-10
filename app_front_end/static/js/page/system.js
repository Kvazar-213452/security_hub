function get_os_data() {
    $.ajax({
        url: "/get_os_data",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            write_os_data(response)
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
        }
    });
}

function write_os_data(response) {
    let bios_info = response['bios_info'];
    let host_info = response['host_info'];
    let operating_system_info = response['operating_system_info'];

    $('#ddcbnxcew33333').html(null);
    $('#ndwe8rfier').html(null);
    $('#bfgtey65yt').html(null);

    let text = `
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Виробник" : (lang_global === "en" ? "Manu facturer" : "")}</p>
            <p class="desc_o">${bios_info[0]['manufacturer']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Імя" : (lang_global === "en" ? "Тame" : "")}</p>
            <p class="desc_o">${bios_info[0]['name']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Дата випуску" : (lang_global === "en" ? "Release date" : "")}</p>
            <p class="desc_o">${bios_info[0]['release_date']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Серійний номер" : (lang_global === "en" ? "Serial number" : "")}</p>
            <p class="desc_o">${bios_info[0]['serial_number']}</p>
            <div class="hr_div"></div>
        </div>
    `;

    $('#ddcbnxcew33333').append(text);

    let data_text = `
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Імя хоста" : (lang_global === "en" ? "Host name" : "")}</p>
            <p class="desc_o">${host_info['hostname']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Арка ядра" : (lang_global === "en" ? "kernel arch" : "")}</p>
            <p class="desc_o">${host_info['kernel_arch']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Платформа" : (lang_global === "en" ? "Platform" : "")}</p>
            <p class="desc_o">${host_info['platform']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Сімейство платформ" : (lang_global === "en" ? "Platform family" : "")}</p>
            <p class="desc_o">${host_info['platform_family']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Версія платформи" : (lang_global === "en" ? "Platform version" : "")}</p>
            <p class="desc_o">${host_info['platform_version']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Час роботи" : (lang_global === "en" ? "Uptime" : "")}</p>
            <p class="desc_o">${host_info['uptime']} [s]</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Номер складання" : (lang_global === "en" ? "Build number" : "")}</p>
            <p class="desc_o">${operating_system_info[0]['build_number']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Дата встановлення" : (lang_global === "en" ? "Install date" : "")}</p>
            <p class="desc_o">${operating_system_info[0]['install_date']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Час останнього завантаження" : (lang_global === "en" ? "Last boot up time" : "")}</p>
            <p class="desc_o">${operating_system_info[0]['last_boot_up_time']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Виробник" : (lang_global === "en" ? "Manufacturer" : "")}</p>
            <p class="desc_o">${operating_system_info[0]['manufacturer']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Серійний номер" : (lang_global === "en" ? "Serial number" : "")}</p>
            <p class="desc_o">${operating_system_info[0]['serial_number']}</p>
            <div class="hr_div"></div>
        </div>
    `;

    $('#bfgtey65yt').append(data_text);
}

function render_data_window_open(response) {
    $('#sifewfewx').html(null);
    response = response['devices']

    for (let i = 0; i < response.length; i++) {
        if (i + 1 !== response.length) {
            let text = `<div class="div_wifi_all">
            <p class="name_wifi_div_all">${response[i]}</p>
            </div>`;
            $('#sifewfewx').append(text);
        }
    }
}

// setings new section
// setings new section
// setings new section

let mas_file_extension = [];

function render_div() {
    $('#extension_select').html(null);

    for (let i = 0; i < mas_file_extension.length; i++) {

        let text = `
        <div class="file_system_div_select">
            <p>${mas_file_extension[i]}</p>
            <div onclick="del_div('${mas_file_extension[i]}', 0)" class="file_system_del_div">del</div>
        </div>
    `;

    $('#extension_select').append(text);
    }
    
}

function add_div() {
    let text_file_extension = $('#file_stystem_dwqdasz').val();

    mas_file_extension.push(text_file_extension);

    render_div()
}

function del_div(text_file_extension) {
    mas_file_extension.splice(text_file_extension, 1);

    render_div()
}

function render_top_10(data) {
    $('#data_bate').html(data.Rootsize);

    for (let i = 0; i < 9; i++) {
        let text = `
        <div class="top_10_extension"><span1>${data.Top[i][0]}</span1><span2>${data.Top[i][1]}</span2><span3>${data.Top[i][2]}</span3></div>
        `;
        
        $('#out_data_extension').append(text);
    }
}

function inoxwd() {
    let dir_ =  $('#file_stystem_dwqdasz_2').val();
    let mas1 = mas_file_extension;

    $.ajax({
        url: '/scan_dir',
        type: 'POST',
        data: JSON.stringify({ dir: dir_, mas1: mas1 }),
        processData: false,
        contentType: 'application/json',
        success: function (response) {
            render_top_10(response);
        }
    });
}
