function lang_change_main(lang) {
  if (lang === "en") {
    $('#bt11n1 p').html("Wi-Fi");
    $('#bt11n2 p').html("System");
    $('#bt11n3 p').html("Antivirus");
    $('#bt11n4 p').html("Cleanup");
    $('#bt11n5 p').html("Encryption");
    $('#bt11n6 p').html("Settings");
    $('#bt11n7 p').html("Reboot");
    $('#bt11n8 p').html("Server");
    $('#bt11n9 p').html("Passwords");
    return $('#bt11n10 p').html("File system");
  } else if (lang === "uk") {
    $('#bt11n1 p').html("Вайфай");
    $('#bt11n2 p').html("Система");
    $('#bt11n3 p').html("Антивірус");
    $('#bt11n4 p').html("Очищення");
    $('#bt11n5 p').html("Шифрування");
    $('#bt11n6 p').html("Налаштування");
    $('#bt11n7 p').html("Перезавантажити");
    $('#bt11n8 p').html("Сервер");
    $('#bt11n9 p').html("Паролі");
    return $('#bt11n10 p').html("Файлова система");
  }
};