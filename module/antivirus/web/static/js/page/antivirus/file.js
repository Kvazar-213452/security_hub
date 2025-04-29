// app_front_end/static/js/page/antivirus/file.js

class FileUpload {
    constructor() {
        this.fileList = [];
        this.$inputFile = $('#upload-files');
        this.$inputContainer = $('#upload-container');
        this.$filesListContainer = $('#files-list-container');
        this.$uploadButton = $('#upload-button');
        this.data_json_exe = null;

        this.bindEvents();
    }

    bindEvents() {
        this.$inputFile.on('click dragstart dragover', () => {
            this.$inputContainer.addClass('active');
        });

        this.$inputFile.on('dragleave dragend drop change', () => {
            this.$inputContainer.removeClass('active');
            this.handleFileSelection();
        });

        this.$uploadButton.on('click', () => {
            this.uploadFile();
        });
    }

    handleFileSelection() {
        const files = Array.from(this.$inputFile[0].files);
        this.fileList = [];

        files.forEach(file => {
            const fileName = file.name;
            const uploadedFiles = {
                name: fileName,
                file: file
            };

            this.fileList.push(uploadedFiles);
            this.updateFileListDisplay(uploadedFiles);
        });
    }

    updateFileListDisplay(uploadedFiles) {
        this.$filesListContainer.html('');
        const content = `
            <div class="form__files-container">
                <span class="form__text">${uploadedFiles.name}</span>
            </div>
        `;
        this.$filesListContainer.append(content);
    }

    uploadFile() {
        if (this.fileList.length === 0) {
            this.showMessage('Please select a file before sending');
            return;
        }

        this.updateStatus('Processing');
        const formData = new FormData();
        formData.append('file', this.fileList[0].file);

        $.ajax({
            url: '/antivirus_bekend',
            method: 'POST',
            data: formData,
            processData: false,
            contentType: false,
            success: (response) => {
                this.handleUploadSuccess(response);
            }
        });
    }

    showMessage(message) {
        if (lang_global === "uk") {
            message_window('Виберіть файл перед відправкою');
        } else if (lang_global === "en") {
            message_window(message);
        }
    }

    updateStatus(status) {
        $("#unix_antivirus_098e3").html(`
            <p id="we332dvc"></p>
            <p id="we332dvc1"></p>
            <br>
            <p id="we332dvc2"></p>
            <p id="we332dvc3"></p>
            <p id="we332dvc4"></p>
            <br>
            <button id="d13dqe021w34fwvqedddd" style="display: none" onclick="window.dk_antivirus.download_json_data()" class="ewwdqszx4wefff"></button>    
        `);

        if (lang_global === "uk") {
            $('#we332dvc').html(status === 'Processing' ? "Обробка" : 'Завершено з помилкою');
        } else if (lang_global === "en") {
            $('#we332dvc').html(status === 'Processing' ? "Processing" : 'Completed with an error');
        }
    }

    handleUploadSuccess(response) {
        $('#d13dqe021w34fwvqedddd').hide();
        clean_div("f2ewds322r3345trg");
        this.processResponseData(response);
    }

    processResponseData(response) {
        response = JSON.parse(response);

        if (response['status'] === 2) {
            this.handleSuccessResponse(response);
        } else {
            this.handleErrorResponse();
        }
    }

    handleSuccessResponse(response) {
        if (lang_global === "uk") {
            $('#we332dvc1').html('Завершено успішно');
            $('#we332dvc3').html('Хеш файлу: ' + response['hash']);
        } else if (lang_global === "en") {
            $('#we332dvc1').html('Completed successfully');
            $('#we332dvc3').html('The hash of the file: ' + response['hash']);
        }

        console.log(response['data'])

        if (response['data'] !== "") {
            this.handleExeFile(response['data']);
        } else {
            this.displayExeError();
        }

        this.handleVirusResponse(response);
    }

    handleExeFile(data) {
        data_json_exe = JSON.parse(data);

        this.displayExeDetails(data_json_exe);
    }

    handleVirusResponse(response) {
        if (response['namber'] === 0) {
            this.displayNoVirusFound();
        } else {
            this.displayVirusFound(response['namber']);
        }
    }

    displayExeError() {
        if (lang_global === "uk") {
            $('#we332dvc4').html('Файл не є exe дані детальні дані отримати невдалось');
        } else if (lang_global === "en") {
            $('#we332dvc4').html('The file is not an exe data detailed data failed to get');
        }
    }

    displayExeDetails(data_json_exe) {
        if (lang_global === "uk") {
            $('#we332dvc4').html('Детальні дані про exe');
            $('#d13dqe021w34fwvqedddd').show().html('Скачати');
        } else if (lang_global === "en") {
            $('#we332dvc4').html('More details about the exe');
            $('#d13dqe021w34fwvqedddd').show().html("Download");
        }
    }

    displayNoVirusFound() {
        if (lang_global === "uk") {
            $('#we332dvc2').html('<span class="f343ffv1 grean_1">Вірусів незнайдено</span>');
        } else if (lang_global === "en") {
            $('#we332dvc2').html('<span class="f343ffv1 grean_1">No viruses found</span>');
        }
    }

    displayVirusFound(namber) {
        if (lang_global === "uk") {
            $('#we332dvc2').html(`<span class="f343ffv read">Вірус знайдено кількість ${namber}</span>`);
        } else if (lang_global === "en") {
            $('#we332dvc2').html(`<span class="f343ffv read">Virus found number ${namber}</span>`);
        }
    }

    handleErrorResponse() {
        if (lang_global === "uk") {
            $('#we332dvc1').html('Завершено з помилкою зачекайте декілька хвилин щоб ваш файл обробився і повторять запрос');
        } else if (lang_global === "en") {
            $('#we332dvc1').html('Completed with an error, please wait a few minutes for your file to be processed and the request will be repeated');
        }
    }

    download_json_data() {
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
}
