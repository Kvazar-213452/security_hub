$let = Read-Host "Type"

if ($let -eq 0) {
    git add -A
    $name = Read-Host "Name"
    git commit -m $name
    git push
} elseif ($let -eq 1) {
    node-sass web/static/prefab/scss/main.scss web/static/css/main.css
    node-sass web/static/prefab/scss/global.scss web/static/css/global.css
} elseif ($let -eq 2) {
    go build
    .\head.exe
} elseif ($let -eq 3) {
    go build -ldflags="-H windowsgui"
} elseif ($let -eq 4) {
    rsrc -ico icon.ico -o icon.syso
} elseif ($let -eq 5) {
    lsc -o web/static/js/ web/static/prefab/LiveScript/lang_main.ls
    lsc -o web/static/js/page/ web/static/prefab/LiveScript/lang.ls

    python .sorse/del_ls_txt.py web/static/js/lang_main.js
    python .sorse/del_ls_txt.py web/static/js/page/lang.js
}
