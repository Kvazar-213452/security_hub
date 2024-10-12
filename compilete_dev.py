import os

let = input("Type ")

if int(let) == 0:
    os.system("git add -A")
    name = input("Name: ")
    os.system(f'git commit -m "{name}"')
    os.system("git push")
elif int(let) == 1:
    os.system("node-sass web/static/prefab/main.scss web/static/css/main.css")
    os.system("node-sass web/static/prefab/global.scss web/static/css/global.css")
elif int(let) == 2:
    os.system("go build -o main.exe main.go")
    os.system(r"main.exe")