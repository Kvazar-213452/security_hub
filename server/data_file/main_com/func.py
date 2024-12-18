import os
import json
from datetime import datetime, timedelta
from threading import Timer

UPLOAD_FOLDER = os.path.join(os.getcwd(), 'static/file')
DB_FILE = os.path.join(os.getcwd(), 'db.json')

def get_time():
    now = datetime.now() + timedelta(hours=1)
    return now.strftime('%H:%M:%S')

def read_db():
    if not os.path.exists(DB_FILE):
        return []
    with open(DB_FILE, 'r', encoding='utf-8') as f:
        return json.load(f)

def write_db(data):
    with open(DB_FILE, 'w', encoding='utf-8') as f:
        json.dump(data, f, ensure_ascii=False, indent=2)

def delete_file_after_delay(file_path, password, delay=43200):
    def delete():
        if os.path.exists(file_path):
            os.remove(file_path)
            print(f"Файл {file_path} видалено.")

        db = read_db()
        updated_db = [entry for entry in db if entry['password'] != password]
        write_db(updated_db)
        print(f"Запис із паролем '{password}' видалено з db.json.")

    Timer(delay, delete).start()