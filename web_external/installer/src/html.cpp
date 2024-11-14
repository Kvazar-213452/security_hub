#include <string>

std::string html_content = R"(
<!DOCTYPE html>
<html lang="en">
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
            overflow-y: hidden;
            overflow-x: hidden;
        }

        .ew32fe,
        .ew32fe1 {
            padding: 7px;
            background-color: #766aff;
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

        .ew32fe1 {background-color: #303030;}

        .ew32fe1 a {color: #ffffff;}

        h1 {
            color: #ffffff;
            position: relative;
            float: left;
            top: 15px;
        }

        .imfe {
            position: relative;
            float: right;
            height: 80px;
        }

        .dest {color: #766affc8;}

        .wdqe333 {color: #ffffff;}

        .dqfeffggg {
            position: relative;
            left: 15px;
            top: 5px;
            color: #ffffff;
            opacity: 0.5;
        }

        .edwe {color: #ffffff;}

        .e332s {color: #766aff;}
    </style>
</head>
<body>
    <h1>Installer Security Hub</h1>
    <img class="imfe" src="https://raw.githubusercontent.com/Kvazar-213452/security_hub/refs/heads/main/web/static/img/9.png">
    <br><br><br>
    <p class="dest">Published on 09.11.2024</p>
    <p class="dest">Version 3</p>
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
            throw new Error(`Error on send: ${response.statusText}`);
        }
        return response.text();
    })
    .then(text => {
        if (text === "0") {
            end_page();
        }
    })
    .catch(error => {
        console.error("Error on send:", error);
    });
}

function dwn_page() {
    let content = `
    <p class="edwe">Installing ...</p>
    <p class="edwe">Installation takes about 3 seconds</p>
    <br>
    <button class="ew32fe1"><a href="/exit">Close Installer</a></button>
    `;

    document.getElementById("ddcbnxcew33333").innerHTML = content;

    dwn();
}

function main_page() {
    let content = `
    <p class="wdqe333">Information</p>
    <ul class="dqfeffggg">
        <li>Shells NM1 NM2 are being installed</li>
        <li>A shortcut will be created on the desktop</li>
        <li>Files will be located at C:\\security_hub</li>
        <li>The program size is 14 MB</li>
        <li>Program with dependencies is 61 MB</li>
    </ul>
    <br>
    <button class="ew32fe1"><a href="/exit">Close Installer</a></button>
    <button onclick='dwn_page()' class="ew32fe"><a>Install</a></button>
    `;
    document.getElementById("ddcbnxcew33333").innerHTML = content;
}

function end_page() {
    let content = `
    <p class="e332s">Completed</p>
    <p class="edwe">You may close the installer</p>
    <br>
    <button class="ew32fe1"><a href="/exit">Close Installer</a></button>
    `;

    document.getElementById("ddcbnxcew33333").innerHTML = content;
}

main_page();
</script>
)";