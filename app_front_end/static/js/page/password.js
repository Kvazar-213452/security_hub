function generatePassword() {
    const length = document.getElementById('passwordLength').value;
    const includeUppercase = document.getElementById('uppercase').checked;
    const includeNumbers = document.getElementById('numbers').checked;
    const includeSpecialChars = document.getElementById('specialChars').checked;

    const lowercaseLetters = 'abcdefghijklmnopqrstuvwxyz';
    const uppercaseLetters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
    const numbers = '0123456789';
    const specialChars = '!@#$%^&*()_+-=[]{}|;:,.<>?';

    let characters = lowercaseLetters;

    if (includeUppercase) {
        characters += uppercaseLetters;
    }

    if (includeNumbers) {
        characters += numbers;
    }

    if (includeSpecialChars) {
        characters += specialChars;
    }

    let password = '';
    for (let i = 0; i < length; i++) {
        const randomIndex = Math.floor(Math.random() * characters.length);
        password += characters[randomIndex];
    }

    $('.passw_wqdsc span').html(password);
}

function get_status_reg() {
    $.ajax({
        url: "/get_password",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            const parsedJson = JSON.parse(response);
            render_key(parsedJson['key']);
        }
    });
}

function add_key_pasw() {
    let data = {
        key: $('#password_d1qwasz').val(),
        pasw: $('#password_d1qwasz1').val()
    }

    $.ajax({
        url: "/add_key_pasw",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function (response) {

        }
    });
}

function render_key(data) {
    $("#pasw_pkjnf2qewvsd").html(null);

    for (let i = 0; i < data.length; i++) {
        let text = `
        <div class="div_pasw_user">
            <p class="table_pasw">${data[i][0]}</p>
            <p class="table_pasw1">${data[i][1]}</p>
            <p onclick="del_key('${data[i][0]}')" class="table_pasw2">del</p>
        </div>
        `;

        $("#pasw_pkjnf2qewvsd").append(text);
    }
}

function del_key(data) {
    $.ajax({
        url: "/del_key_pasw",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: data}),
        success: function (response) {
            get_status_reg();
        }
    });
}

function get_status_reg_hide_pasw() {
    $.ajax({
        url: "/reg_status",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            reg_login = response['acsses']

            if (reg_login == 0) {
                $("#pasword_1_btn_page1").hide();
            }
        }
    });
}

get_status_reg();