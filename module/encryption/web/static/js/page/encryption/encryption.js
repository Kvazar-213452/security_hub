// app_front_end/static/js/page/encryption/encryption.js

class FileEncryption {
    constructor() {
        this.fileInput = $('#file_dwqdw');
        this.filesListContainer = $('#files-list-container');
        this.uploadContainer = $('#upload-container');
        this.langGlobal = lang_global;
        this.fileList = [];
        this.resultContainer = $('#errewfffee');
        this.statusContainer = $('#bbbnsqee343');
    }

    init() {
        this.addFileInputEventListeners();
    }

    addFileInputEventListeners() {
        this.fileInput.on('click dragstart dragover', () => {
            this.uploadContainer.addClass('active');
        });

        this.fileInput.on('dragleave dragend drop change', () => {
            this.uploadContainer.removeClass('active');
            const files = Array.from(this.fileInput[0].files);
            this.handleFileSelection(files);
        });
    }

    handleFileSelection(files) {
        this.fileList = files.map(file => ({
            name: file.name,
            file: file
        }));

        this.updateFilesList();
    }

    updateFilesList() {
        this.filesListContainer.html('');
        this.fileList.forEach(file => {
            const content = `
                <div class="form__files-container">
                    <span class="form__text">${file.name}</span>
                </div>
            `;
            this.filesListContainer.append(content);
        });
    }

    startEncryption() {
        if (this.fileList.length === 0) {
            this.showMessage('No file selected');
            return;
        }

        this.showProcessingMessage();

        const formData = new FormData();
        formData.append('file', this.fileList[0].file);
        formData.append('filename', this.fileList[0].name);
        formData.append('type', type);

        this.sendRequest(formData);
    }

    showProcessingMessage() {
        if (this.langGlobal === 'uk') {
            this.statusContainer.html('Обробка');
        } else if (this.langGlobal === 'en') {
            this.statusContainer.html('Processing');
        }
    }

    sendRequest(formData) {
        $.ajax({
            url: '/encryption_file',
            type: 'POST',
            processData: false,
            contentType: false,
            data: formData,
            success: (response) => this.handleResponse(response)
        });
    }

    handleResponse(response) {
        this.showMessage('Успішно');

        this.encryptionFileEnd(response);

        const link = document.createElement('a');
        link.href = '/static/data/main.enc';
        link.download = 'main.enc';
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    }

    showMessage(messageKey) {
        if (this.langGlobal === 'uk') {
            message_window(messageKey);
        } else if (this.langGlobal === 'en') {
            message_window(messageKey);
        }
    }

    encryptionFileEnd(response) {
        this.resultContainer.html(`
            <p class="deferghhhh777">Your decryption key:</p>
            <br>
            <span class="dwedffvvvvv">${response}</span>
        `);
    }
}
