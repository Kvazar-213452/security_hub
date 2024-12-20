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
           console.log(response)
        }
    });
}
