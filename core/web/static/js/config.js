// core/static/js/config.js

let lang_global;
let reg_login;
let db_lang = [];
let mmain_buuton = ['btn0', 'btn1', 'btn2', 'btn3', 'btn4', 'btn5', 'btn6', 'btn7', 'btn8', 'btn9', 'btn10', 'btn11', 'btn12'];
let lang_db = {
    "uk": {
        "relon": "Перезавантажити",
        "btn11": "Реєстарція",
        "btn12": "Налаштування"
    },
    "en": {
        "relon": "Reboot",
        "btn11": "register",
        "btn12": "settings"
    }
}

// html
let html_1 = `
<div class="menu no_select_o">
    <div class="menu_logo"><img src="/static/img/9.png"><p>Security hub</p></div>
    <br><br>
    <div class="menu_div">
        <div onclick="module_integrated('register', 'btn11')" id="btn11" class="button"><img src="/static/img/15.png"><p></p></div>
        <div onclick="module_integrated('settings', 'btn12')" id="btn12" class="button"><img src="/static/img/3.png"><p></p></div>
        <br><br>
        <a href=""><div id="relon" class="button"><img src="/static/img/12.png"><p></p></div></a>
    </div>
</div>

<div class="console"></div>
<iframe id="iframe"></iframe>
`;
