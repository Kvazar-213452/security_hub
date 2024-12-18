$let = Read-Host "Type"

# ⣾⡇⣿⣿⡇⣾⣿⣿⣿⣿⣿⣿⣿⣿⣄⢻⣦⡀⠁⢸⡌⠻⣿⣿⣿⡽⣿⣿
# ⡇⣿⠹⣿⡇⡟⠛⣉⠁⠉⠉⠻⡿⣿⣿⣿⣿⣿⣦⣄⡉⠂⠈⠙⢿⣿⣝⣿
# ⠤⢿⡄⠹⣧⣷⣸⡇⠄⠄⠲⢰⣌⣾⣿⣿⣿⣿⣿⣿⣶⣤⣤⡀⠄⠈⠻⢮
# ⠄⢸⣧⠄⢘⢻⣿⡇⢀⣀⠄⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⡀⠄⢀
# ⠄⠈⣿⡆⢸⣿⣿⣿⣬⣭⣴⣿⣿⣿⣿⣿⣿⣿⣯⠝⠛⠛⠙⢿⡿⠃⠄⢸
# ⠄⠄⢿⣿⡀⣿⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣿⣿⣿⣿⡾⠁⢠⡇⢀
# ⠄⠄⢸⣿⡇⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣏⣫⣻⡟⢀⠄⣿⣷⣾
# ⠄⠄⢸⣿⡇⠄⠈⠙⠿⣿⣿⣿⣮⣿⣿⣿⣿⣿⣿⣿⣿⡿⢠⠊⢀⡇⣿⣿
# ⠒⠤⠄⣿⡇⢀⡲⠄⠄⠈⠙⠻⢿⣿⣿⠿⠿⠟⠛⠋⠁⣰⠇⠄⢸⣿⣿⣿

if ($let -eq 0) {
    git add -A
    $name = Read-Host "Name"
    git commit -m $name
    git push
} elseif ($let -eq 1) {
    node-sass app_front_end/static/prefab/scss/main.scss app_front_end/static/css/main.css
    node-sass app_front_end/static/prefab/scss/global.scss app_front_end/static/css/global.css
} elseif ($let -eq 2) {
    Push-Location
    try {
        cd app_back_end
        go build
        .\head.exe
    } finally {
        Pop-Location
    }
} elseif ($let -eq 3) {
    Push-Location
    try {
        cd app_back_end
        go build -ldflags="-H windowsgui"
        .\head.exe
    } finally {
        Pop-Location
    }
} elseif ($let -eq 4) {
    Push-Location
    try {
        cd app_back_end
        rsrc -ico icon.ico -o icon.syso
        .\head.exe
    } finally {
        Pop-Location
    }
} elseif ($let -eq 5) {
    lsc -o app_front_end/static/js/ app_front_end/static/prefab/LiveScript/lang_main.ls
    lsc -o app_front_end/static/js/page/ app_front_end/static/prefab/LiveScript/lang.ls

    python .sorse/other/del_ls_txt.py app_front_end/static/js/lang_main.js
    python .sorse/other/del_ls_txt.py app_front_end/static/js/page/lang.js
} elseif ($let -eq 6) {
    Push-Location
    try {
        cd auto_update
        go build -ldflags="-H windowsgui"
        .\head.exe
    } finally {
        Pop-Location
    }
} elseif ($let -eq 7) {
    Push-Location
    try {
        cd auto_update
        go build
        .\head.exe
    } finally {
        Pop-Location
    }
}
