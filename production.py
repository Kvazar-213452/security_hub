import os
import shutil
import subprocess

# production.py

def build_and_deploy():
    production_path = ".production"

    os.makedirs(production_path, exist_ok=True)

    os.chdir("app_back_end")
    subprocess.run(["go", "build", "-ldflags=-H windowsgui"], check=True)

    app_back_end_path = os.path.join("..", production_path, "app_back_end")
    os.makedirs(app_back_end_path, exist_ok=True)
    shutil.move("head.exe", os.path.join(app_back_end_path, "head.exe"))

    for folder in ["library", "data"]:
        src = folder
        dest = os.path.join(app_back_end_path, folder)
        if os.path.exists(src):
            shutil.copytree(src, dest, dirs_exist_ok=True)

    os.chdir("..")

    for folder in ["app_front_end", "core", "data"]:
        src = folder
        dest = os.path.join(production_path, folder)
        if os.path.exists(src):
            shutil.copytree(src, dest, dirs_exist_ok=True)

    os.chdir("auto_update")
    subprocess.run(["go", "build", "-ldflags=-H windowsgui"], check=True)

    auto_update_path = os.path.join("..", production_path, "auto_update")
    os.makedirs(auto_update_path, exist_ok=True)
    shutil.move("head.exe", os.path.join(auto_update_path, "head.exe"))

    html_tmp_update_src = "html_tmp_update"
    html_tmp_update_dest = os.path.join(auto_update_path, "html_tmp_update")
    if os.path.exists(html_tmp_update_src):
        shutil.copytree(html_tmp_update_src, html_tmp_update_dest, dirs_exist_ok=True)

    os.chdir("..")

    nm3_path = os.path.join(production_path, "core", "NM3")
    if os.path.exists(nm3_path):
        shutil.rmtree(nm3_path)

    print("End prod")    

if __name__ == "__main__":
    build_and_deploy()
