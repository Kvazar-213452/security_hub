import os

let = input("Type ")

if int(let) == 0:
    os.system("git add -A")
    name = input("Name: ")
    os.system(f'git commit -m "{name}"')
    os.system("git push")
elif int(let) == 1:
    os.system("node-sass front_end/static/prefab/main.scss front_end/static/css/main.css")
    os.system("node-sass front_end/static/prefab/global.scss front_end/static/css/global.css")
elif int(let) == 2:
    os.system("go build -o main.exe main.go")
    os.system(r"main.exe")