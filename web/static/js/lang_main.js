function lang_change_main(lang) {
    if (lang === "en") {
        $('#btn1 p').html("Wi-Fi");
        $('#btn2 p').html("System");
        $('#btn3 p').html("Antivirus");
        $('#btn4 p').html("Cleanup");
        $('#btn5 p').html("Encryption");
        $('#btn6 p').html("Settings");        
    } else if (lang === "uk") {
        $('#btn1 p').html("Вайфай");
        $('#btn2 p').html("Система");
        $('#btn3 p').html("Антивірус");
        $('#btn4 p').html("Очищення");
        $('#btn5 p').html("Шифрування");
        $('#btn6 p').html("Налаштування");
    }
}
