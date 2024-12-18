import os
import time
from flask import Flask, request, send_from_directory
from werkzeug.utils import secure_filename
from flask_cors import CORS
from main_com.func import get_time, read_db, write_db, delete_file_after_delay, UPLOAD_FOLDER

app = Flask(__name__)
CORS(app)

app.config['UPLOAD_FOLDER'] = UPLOAD_FOLDER
os.makedirs(UPLOAD_FOLDER, exist_ok=True)

@app.route('/')
def index():
    return 'unix server'

@app.route('/upload', methods=['POST'])
def upload_file():
    if 'file' not in request.files or 'password' not in request.form:
        return 'Невірний запит. Перевірте, чи передано файл та пароль.', 400

    file = request.files['file']
    password = request.form['password']
    if file.filename == '':
        return 'Файл не вибрано.', 400

    original_name = secure_filename(file.filename)
    unique_suffix = f'renamed-{int(time.time())}{os.path.splitext(original_name)[1]}'
    new_file_path = os.path.join(app.config['UPLOAD_FOLDER'], unique_suffix)

    file.save(new_file_path)
    current_time = get_time()

    db = read_db()
    db.append({'password': password, 'fileName': unique_suffix, 'time': current_time})
    write_db(db)

    delete_file_after_delay(new_file_path, password)

    return 'good'

@app.route('/search', methods=['POST'])
def search_file():
    data = request.get_json()
    if not data or 'searchPassword' not in data:
        return 'Невірний запит. Перевірте, чи передано пароль.', 400

    search_password = data['searchPassword']
    db = read_db()
    found_entry = next((entry for entry in db if entry['password'] == search_password), None)

    if found_entry:
        file_path = os.path.join(app.config['UPLOAD_FOLDER'], found_entry['fileName'])
        if os.path.exists(file_path):
            return send_from_directory(app.config['UPLOAD_FOLDER'], found_entry['fileName'], as_attachment=True)
        else:
            return 'Файл не знайдено на сервері.', 404

    return 'Пароль не знайдено.', 404


if __name__ == '__main__':
    app.run(port=3000, debug=True)