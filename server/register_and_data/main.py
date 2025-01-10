from flask import Flask, request, jsonify
from flask_cors import CORS
import smtplib
from email.mime.text import MIMEText
import json
from main_com.func import save_to_db, decript

app = Flask(__name__)

with open("config.json", "r") as file:
    config_data = json.load(file)

sender = config_data["sender"]
password = config_data["password"]

server = smtplib.SMTP("smtp.gmail.com", 587)
server.starttls()

def send_text_email(message, subject, receiver):
    try:
        server.login(sender, password)
        msg = MIMEText(message)
        msg["Subject"] = subject
        msg["From"] = sender
        msg["To"] = receiver
        print(receiver)
        server.sendmail(sender, receiver, msg.as_string())
        return "The text message was sent successfully!"
    except Exception as _ex:
        return f"Error: {_ex}"

@app.route('/send_email', methods=['POST'])
def send_email():
    data = request.json
    print(data)
    try:
        receiver = decript(data.get('receiver'))
        print(receiver)
        code = decript(data.get('code'))
        subject = "Notification"
        message = f"Code: {code}"
        result = send_text_email(message, subject, receiver)
        return jsonify({'status': 'success', 'message': result})
    except ValueError as e:
        return jsonify({'status': 'error', 'message': str(e)})

@app.route('/save_user', methods=['POST'])
def save_user():
    try:
        user_data = request.get_json()

        if save_to_db(user_data):
            return jsonify({"message": "1"}), 200
        else:
            return jsonify({"message": "0"}), 200

    except Exception as e:
        return jsonify({"message": "Помилка при збереженні даних", "error": str(e)}), 500

@app.route('/version', methods=['POST'])
def version():
    return jsonify({'version': config_data["version"]})

@app.route('/check', methods=['POST'])
def check():
    return jsonify({'version': None})

@app.route('/get_password', methods=['POST'])
def get_password():
    try:
        user_data = request.get_json()
        gmail = user_data.get("gmail")

        if not gmail:
            return jsonify({"message": "Gmail не надано"}), 400

        with open('db.json', 'r') as f:
            db_data = json.load(f)

        for entry in db_data:
            if entry.get("gmail") == gmail:
                return jsonify({"key": entry.get("key")}), 200

        return jsonify({"message": "Користувача з таким gmail не знайдено"}), 404

    except Exception as e:
        return jsonify({"message": "Помилка при обробці запиту", "error": str(e)}), 500
    
@app.route('/add_key_pasw', methods=['POST'])
def add_key_pasw():
    try:
        user_data = request.get_json()
        gmail = user_data.get("gmail")
        key = user_data.get("key")
        pasw = user_data.get("pasw")

        if not gmail:
            return jsonify({"message": "Gmail не надано"}), 400

        with open('db.json', 'r') as f:
            db_data = json.load(f)

        for entry in db_data:
            if entry.get("gmail") == gmail:
                if isinstance(entry.get("key"), list):
                    entry["key"].append([key, pasw])
                else:
                    entry["key"] = [[key, pasw]]

                with open('db.json', 'w') as f:
                    json.dump(db_data, f, indent=4)

                return jsonify({"message": "Дані успішно оновлені", "key": entry.get("key")}), 200

        return jsonify({"message": "Користувача з таким gmail не знайдено"}), 404

    except Exception as e:
        return jsonify({"message": "Помилка при обробці запиту", "error": str(e)}), 500
    
@app.route('/login', methods=['POST'])
def login():
    try:
        user_data = request.get_json()
        name = user_data.get("name")
        password = user_data.get("password")

        if not name:
            return jsonify({"message": "Gmail не надано"}), 200

        with open('db.json', 'r') as f:
            db_data = json.load(f)

        for entry in db_data:
            if entry.get("name") == name and entry.get("pasw") == password:
                return jsonify({"status": "1"}), 200

        return jsonify({"status": "0"}), 200

    except Exception as e:
        return jsonify({"status": "0", "error": str(e)}), 200
    
@app.route('/del_key_pasw', methods=['POST'])
def del_key_pasw():
    try:
        print(request.get_json())
        user_data = request.get_json()
        gamil = user_data.get("key")
        value = user_data.get("pasw")

        if not gamil or not value:
            return jsonify({"message": "Gmail або value не надано"}), 200

        with open('db.json', 'r') as f:
            db_data = json.load(f)

        for entry in db_data:
            if entry.get("gmail") == gamil:
                key_data = entry.get("key", [])

                key_data = [item for item in key_data if item[0] != value]

                entry["key"] = key_data

                with open('db.json', 'w') as f:
                    json.dump(db_data, f, indent=4)

                return jsonify({"status": "1"}), 200

        return jsonify({"status": "0"}), 200

    except Exception as e:
        return jsonify({"status": "0"}), 500

if __name__ == "__main__":
    CORS(app)
    app.run(debug=True)