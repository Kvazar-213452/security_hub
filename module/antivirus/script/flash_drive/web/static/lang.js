let mas_lang = {
    "en": [
        ["#modla_install_32ffff", "Data processing"]
    ],

    "uk": [
        ["#modla_install_32ffff", "Обробка даних"]
    ]
};
  
function lang_change_page(lang) {
    for (let i = 0; i < mas_lang[lang].length; i++) {
        $(mas_lang[lang][i][0]).html(mas_lang[lang][i][1]);
    }
}
  