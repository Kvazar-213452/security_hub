// module/server/web/static/js/page/lang.js

let mas_lang = {
  "en": [
    ['#lang_system_vfd8723ed span', 'System Data'],
    ['#lang_server_verdfvcww', 'Secure file transfer over the network'],
    ['#server_ddqd212dqwas', 'Select file'],
    ['#server_099kfvv', 'Password for file access:'],
    ['#server_Pfj23wes', 'Password:'],
    ['#server_0ij2p3ve', 'Search file by password'],
    ['#server_0jdccc', 'Hidden files'],
    ['#sever_pi913ofv', 'Upload to server'],
    ['#server_lang_werthrgjhfbvd', 'Number of uploaded files'],
    ['#duw1291291e12fq', 'Open website'],
    ['.server_0312edcccc', 'Download']
  ],

  "uk": [
    ['#lang_system_vfd8723ed span', 'Системні дані'],
    ['#lang_server_verdfvcww', 'Безпечна передача файлів через мережу'],
    ['#server_ddqd212dqwas', 'Виберіть файл'],
    ['#server_099kfvv', 'Пароль для доступу до файлу:'],
    ['#server_Pfj23wes', 'Пароль:'],
    ['#server_0ij2p3ve', 'Пошук файлу за паролем'],
    ['#server_0jdccc', 'Приховані файли'],
    ['#sever_pi913ofv', 'Завантаження на сервер'],
    ['#server_lang_werthrgjhfbvd', 'Кількість завантажених файлів'],
    ['#duw1291291e12fq', 'Відкрити сайт'],
    ['.server_0312edcccc', 'Завантажити'],
  ]
};

function lang_change_page(lang) {
  for (let i = 0; i < mas_lang[lang].length; i++) {
    $(mas_lang[lang][i][0]).html(mas_lang[lang][i][1]);
  }
}
