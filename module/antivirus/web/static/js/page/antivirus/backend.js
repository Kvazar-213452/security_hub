// app_front_end/static/js/page/antivirus/backend.js

function antivirus_bekend_scan_dir() {
    let inputValue = $('#fkwe9203f1d').val();
    $("#hfweo23fwesd").html(null);
    if (lang_global === "uk") {
        $("#hfweo23fwesd").append("<p class='scan_dir'>Сканування</p>");
    } else if (lang_global === "en") {
        $("#hfweo23fwesd").append("<p class='scan_dir'>Scanning</p>");
    }
    $.ajax({
        url: '/antivirus_bekend_scan_dir',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({dir: inputValue}),
        success: function (response) {
            render_data_scan_dir(response);
        }
    });
}

function render_data_scan_dir(response) {
    $("#hfweo23fwesd").html(null);
    if (response["detected_viruses"].length === 0) {
        if (lang_global === "uk") {
            $("#hfweo23fwesd").append("<p class='scan_dir'>Загроз не виявлено</p>");
        } else if (lang_global === "en") {
            $("#hfweo23fwesd").append("<p class='scan_dir'>No threats detected</p>");
        }
        render_scaned_file(response);
    } else {
        scan_virus_file = response["detected_viruses"];
        if (lang_global === "uk") {
            $("#hfweo23fwesd").append(`<p class='scan_dirf'>Виявлено загрози [${response["detected_viruses"].length}]</p>`);
        } else if (lang_global === "en") {
            $("#hfweo23fwesd").append(`<p class='scan_dirf'>Threats detected [${response["detected_viruses"].length}]</p>`);
        }
        
        render_scaned_file(response);
        $("#hfweo23fwesd").append("<br><br>");
        if (lang_global === "uk") {
            $("#hfweo23fwesd").append("<p class='ooindq_dlffff_13'>Загрози</p>");
        } else if (lang_global === "en") {
            $("#hfweo23fwesd").append("<p class='ooindq_dlffff_13'>Threats</p>");
        }
        $("#hfweo23fwesd").append("<br>");
        $("#hfweo23fwesd").append("<div class='ump_u3fgbbb'><div id='ump_u3fgbbb_1_1' class='padoge'></div></div>");
        render_unixo(response["detected_viruses"]);
    }
}

function render_scaned_file(response) {
    $("#hfweo23fwesd").append("<br><br>");
    if (lang_global === "uk") {
        $("#hfweo23fwesd").append(`<p class='ooindq_dlffff'>Відскановані файли [${response["total_exe_files"]}]</p>`);
    } else if (lang_global === "en") {
        $("#hfweo23fwesd").append(`<p class='ooindq_dlffff'>Scanned files [${response["total_exe_files"]}]</p>`);
    }
    $("#hfweo23fwesd").append("<br>");
    $("#hfweo23fwesd").append("<div class='ump_u3fgbbb'><div id='ump_u3fgbbb_1' class='padoge'></div></div>");
    for (let i = 0; i < response["total_exe_files"]; i++) {
        $("#ump_u3fgbbb_1").append(`
            <div class="div_pdq32fwesd">
                <p class="iewfv2222c">${response["checked_files"][i]['path']}</p>
                <p class="iewfv2222c1">${response["checked_files"][i]['hash']}</p>
            </div>
        `);
    }
}

function del_file_scan(index) {
    path = scan_virus_file[index]["path"];
    $.ajax({
        url: '/antivirus_bekend_del_file',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({path: path}),
        success: function (response) {
            console.log(response)
            if (response["status"] == 0) {
                if (lang_global === "uk") {
                    message_window('Файл не видалено');
                } else if (lang_global === "en") {
                    message_window('File not deleted');
                }
            } else {
                if (lang_global === "uk") {
                    message_window('Файл видалено');
                } else if (lang_global === "en") {
                    message_window('File deleted');
                }
                for (let i = 0; i < scan_virus_file.length; i++) {
                    if (scan_virus_file[i]["path"] === path) {
                        scan_virus_file.splice(i, 1);
                        render_unixo(scan_virus_file);
                    }
                }
            }
        }
    });
}

function render_unixo(detected_viruses) {
    $("#ump_u3fgbbb_1_1").html(null);
    if (lang_global === "uk") {
        for (let i = 0; i < detected_viruses.length; i++) {
            $("#ump_u3fgbbb_1_1").append(`
                <div class="div_pdq32fwesd">
                    <p class="iewfv2222c">${detected_viruses[i]['path']}</p>
                    <p class="iewfv2222c1">${detected_viruses[i]['hash']}</p>
                    <button onclick="del_file_scan(${i})" class="po12ldqwecsdbgb">Видалити</button>
                </div>
            `);
        }
    } else if (lang_global === "en") {
        for (let i = 0; i < detected_viruses.length; i++) {
            $("#ump_u3fgbbb_1_1").append(`
                <div class="div_pdq32fwesd">
                    <p class="iewfv2222c">${detected_viruses[i]['path']}</p>
                    <p class="iewfv2222c1">${detected_viruses[i]['hash']}</p>
                    <button onclick="del_file_scan(${i})" class="po12ldqwecsdbgb">Remove</button>
                </div>
            `);
        }
    }
}
