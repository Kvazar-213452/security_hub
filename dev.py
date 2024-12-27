import subprocess
import os

def run_command(command):
    result = subprocess.run(command, shell=True, check=True, text=True)
    return result

let = int(input("Type: "))

# ⣾⡇⣿⣿⡇⣾⣿⣿⣿⣿⣿⣿⣿⣿⣄⢻⣦⡀⠁⢸⡌⠻⣿⣿⣿⡽⣿⣿
# ⡇⣿⠹⣿⡇⡟⠛⣉⠁⠉⠉⠻⡿⣿⣿⣿⣿⣿⣦⣄⡉⠂⠈⠙⢿⣿⣝⣿
# ⠤⢿⡄⠹⣧⣷⣸⡇⠄⠄⠲⢰⣌⣾⣿⣿⣿⣿⣿⣿⣶⣤⣤⡀⠄⠈⠻⢮
# ⠄⢸⣧⠄⢘⢻⣿⡇⢀⣀⠄⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⡀⠄⢀
# ⠄⠈⣿⡆⢸⣿⣿⣿⣬⣭⣴⣿⣿⣿⣿⣿⣿⣿⣯⠝⠛⠛⠙⢿⡿⠃⠄⢸
# ⠄⠄⢿⣿⡀⣿⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣿⣿⣿⣿⡾⠁⢠⡇⢀
# ⠄⠄⢸⣿⡇⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣏⣫⣻⡟⢀⠄⣿⣷⣾
# ⠄⠄⢸⣿⡇⠄⠈⠙⠿⣿⣿⣿⣮⣿⣿⣿⣿⣿⣿⣿⡿⢠⠊⢀⡇⣿⣿
# ⠒⠤⠄⣿⡇⢀⡲⠄⠄⠈⠙⠻⢿⣿⣿⠿⠿⠟⠛⠋⠁⣰⠇⠄⢸⣿⣿⣿

if let == 0:
    subprocess.run(["git", "add", "-A"])
    name = input("Name: ")
    subprocess.run(["git", "commit", "-m", name])
    subprocess.run(["git", "push"])
elif let == 1:
    subprocess.run(["sass", "app_front_end/static/prefab/scss/main.scss", "app_front_end/static/css/main.css"])
    subprocess.run(["sass", "app_front_end/static/prefab/scss/global.scss", "app_front_end/static/css/global.css"])
elif let == 2:
    try:
        os.chdir("app_back_end")
        subprocess.run(["go", "build"])
        subprocess.run(["./head.exe"])
    finally:
        os.chdir("../..")
elif let == 3:
    try:
        os.chdir("app_back_end")
        subprocess.run(["go", "build", "-ldflags='-H windowsgui'"])
        subprocess.run(["./head.exe"])
    finally:
        os.chdir("../..")
elif let == 4:
    try:
        os.chdir("app_back_end")
        subprocess.run(["rsrc", "-ico", "icon.ico", "-o", "icon.syso"])
        subprocess.run(["./head.exe"])
    finally:
        os.chdir("../..")
