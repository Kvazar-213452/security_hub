let visualization_mas = ['visualization1', 'visualization2'];
let vsgretdbgc = ['vsgretdbgc1', 'vsgretdbgc2'];
let shell_NM = ['shell_NM', 'shell_NM1'];
let frg45th9nd = ['frg45th9nd', 'frg45th9nd1', 'frg45th9nd2'];
let setingss_vdwewe = ['setingss_vdwewe', 'setingss_vdwewe1'];
let mmain_buuton = ['btn1', 'btn2', 'btn3', 'btn4', 'btn5', 'btn6', 'btn7', 'btn8', 'btn9', 'btn10'];
let vvw2311323ferererg3g3g3 = ['vvw2311323ferererg3g3g3', 'vvw2311323ferererg3g3g31'];

let mas_sonar = [
    ['dwdc21e12d', 'dwdc21e12d1', 'dwdc21e12d2'],
    ['wifi_0jdccc']
]

const unsafeProtocols = ["WEP", "WPA", "HTTP", "FTP", "Telnet", "RDP", "SNMP", "ICMP"];

let antivirus_flash_drive;
let lang_global;
let data_json_exe;

let data_cleaning = {
    backup: 0,
    wifi: 0,
    desktop: 0,
    doskey: 0
};

// html
let html_1 = `
<div class="menu">
    <div class="menu_logo"><img src="/static/img/9.png"><p>Alter++</p></div>
    <br><br>
    <div class="menu_div">
        <div onclick="page_iframe('/wifi', 'btn1')" id="btn1" class="button"><img src="/static/img/4.png"><p></p></div>
        <div onclick="page_iframe('/system', 'btn2')" id="btn2" class="button"><img src="/static/img/6.png"><p></p></div>
        <div onclick="page_iframe('/antivirus', 'btn3')" id="btn3" class="button"><img src="/static/img/7.png"><p></p></div>
        <div onclick="page_iframe('/cleaning', 'btn4')" id="btn4" class="button"><img src="/static/img/8.png"><p></p></div>
        <div onclick="page_iframe('/encryption', 'btn5')" id="btn5" class="button"><img src="/static/img/11.png"><p></p></div>
        <div onclick="page_iframe('/server', 'btn8')" id="btn8" class="button"><img src="/static/img/13.png"><p></p></div>
        <div onclick="page_iframe('/password', 'btn9')" id="btn9" class="button"><img src="/static/img/14.png"><p></p></div>
        <div onclick="page_iframe('/file_system', 'btn10')" id="btn10" class="button"><img src="/static/img/15.png"><p></p></div>
        <div onclick="page_iframe('/settings', 'btn6')" id="btn6" class="button"><img src="/static/img/3.png"><p></p></div>
        <br><br>
        <a href=""><div id="btn7" class="button"><img src="/static/img/12.png"><p></p></div></a>
    </div>
</div>

<div class="console"></div>
<iframe id="iframe"></iframe>
`;