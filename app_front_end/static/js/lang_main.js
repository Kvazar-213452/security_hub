// app_front_end/static/js/lang_main.js

function lang_change_main(lang) {
  if (lang === "en") {
    $('#btn1 p').html("Wi-Fi");
    $('#btn2 p').html("System");
    $('#btn3 p').html("Antivirus");
    $('#btn4 p').html("Cleanup");
    $('#btn5 p').html("Encryption");
    $('#btn6 p').html("Settings");
    $('#btn7 p').html("Reboot");
    $('#btn8 p').html("Server");
    $('#btn9 p').html("Passwords");
    $('#btn10 p').html("Register");
  } else if (lang === "uk") {
    $('#btn1 p').html("Вайфай");
    $('#btn2 p').html("Система");
    $('#btn3 p').html("Антивірус");
    $('#btn4 p').html("Очищення");
    $('#btn5 p').html("Шифрування");
    $('#btn6 p').html("Налаштування");
    $('#btn7 p').html("Перезавантажити");
    $('#btn8 p').html("Сервер");
    $('#btn9 p').html("Паролі");
    $('#btn10 p').html("Реєстрація");
  }
};
