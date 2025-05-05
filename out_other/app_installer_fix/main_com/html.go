package main_com

var Html_code string = `
<html style="height: 100%">
<head>
    <meta charset="utf-8">
    <title>SPX</title>
    <style>
        .main_div {
            position: absolute;
            top: 50%;
            left: 50%;
            height: 70%;
            width: 60%;
            background-color: #181822;
            transform: translate(-50%, -50%);
            padding: 10px;
        }

        @font-face {
            font-family: "MyCustomFont";
            src: url("https://spx-security-hub.wuaze.com/static/Minecraft_1.1.ttf") format("truetype");
        }

        body {
            background-color: #22222e;
            font-family: "MyCustomFont", sans-serif;
            color: #fff;
            overflow-y: hidden;
            overflow-x: hidden;
            font-size: 16px;
        }

        #modal1{
        display: none;
        position: fixed;
        z-index: 10000;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
        overflow: auto;
        background-color: rgba(0, 0, 0, 0.4); 
        }

        #modal-content {
        position: fixed;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        width: 300px;
        height: 120px;
        overflow-y: hidden;
        overflow-x: hidden;
        background-color: #22222e;
        padding: 20px;
        }

        .load_title {
        position: relative;
        text-align: center;
        font-size: 20px;
        color: #68ff9d;
        }

        .load_animat_div {
        position: absolute;
        top: 55%;
        left: 50%;
        font-size: 18px;
        transform: translate(-50%, -50%);
        }

        .div_animation, .div_animation1, .div_animation2, .div_animation3 {
        width: 20px;
        height: 20px;
        background-color: #181822;
        float: left;
        margin-right: 20px;
        }

        .f54jo4evf {
            color: #68ff9d;
            position: relative;
            text-align: center;
            font-size: 22px;
        }

        .r4356ytgfvc {
            color: #68ff9d;
        }

        * {
            margin: 0;
            padding: 0;
            font-family: "MyCustomFont", sans-serif;
        }

        .dqfeffggg{
            position: relative;
            left: 20px;
            top: 5px;
            color: #ffffff;
        }

        .dqfeffggg li {list-style-type: square;}
        .dqfeffggg li::marker {color: #68ff9d; }

        .f3jf3ref {
            border: none;
            padding: 5px;
            font-size: 16px;
            color: #68ff9d;
            background-color: #22222e;
            cursor: pointer;
        }
    </style>
</head>
<body>
    <div class="main_div">
        <p class="f54jo4evf">Інсталяція програми</p>
        <br><br>
        <p>Опубліковано 09.11.2024</p>
        <p>Версія 9</p>
        <br>
        <p class="r4356ytgfvc">Інормація</p>
        <ul class="dqfeffggg">
            <li>Інсталюються оболонки NM1 NM2</li>
            <li>Створиться ярлик на робочому столі</li>
            <li>Файли будуть по шляху C:\security_hub</li>
            <li>Програма важить 14 МБ</li>
            <li>Програма з залежностями важить 61 МБ</li>
        </ul>
        <br><br>
        <button onclick="install()" class="f3jf3ref">Інсталювати</button>
        <!-- modal -->

        <div id="modal1"> 
            <div id="modal-content">
                <p class="load_title">Завантаження</p>
                <div class="load_animat_div">
                    <div class="div_animation"></div>
                    <div class="div_animation1"></div>
                    <div class="div_animation2"></div>
                    <div class="div_animation3"></div>
                </div>
            </div>
        </div>
    </div>
</body>
<script>
let animation = false;

function clos(name) {
    const el = document.getElementById(name);
    if (el) el.style.display = "none";
}

function openModal(name) {
    const el = document.getElementById(name);
    if (el) el.style.display = "block";
}

const divs = [
  document.querySelector('.div_animation'),
  document.querySelector('.div_animation1'),
  document.querySelector('.div_animation2'),
  document.querySelector('.div_animation3'),
];

let current = 0;
setInterval(() => {
    if (animation) {
        divs.forEach(div => div.style.backgroundColor = '#181822');
        divs[current].style.backgroundColor = '#68ff9d';
        current = (current + 1) % divs.length;
    }
}, 1000);

function install() {
    openModal("modal1");
    animation = true;

    fetch("/install", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(null)
    })
    .then(response => response.json())
    .then(data => {
        clos("modal1");
        animation = false;

        const mainDiv = document.querySelector(".main_div");
        if (mainDiv) mainDiv.innerHTML = "Виконано";

		end();
    })
    .catch(error => {
        console.error("Error during installation:", error);
    });
}

function end() {
    fetch("/off", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(null)
    })
    .then(response => response.json())
    .then(data => {

    })
    .catch(error => {
        console.error("Error during installation:", error);
    });
}
</script> 
</html>
`
