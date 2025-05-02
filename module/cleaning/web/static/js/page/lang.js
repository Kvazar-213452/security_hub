// module/cleaning/web/static/js/page/lang.js

let mas_lang = {
  "en": [
    ['#vc728i32000 span', 'PC Cleanup'],
    ['#bvh9iuweddd', 'Unfortunately, the script does not clean the System32 folder. Sorry'],
    ['#bx66210doddddw', 'Cleanup Information'],
    ['#bo9ho2_0dddd', 'Clear DNS and ARP tables'],
    ['#wifi_0jdccc', 'Information'],
    ['#d_odw2983fff', 'File System Logs'],
    ['#b8cq82d3333', 'Remove Open Path History'],
    ['#vc972odwe9', 'Clear Internet Cache'],
    ['#b89bdq9822222', 'Remove Specific Files'],
    ['#bc92idw092222d', 'Remove Temporary Files, Caches, and Logs'],
    ['#vwer43erfdfcff', 'Additional Settings'],
    ['#bvpuwe033230_233', 'Remove Windows Backups'],
    ['#v21qp9wd000222s', 'The "Remove Windows Backups" setting deletes saved system backups created for restoring Windows in case of failure or data loss. Backups may include system images, restore points, and automatic file archives.'],
    ['.cbq98222fz233', 'Install'],
    ['#cbwq982fff45t5', 'Remove all Wi-Fi profiles'],
    ['#cbqw278d436666', 'The "Remove all Wi-Fi profiles" setting deletes saved wireless network profiles your device has connected to.'],
    ['#bv89qwdasdwededddd', 'Remove Remote Desktop Connection settings.'],
    ['#vb0928923ee3ddd', 'The "Remove Remote Desktop Connection settings" deletes saved configurations and connection history for remote desktops.'],
    ['#bviubiwdd3333r', 'Restart DOSKEY Command Processor'],
    ['#bciu292ed45675663', 'The "Restart DOSKEY Command Processor" setting updates the DOSKEY command processor.\n<br><br>\nRestarting DOSKEY clears the command history and resets all defined macros. This can be useful for resetting the command line environment to its initial state.'],
    ['#bf0qwp32r4t5651222', 'Clean Computer'],
    ["#modla_install_32ffff", "Data processing"]
  ],

  "uk": [
    ['#vc728i32000 span', 'Очищення ПК'],
    ['#wifi_0jdccc', 'Інформація'],
    ['#bvh9iuweddd', 'На жаль, скрипт не очищує папку System32. Вибачте'],
    ['#bx66210doddddw', 'Інформація очистки'],
    ['#bo9ho2_0dddd', 'Очищення DNS та ARP-таблиць'],
    ['#d_odw2983fff', 'Журнали файлової системи'],
    ['#b8cq82d3333', 'Видалення історії відкритих шляхів'],
    ['#vc972odwe9', 'Очищення кешу інтернету'],
    ['#b89bdq9822222', 'Видалення конкретних файлів'],
    ['#bc92idw092222d', 'Видалення тимчасових файлів, кешів та логів'],
    ['#vwer43erfdfcff', 'Додаткові налаштування'],
    ['#bvpuwe033230_233', 'Видалити резервні копії Windows'],
    ['#v21qp9wd000222s', 'Налаштування "Видалити резервні копії Windows" видаляє збережені копії системи, створені для відновлення Windows у разі збоїв або втрати даних. Резервні копії можуть включати зображення системи (резервні копії системи), точки відновлення, а також автоматичні архіви файлів.'],
    ['.cbq98222fz233', 'Встановити'],
    ['#cbwq982fff45t5', 'Видалити всі профілі Wi-Fi'],
    ['#cbqw278d436666', 'Налаштування "Видалити всі профілі Wi-Fi" видалить збережені профілі бездротових мереж, до яких ваше пристрій підключалось.'],
    ['#bv89qwdasdwededddd', 'Видалити налаштування підключення до віддаленого робочого столу.'],
    ['#vb0928923ee3ddd', 'Налаштування "Видалити налаштування підключення до віддаленого робочого столу" видаляє збережені конфігурації та історію підключень до віддалених робочих столів.'],
    ['#bviubiwdd3333r', 'Перезавантажити процесор команд doskey'],
    ['#bciu292ed45675663', 'Налаштування "Перезавантажити процесор команд doskey" оновлює процесор команд doskey.\n<br><br>\nПерезавантаження doskey очищає історію команд і скидає всі визначені макроси. Це може бути корисно для скидання середовища командного рядка до його початкового стану.'],
    ['#bf0qwp32r4t5651222', 'Очищення комп\'ютера'],
    ["#modla_install_32ffff", "Data processing"]
  ]
};

function lang_change_page(lang) {
  for (let i = 0; i < mas_lang[lang].length; i++) {
    $(mas_lang[lang][i][0]).html(mas_lang[lang][i][1]);
  }
}
