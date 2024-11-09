#include <string>

std::string html_content = R"(
<!DOCTYPE html>
<html lang="uk">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Installer Security Hub</title>
    <link href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;700&display=swap" rel="stylesheet">
    <style>
        @import url('https://fonts.googleapis.com/css2?family=Roboto:wght@400;700&display=swap');

        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        a {
            color: #000000;
            text-decoration: none;
        }

        body {
            background-color: #202020;
            width: 500px;
            font-family: "Roboto", sans-serif;
            height: 600px;
            padding: 20px;
        }

        .ew32fe,
        .ew32fe1 {
            padding: 7px;
            background-color: #55c959;
            font-size: 16px;
            border: none;
            border-radius: 5px;
            transition: all .3s;
            cursor: pointer;
            position: relative;
            float: right;
            margin-right: 20px;
            left: 20px;
        }

        .ew32fe:hover,
        .ew32fe1:hover {opacity: 0.6;}

        .ew32fe1{background-color: #303030;}

        .ew32fe1 a{color: #ffffff}

        h1{
            color: #ffffff;
            position: relative;
            float: left;
            top: 15px;
        }

        .imfe{
            position: relative;
            float: right;
            height: 80px;
        }

        .dest{color: #55c959c8;}

        .wdqe333{color: #ffffff;}

        .dqfeffggg{
            position: relative;
            left: 15px;
            top: 5px;
            color: #ffffff;
            opacity: 0.5;
        }

        .edwe{color: #ffffff;}

        .e332s{color: #55c959;}
    </style>
</head>
<body>
    <h1>Інсталятор Security Hub</h1>
    <img class="imfe" src="http://spx-security-hub.wuaze.com/static/img/9.png">
    <br><br><br>
    <p class="dest">Опубліковано 09.11.2024</p>
    <p class="dest">Версія 3</p>
    <br>
    <div id="ddcbnxcew33333"></div>
</body>
</html>

<script>
function dwn() {
    fetch("/dwn", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(null)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`Помилка при відправці: ${response.statusText}`);
        }
        return response.text();
    })
    .then(text => {
        if (text === "0") {
            end_page();
        }
    })
    .catch(error => {
        console.error("Помилка при відправці:", error);
    });
}

function dwn_page() {
    let content = `
    <p class="edwe">Інсталяція ...</p>
    <p class="edwe">Інсталяція триватє в середньому 3 секунди</p>
    <br>
    <button class="ew32fe1"><a href="/exit">Закрити інсталятор</a></button>
    `;

    document.getElementById("ddcbnxcew33333").innerHTML = content;

    dwn();
}

function main_page() {
    let content = `
    <p class="wdqe333">Інормація</p>
    <ul class="dqfeffggg">
        <li>Інсталюються оболонки NM1 NM2</li>
        <li>Створиться ярлик на робочому столі</li>
        <li>Файли будуть по шляху C:\\security_hub</li>
        <li>Програма важить 14 МБ</li>
        <li>Програма з залежностями важить 61 МБ</li>
    </ul>
    <br>
    <button class="ew32fe1"><a href="/exit">Закрити інсталятор</a></button>
    <button onclick='dwn_page()' class="ew32fe"><a>Інсталювати</a></button>
    `;
    document.getElementById("ddcbnxcew33333").innerHTML = content;
}

function end_page() {
    let content = `
    <p class="e332s">Завершино</p>
    <p class="edwe">Можете закрити інсталятор</p>
    <br>
    <button class="ew32fe1"><a href="/exit">Закрити інсталятор</a></button>
    `;

    document.getElementById("ddcbnxcew33333").innerHTML = content;
}

main_page();
</script>
)";