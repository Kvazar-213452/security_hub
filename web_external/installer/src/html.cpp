#include <string>

std::string html_content_core = R"(
<style>
    iframe{
        position: fixed;
        height: 100%;
        width: 100%;
        top: 0%;
        left: 0%;
    }
</style>
<iframe src="http://127.0.0.1:59093/" frameborder="0"></iframe>
)";

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
        }

        body {
            background-color: #17181f;
            height: 100%;
            width: 100%;
            font-family: "Roboto", sans-serif;
            color: #dbdbdb;
        }

        .dqwdwdd4 {
            position: relative;
            text-align: center;
            font-weight: bold;
            font-size: 30px;
            color: #55c959;
            text-decoration: underline;
            text-decoration-thickness: 2px;
            text-underline-offset: 4px;
        }

        .dwqdwccccc {
            position: absolute;
            top: 50%;
            left: 50%;
            padding: 15px;
            transform: translate(-50%, -50%);
            height: 75%;
            width: 75%;
            background-color: #20252c;
        }

        .dqfeffggg {
            padding: 15px;
            font-size: 17px;
        }

        .ew32fe {
            display: block;
            margin: 0 auto;
            padding: 10px;
            font-weight: bold;
            color: #dbdbdb;
            background-color: #17181f;
            font-size: 18px;
            border: none;
            border: solid 2px #ffffff00;
            border-radius: 5px;
            transition: all .3s;
            cursor: pointer;
        }

        .ew32fe:hover {
            border: solid 2px #55c959;
            color: #55c959;
        }
    </style>
</head>
<body>
    <div class="dwqdwccccc">
        <p class="dqwdwdd4">Installer Security Hub 1.1</p>
        <br><br>
        <ul class="dqfeffggg">
            <li>Інсталюється ядро ASW для відображення програми вага 144КБ</li><br>
            <li>Створиться ярлик на робочому столі</li><br>
            <li>Файли будуть по шляху C:\security_hub</li><br>
            <li>Програма важить 50 МБ</li><br>
            <li>Інсталюється бази даних "web"</li><br>
            <li>Інсталюється бази даних "bekend"</li>
        </ul>
        <br><br>
        <button class="ew32fe">Інсталювати</button>    
    </div>
</body>
</html>
)";