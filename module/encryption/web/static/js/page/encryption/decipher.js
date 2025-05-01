// app_front_end/static/js/page/encryption/decipher.js

class FileDecryption {
    constructor() {
        this.$fileInput = $('#file_dwqdw1');
        this.$filesListContainer = $('#files-list-container');
        this.$uploadContainer = $('#upload-container');
        this.langGlobal = lang_global;
        this.fileList = [];
    }

    init() {
        this.addFileInputEventListeners();
    }

    addFileInputEventListeners() {
        this.$fileInput.on('click dragstart dragover', this.activateInputContainer.bind(this));
        this.$fileInput.on('dragleave dragend drop change', this.deactivateInputContainer.bind(this));
        this.$fileInput.on('change', this.handleFileSelection.bind(this));
    }

    activateInputContainer() {
        this.$uploadContainer.addClass('active');
    }

    deactivateInputContainer() {
        this.$uploadContainer.removeClass('active');
    }

    handleFileSelection() {
        const files = Array.from(this.$fileInput[0].files);
        this.fileList = files.map(file => ({
            name: file.name,
            file: file
        }));

        this.updateFilesList();
    }

    updateFilesList() {
        this.$filesListContainer.empty();

        this.fileList.forEach(file => {
            const content = `
                <div class="form__files-container">
                    <span class="form__text">${file.name}</span>
                </div>
            `;
            this.$filesListContainer.append(content);
        });
    }

    start() {
        if (this.fileList.length === 0) {
            this.showMessage('No file selected', 'uk', 'en');
            return;
        }

        if (lang_global === "uk") {
            $('#bbbnsqee343').html("Обробка");
        } else if (lang_global === "en") {
            $('#bbbnsqee343').html("Обробка");           
        }

        const formData = new FormData();
        formData.append('file', this.fileList[0].file);
        formData.append('key', $('#cwwzevbnnn').val());
        formData.append('type', type_dec);
        console.log(type_dec)

        this.sendRequest(formData);
    }

    showProcessingMessage() {
        const message = this.langGlobal === "uk" ? "Обробка" : "Processing";
        $('#bbbnsqee343').html(message);
    }

    sendRequest(formData) {
        $.ajax({
            url: '/decipher_file',
            type: 'POST',
            processData: false,
            contentType: false,
            data: formData,
            success: function (response) {
                console.log(response)
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
            error: (jqXHR, textStatus, errorThrown) => {
                if (lang_global === "uk") {
                    message_window('Помилка');
                } else if (lang_global === "en") {
                    message_window('Error');                
                }
            }
        });
    }

    showMessage(messageKey, langUk, langEn) {
        const message = this.langGlobal === langUk ? messageKey : messageKey;
        message_window(message);
    }
}
