$let = Read-Host "Type"

if ($let -eq 0) {
    git add -A
    $name = Read-Host "Name"
    git commit -m $name
    git push
} elseif ($let -eq 1) {
    node-sass web/static/prefab/main.scss web/static/css/main.css
    node-sass web/static/prefab/global.scss web/static/css/global.css
} elseif ($let -eq 2) {
    go build
    .\head.exe
} elseif ($let -eq 3) {
    go build -ldflags="-H windowsgui"
    .\head.exe
} elseif ($let -eq 4) {
    cd web_external
}
