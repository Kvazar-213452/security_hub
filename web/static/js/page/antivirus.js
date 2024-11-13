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

    $('#dwdefw4f4ewqe').text(response['dns'] + response['ssl'] + response['url']); 
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
                console.log(response)
                if (response == 0) {
                    if (lang_global === "uk") {
                        $('#we332dvc').html('<span class="f343ffv1">Вірусів незнайдено</span>');
                    } else if (lang_global === "en") {
                        $('#we332dvc').html('<span class="f343ffv1">No viruses found</span>');
                    }
                } else {
                    if (lang_global === "uk") {
                        $('#we332dvc').html('<span class="f343ffv">Вірус знайдено</span>');
                    } else if (lang_global === "en") {
                        $('#we332dvc').html('<span class="f343ffv">Virus found</span>');
                    }
                }
            },
            error: function (error) {
                console.error('Помилка відправки:', error);
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
            if (lang_global === "uk") {
                message_window('Значення встановлено');
            } else if (lang_global === "en") {
                message_window('The value is set');
            }

            config_bg();
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}