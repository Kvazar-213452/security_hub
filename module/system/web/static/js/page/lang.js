// module/system/web/static/js/page/lang.js

let mas_lang = {
  "en": [
    ['#lang_system_vfd8723ed span', 'System Data'],
    ['#lang_system_verdfvcww', 'Get more information about BIOS and Host'],
    ['#lang_system_wdeewds', 'BIOS Data'],
    ['#lang_system_wfr839wefsff', 'Host Data'],
    ['#lang_system_v00qwdweee', 'Open Programs'],
    ['#system_0jdccc', 'System Information'],
    ['#system_0jdccc1', 'File System'],
    ['#system_fqfewvsfdwq', 'Scan Settings'],
    ['#system_0dqdwas', 'Extension selection:'],
    ['#system_dqqdwas444', 'Send'],
    ['#systme_pj2eqdwa35tg', 'Send Data'],
    ['#systme_fqwfegwrehj1', 'Information']
  ],

  "uk": [
    ['#lang_system_vfd8723ed span', 'Системні дані'],
    ['#lang_system_verdfvcww', 'Отримайте більше інформації про bios та host'],
    ['#lang_system_wdeewds', 'Дані біоса'],
    ['#lang_system_wfr839wefsff', 'Дані хоста'],
    ['#lang_system_v00qwdweee', 'Відкриті програми'],
    ['#system_0jdccc', 'Системна інформація'],
    ['#system_0jdccc1', 'Файлова система'],
    ['#system_fqfewvsfdwq', 'Налаштування сканування'],
    ['#system_0dqdwas', 'Вибір розширення:'],
    ['#system_dqqdwas444', 'Відправити'],
    ['#systme_pj2eqdwa35tg', 'Відправити дані'],
    ['#systme_fqwfegwrehj1', 'Інформація']
  ]
};

function lang_change_page(lang) {
  for (let i = 0; i < mas_lang[lang].length; i++) {
    $(mas_lang[lang][i][0]).html(mas_lang[lang][i][1]);
  }
}
