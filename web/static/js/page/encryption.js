function encryption_file_start() {
    const fileInput = document.getElementById('file_dwqdw');
    const files = fileInput.files;

    if (files.length === 0) {
        if (lang_global === "uk") {
            message_window('Файл не вибрано');
        } else if (lang_global === "en") {
            message_window('No file selected');
        }

        return;
    }

    if (lang_global === "uk") {
        $('#bbbnsqee343').html("Обробка");
    } else if (lang_global === "en") {
        $('#bbbnsqee343').html("Processing");       
    }

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
        <p class="deferghhhh777">Your decryption key:</p> 
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
        if (lang_global === "uk") {
            message_window('Файл не вибрано');
        } else if (lang_global === "en") {
            message_window('No file selected');
        }

        return;
    }

    if (lang_global === "uk") {
        $('#bbbnsqee343').html("Обробка");
    } else if (lang_global === "en") {
        $('#bbbnsqee343').html("Processing");       
    }

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
                if (lang_global === "uk") {
                    message_window('Помилка');
                } else if (lang_global === "en") {
                    message_window('Error');                
                }
            } else {
                if (lang_global === "uk") {
                    message_window('Успішно');
                    $('#bbbnsqee343').html("Успішно"); 
                } else if (lang_global === "en") {
                    message_window('Success');
                    $('#bbbnsqee343').html("Success");                    
                }

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
