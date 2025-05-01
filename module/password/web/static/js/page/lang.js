// module/password/web/static/js/page/lang.js

let mas_lang = {
  "en": [
    ['#pasw_fwefc2dewsd span', 'Best passwords'],
    ['#pasw_dwqvnbb455', 'Password generation with subsequent storage in an encrypted database'],
    ['#pasword_1_btn_page', 'Generation'],
    ['#pasword_1_btn_page1', 'Cloud passwords'],
    ['#pasw_deqfwre', 'Password generation'],
    ['#pasw_pou823rd', 'Password length:'],
    ['#pasw_o218dqafghh', 'Include uppercase letters:'],
    ['#pasw_dq323fh9vvv', 'Include numbers:'],
    ['#pasw_o290rf323f', 'Include special characters:'],
    ['#pasw_dkq30f3q23r unix', 'Generated password:'],
    ['#pasw_p00o1edcc', 'Generate password'],
    ['#pasw_ok2f3web', 'Add password'],
    ['#pasw_d3vnb4354', 'Send']
  ],

  "uk": [
    ['#pasw_fwefc2dewsd span', 'Кращі паролі'],
    ['#pasw_dwqvnbb455', 'Генерація паролів з подальшим зберіганням в зашифрованій базі даних'],
    ['#pasword_1_btn_page', 'Генерація'],
    ['#pasword_1_btn_page1', 'Хмарні паролі'],
    ['#pasw_deqfwre', 'Генерація пароля'],
    ['#pasw_pou823rd', 'Довжина пароля:'],
    ['#pasw_o218dqafghh', 'Включити великі літери:'],
    ['#pasw_dq323fh9vvv', 'Включити цифри:'],
    ['#pasw_o290rf323f', 'Включити спеціальні символи:'],
    ['#pasw_dkq30f3q23r unix', 'Згенерований пароль:'],
    ['#pasw_p00o1edcc', 'Згенерувати пароль'],
    ['#pasw_ok2f3web', 'Додати пароль'],
    ['#pasw_d3vnb4354', 'Відправити']
  ]
};

function lang_change_page(lang) {
  for (let i = 0; i < mas_lang[lang].length; i++) {
    $(mas_lang[lang][i][0]).html(mas_lang[lang][i][1]);
  }
}
