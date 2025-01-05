function antivirus_web_start() {
    let inputValue = $('#fkwe9203f').val();
    const dataToSend = {url_site: [inputValue]};

    if (lang_global === "uk") {
        $('#dwdefw4f4').text('Перевірка.....');
    } else if (lang_global === "en") {
        $('#dwdefw4f4').text('Audit.....');
    }

    $.ajax({
        url: '/antivirus_web',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(dataToSend),
        success: function (response) {
            antivirus_web_end(response)
        }
    });
}

function antivirus_web_end(response) {
    $('#dqdcew336g').show();

    clean_div("dq13892r2323233313");

    if (lang_global === "uk") {
        $('#dwdefw4f4').text('Завершено'); 
    } else if (lang_global === "en") {
        $('#dwdefw4f4').text('Completed'); 
    }

    if (lang_global === "uk") {
        $('#we31f3qecsdx13rv1').prepend("SSL сервтифікат: ");
        $('#we31f3qecsdx13rv').prepend("Загрози: ");
        $('#we31f3qecsdx13rv2').prepend("DNS зміни: ");

        if (response['ssl'] === 1) {
            $('#dwdefw4f4ewqe').html("працює");
        } else {
            $('#dwdefw4f4ewqe').html("не працює");
        }
    
        if (response['url'] === 1) {
            $('#ewfsdt4w43tgfd321').html("не знайдено");
        } else {
            $('#ewfsdt4w43tgfd3211').html("pнайдені");
        }
    } else if (lang_global === "en") {
        $('#we31f3qecsdx13rv1').prepend("SSL Certificate: ");
        $('#we31f3qecsdx13rv').prepend("Threats: ");
        $('#we31f3qecsdx13rv2').prepend("DNS Changes: ");
        
        if (response['ssl'] === 1) {
            $('#dwdefw4f4ewqe').html("working");
        } else {
            $('#dwdefw4f4ewqe').html("not working");
        }
        
        if (response['url'] === 1) {
            $('#ewfsdt4w43tgfd321').html("Not found");
        } else {
            $('#ewfsdt4w43tgfd3211').html("Found");
        }        
    }

    $('#qefwfvd244ttff').html(response['dns']);
}

const fileUpload = () => {
    const $inputFile = $('#upload-files');
    const $inputContainer = $('#upload-container');
    const $filesListContainer = $('#files-list-container');
    const $uploadButton = $('#upload-button');
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
            if (lang_global === "uk") {
                message_window('Виберіть файл перед відправкою'); 
            } else if (lang_global === "en") {
                message_window('Select a file before sending'); 
            }
            return;
        }

        if (lang_global === "uk") {
            $('#we332dvc').html("Обробка");
        } else if (lang_global === "en") {
            $('#we332dvc').html("Processing");
        }

        const formData = new FormData();
        formData.append('file', fileList[0].file);

        $.ajax({
            url: '/antivirus_bekend',
            method: 'POST',
            data: formData,
            processData: false,
            contentType: false,
            success: function (response) {
                $('#d13dqe021w34fwvqedddd').hide();
                clean_div("f2ewds322r3345trg");
                data_bekend_solver(response);
            }
        });
    });
};

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
                $("#bg_dqwderfd").css("background-color", "#181822");
            } else {
                $("#bg_dqwderfd").css("background-color", "#373745");
            }
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
            if (lang_global === "uk") {
                message_window('Значення встановлено');
            } else if (lang_global === "en") {
                message_window('The value is set');
            }

            config_bg();
        }
    });
}

function data_bekend_solver(response) {
    response = JSON.parse(response);

    if (response['status'] === 2) {
        if (lang_global === "uk") {
            $('#we332dvc1').html('Завершено успішно');
            $('#we332dvc3').html('Хеш файлу: ' + response['hash']);
        } else if (lang_global === "en") {
            $('#we332dvc1').html('Completed successfully');
            $('#we332dvc3').html('The hash of the file: ' + response['hash']);
        }
    
        if (response['data'] != "") {
            data_json_exe = JSON.parse(response['data']);
    
            if (data_json_exe === "") {
                if (lang_global === "uk") {
                    $('#we332dvc4').html('Файл не є exe дані детальні дані отримати невдалось');
                    $('#d13dqe021w34fwvqedddd').hide();
                } else if (lang_global === "en") {
                    $('#we332dvc4').html('The file is not an exe data detailed data failed to get');
                    $('#d13dqe021w34fwvqedddd').hide();
                }
            } else {
                if (lang_global === "uk") {
                    $('#we332dvc4').html('Детальні дані про exe');
                    $('#d13dqe021w34fwvqedddd').show();
                    $('#d13dqe021w34fwvqedddd').html('Скачати');
                } else if (lang_global === "en") {
                    $('#we332dvc4').html('More details about the exe');
                    $('#d13dqe021w34fwvqedddd').show();
                    $('#d13dqe021w34fwvqedddd').html("Download");
                }
            }
        }
    
        if (response['namber'] === 0) {
            if (lang_global === "uk") {
                $('#we332dvc2').html('<span class="f343ffv1 grean_1">Вірусів незнайдено</span>');
            } else if (lang_global === "en") {
                $('#we332dvc2').html('<span class="f343ffv1 grean_1">No viruses found</span>');
            }
        } else {
            if (lang_global === "uk") {
                $('#we332dvc2').html(`<span class="f343ffv read">Вірус знайдено кількість ${response['namber']}</span>`);
            } else if (lang_global === "en") {
                $('#we332dvc2').html(`<span class="f343ffv read">Virus found number ${response['namber']}</span>`);
            }
        }
    } else {
        if (lang_global === "uk") {
            $('#we332dvc1').html('Завершено з помилкою зачекайте декілька хвилин щоб ваш файл обробився і повторять запрос');
        } else if (lang_global === "en") {
            $('#we332dvc1').html('Completed with an error, please wait a few minutes for your file to be processed and the request will be repeated');
        }
    }
}

function download_json_data() {
    const jsonData = JSON.stringify(data_json_exe, null, 2); 

    const blob = new Blob([jsonData], { type: "application/json" });
    const url = URL.createObjectURL(blob);
    const a = document.createElement("a");

    a.href = url;
    a.download = "data.json";
    document.body.appendChild(a);
    a.click();

    document.body.removeChild(a);
    URL.revokeObjectURL(url);
}

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
