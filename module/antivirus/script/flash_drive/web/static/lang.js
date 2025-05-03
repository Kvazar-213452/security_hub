let mas_lang = {
    "en": [
        ["#modla_install_32ffff", "Data processing"],
        ["#f23f34trtyhy566hp", "Scanned data"],
        ["#f32pol24vnza", "Scanned files"],
        ["#fpkpzj021r23ovew", "Malicious files"]
    ],

    "uk": [
        ["#modla_install_32ffff", "Обробка даних"],
        ["#f23f34trtyhy566hp", "Проскановані дані"],
        ["#f32pol24vnza", "Відскановані файли"],
        ["#fpkpzj021r23ovew", "Шкідливі файли"]
    ]
};
  
function lang_change_page(lang) {
    for (let i = 0; i < mas_lang[lang].length; i++) {
        $(mas_lang[lang][i][0]).html(mas_lang[lang][i][1]);
    }
}
  