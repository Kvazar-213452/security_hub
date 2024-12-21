#!/bin/bash

echo "Type:"
read let

if [ "$let" -eq 0 ]; then
    git add -A
    echo "Name:"
    read name
    git commit -m "$name"
    git push
elif [ "$let" -eq 1 ]; then
    node-sass app_front_end/static/prefab/scss/main.scss app_front_end/static/css/main.css
    node-sass app_front_end/static/prefab/scss/global.scss app_front_end/static/css/global.css
elif [ "$let" -eq 2 ]; then
    pushd app_back_end > /dev/null
    go build
    ./head
    popd > /dev/null
elif [ "$let" -eq 3 ]; then
    pushd app_back_end > /dev/null
    go build -ldflags="-H windowsgui"
    ./head
    popd > /dev/null
elif [ "$let" -eq 4 ]; then
    pushd app_back_end > /dev/null
    rsrc -ico icon.ico -o icon.syso
    ./head
    popd > /dev/null
else
    echo "Invalid option!"
fi
