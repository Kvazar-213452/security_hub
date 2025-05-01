// module/wifi/web/static/js/page/lang.js

let mas_lang = {
  "en": [
    ['#lang_wefsdeeeeee span', 'Wi-Fi Manager'],
    ['#lang_nfefdfdvghyt', 'Get more information about your Wi-Fi'],
    ['#lang_dqwedddcwww', "Connection ="],
    ['#wifi_0jdccc', 'Information'],
    ['#wifi_0jdccc1', 'Packets']
  ],

  "uk": [
    ['#lang_wefsdeeeeee span', 'Вайфай менеджер'],
    ['#lang_nfefdfdvghyt', 'Отримайте більше інформації про ваш вайфай'],
    ['#lang_dqwedddcwww', "З'єднання ="],
    ['#wifi_0jdccc', 'Інформація'],
    ['#wifi_0jdccc1', 'Пакети']
  ]
};

function lang_change_page(lang) {
  for (let i = 0; i < mas_lang[lang].length; i++) {
    $(mas_lang[lang][i][0]).html(mas_lang[lang][i][1]);
  }
}
