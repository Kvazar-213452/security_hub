// app_front_end/static/js/page/lang.js

let mas_lang = {
  "en": [
    ['#antivirus_errfee2 span', 'Antivirus'],
    ['#antivirus_vef093f', 'Stay informed and secure'],
    ['#dwdc21e12d', 'Check the website'],
    ['#dwdc21e12d1', 'Check the file'],
    ['#dwdc21e12d2', 'In the background BETA'],
    ['#dwdc21e12d3', 'Scan Path'],
    ['#dwdc21e12d4', "Resource"],
    ['#antivirus_894534ffvvv', 'Check'],
    ['#antivirus_894534ffvvv55555', 'Check'],
    ['#h3ruiwefer24f', 'Description'],
    ['#fewrvw243rgefvcc', 'Features in the background are automatically activated at startup'],
    ['#vbs612dwes655', 'USB Flash Drive Monitoring'],
    ['#vb92354gu04ttg', 'This setting activates the USB media monitoring feature, which warns the user when new flash drives are plugged in. After enabling this feature, the program constantly scans available USB media and reports on infected files.'],
    ['#bg_dqwderfd', 'Set'],
    ['#upload-button', 'Quick Check'],
    ["#vb92354gu04ttg1", "After setting the value, you must restart the module to start monitoring."]
  ],

  "uk": [
    ['#antivirus_errfee2 span', 'Антивірус'],
    ['#antivirus_vef093f', 'Будьте в інформаційній безпеці'],
    ['#dwdc21e12d', 'Перевірити сайт'],
    ['#dwdc21e12d1', 'Перевірити файл'],
    ['#dwdc21e12d2', 'На фоні BETA'],
    ['#dwdc21e12d3', 'Сканування шляху'],
    ['#dwdc21e12d4', "Ресурси"],
    ['#antivirus_894534ffvvv', 'Перевірити'],
    ['#antivirus_894534ffvvv55555', 'Перевірити'],
    ['#h3ruiwefer24f', 'Опис'],
    ['#fewrvw243rgefvcc', 'Функції на фоні вмикаються автоматично при запуску'],
    ['#vbs612dwes655', 'Моніторинг флешок'],
    ['#vb92354gu04ttg', 'Це налаштування активує функцію моніторингу USB-носіїв, яка попереджає користувача при підключенні нових флешок. Після ввімкнення цієї функції програма постійно сканує доступні USB-носії і повідомляє про заражені файли.'],
    ['#bg_dqwderfd', 'Встановити'],
    ['#upload-button', 'Швидка перевірка'],
    ["#vb92354gu04ttg1", "Після встановлення значення необхідно перезапустити модуль, щоб почати моніторинг."]
  ]
};

function lang_change_page(lang) {
  for (let i = 0; i < mas_lang[lang].length; i++) {
    $(mas_lang[lang][i][0]).html(mas_lang[lang][i][1]);
  }
}
