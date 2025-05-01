// app_front_end/static/js/page/lang.js

let mas_lang = {
  "en": [
    ['#bnoewfcc_i033e span', 'Encrypt File'],
    ['#bvoif2032991367', 'Don’t lose your keys'],
    ['.brbtedfe4443', 'Encrypt File'],
    ['.brbtedfe44431', 'Decrypt File'],
    ['.vb0f0234456465', 'Send']
  ],

  "uk": [
    ['#bnoewfcc_i033e span', 'Шифрування файлу'],
    ['#bvoif2032991367', 'Не втратьте ключі'],
    ['.brbtedfe4443', 'Зашифрувати файл'],
    ['.brbtedfe44431', 'Розшифрувати файл'],
    ['.vb0f0234456465', 'Відправити']
  ]
};

function lang_change_page(lang) {
  for (let i = 0; i < mas_lang[lang].length; i++) {
    $(mas_lang[lang][i][0]).html(mas_lang[lang][i][1]);
  }
}
