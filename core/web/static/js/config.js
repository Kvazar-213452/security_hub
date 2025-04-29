// core/static/js/config.js

let antivirus_flash_drive;
let lang_global;
let data_json_exe;
let reg_login;
let scan_virus_file;

// html
let html_1 = `
<div class="menu no_select_o">
    <div class="menu_logo"><img src="/static/img/9.png"><p>Security hub</p></div>
    <br><br>
    <div class="menu_div">
        <div onclick="page_iframe('/wifi', 'btn1')" id="btn1" class="button"><img src="/static/img/4.png"><p></p></div>
        <div onclick="page_iframe('/system', 'btn2')" id="btn2" class="button"><img src="/static/img/6.png"><p></p></div>
        <div onclick="page_iframe('/antivirus', 'btn3')" id="btn3" class="button"><img src="/static/img/7.png"><p></p></div>
        <div onclick="page_iframe('/cleaning', 'btn4')" id="btn4" class="button"><img src="/static/img/8.png"><p></p></div>
        <div onclick="page_iframe('/encryption', 'btn5')" id="btn5" class="button"><img src="/static/img/11.png"><p></p></div>
        <div onclick="page_iframe('/server', 'btn8')" id="btn8" class="button"><img src="/static/img/13.png"><p></p></div>
        <div onclick="page_iframe('/password', 'btn9')" id="btn9" class="button"><img src="/static/img/14.png"><p></p></div>
        <div onclick="page_iframe('/register', 'btn10')" id="btn10" class="button"><img src="/static/img/15.png"><p></p></div>
        <div onclick="page_iframe('/settings', 'btn6')" id="btn6" class="button"><img src="/static/img/3.png"><p></p></div>
        <br><br>
        <a href=""><div id="btn7" class="button"><img src="/static/img/12.png"><p></p></div></a>
    </div>
</div>

<div class="console"></div>
<iframe id="iframe"></iframe>
`;

// cleanup
let bg_color1 = "rgb(24, 24, 34)";
let bg_color2 = "rgb(89, 89, 89)";

// schedule
let main_color = "#68ff9d";
let text_color = "#ffffffd4";
