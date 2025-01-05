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
    let data = response['data'];
    let jsonData = JSON.parse(data);

    $('#ddcbnxcew33333').html(null);
    $('#ndwe8rfier').html(null);
    $('#bfgtey65yt').html(null);

    let text = `
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Операційна система" : (lang_global === "en" ? "Operating System" : "")}</p>
            <p class="desc_o">Name: ${jsonData['OS']['Name']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Архітектура" : (lang_global === "en" ? "Architecture" : "")}</p>
            <p class="desc_o">${jsonData['Architecture']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Диск" : (lang_global === "en" ? "Disk" : "")}</p>
            <p class="desc_o">FreeSpace: ${jsonData['Disk']['FreeSpace']}</p>
            <p class="desc_o_1">TotalSpace: ${jsonData['Disk']['TotalSpace']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Память" : (lang_global === "en" ? "Memory" : "")}</p>
            <p class="desc_o">FreeMemory: ${jsonData['Memory']['FreeMemory']}</p>
            <p class="desc_o_1">FreeVirtualMemory: ${jsonData['Memory']['FreeVirtualMemory']}</p>
            <p class="desc_o_2">TotalMemory: ${jsonData['Memory']['TotalMemory']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Кількість процесорів" : (lang_global === "en" ? "Processor сount" : "")}</p>
            <p class="desc_o">ProcessorCount: ${jsonData['ProcessorCount']}</p>
            <div class="hr_div"></div>
        </div>
        <div class="div_info_os">
            <p class="name_o">${lang_global === "uk" ? "Час роботи системи" : (lang_global === "en" ? "System uptime" : "")}</p>
            <p class="desc_o">time: ${jsonData['SystemUptime']['Days']}:${jsonData['SystemUptime']['Hours']}:${jsonData['SystemUptime']['Minutes']}:${jsonData['SystemUptime']['Seconds']}</p>
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
        if (i + 1 !== response.length) {
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
        }
    });
}

function resource_info_render_data(response) {
    response = response['data'];

    $('#dwqwdwfcfff44').text(response[0] || 'N/A');
    // $('#rggwiovnewcee').text(response[1] || 'N/A');
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
