// app_front_end/static/js/page/password/generate.js

function generatePassword() {
    const length = $('#passwordLength').val();
    const includeUppercase = $('#uppercase').prop('checked');
    const includeNumbers = $('#numbers').prop('checked');
    const includeSpecialChars = $('#specialChars').prop('checked');

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
