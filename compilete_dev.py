import os
import subprocess

let = input("Type: ")

if let == "0":
    os.system("git add -A")
    name = input("Name: ")
    os.system(f'git commit -m "{name}"')
    os.system("git push")

elif let == "1":
    os.system("sass app_front_end/static/prefab/scss/main.scss app_front_end/static/css/main.css")
    os.system("sass app_front_end/static/prefab/scss/global.scss app_front_end/static/css/global.css")

elif let == "2":
    try:
        os.chdir("app_back_end")
        subprocess.run(["go", "build"])
        subprocess.run(["./head.exe"])
    finally:
        os.chdir("..")

elif let == "3":
    try:
        os.chdir("app_back_end")
        subprocess.run(["go", "build", "-ldflags=-H windowsgui"])
        subprocess.run(["./head.exe"])
    finally:
        os.chdir("..")

elif let == "4":
    try:
        os.chdir("app_back_end")
        subprocess.run(["rsrc", "-ico", "icon.ico", "-o", "icon.syso"])
        subprocess.run(["./head.exe"])
    finally:
        os.chdir("..")
