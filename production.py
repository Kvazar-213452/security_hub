import os
import shutil

# production.py

production_path = ".production"
core_path = os.path.join(production_path, "core")
des_source = os.path.join("core", "des")
des_target = os.path.join(core_path, "des")

if os.path.exists(production_path):
    shutil.rmtree(production_path)

os.makedirs(core_path, exist_ok=True)

items_to_copy = ["main.exe", "starter.md", "web", "data"]

for name in items_to_copy:
    src = os.path.join("core", name)
    dst = os.path.join(core_path, name)
    if os.path.exists(src):
        if os.path.isfile(src):
            shutil.copy2(src, dst)
        elif os.path.isdir(src):
            shutil.copytree(src, dst)
    else:
        print(f"[!] Не знайдено: {src}")

os.makedirs(des_target, exist_ok=True)

des_items = ["library", "web", "starter.md", "result.json", "head.exe"]

for name in des_items:
    src = os.path.join(des_source, name)
    dst = os.path.join(des_target, name)
    if os.path.exists(src):
        if os.path.isfile(src):
            shutil.copy2(src, dst)
        elif os.path.isdir(src):
            shutil.copytree(src, dst)
    else:
        print(f"[!] Не знайдено: {src}")
