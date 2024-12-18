from flask import Flask, request, jsonify
from flask_cors import CORS
import smtplib
from email.mime.text import MIMEText
import json
from Crypto.Cipher import AES
from Crypto.Util.Padding import unpad
from Crypto.Protocol.KDF import scrypt
import binascii

app = Flask(__name__)

key = b"3dp4g9DI8h7MzjVz3dp4g9DI8h7MzjVz"
iv = b"1234567890abcdef"

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
    
def decript(text):
    text = text.strip()

    if len(text) % 2 != 0:
        raise ValueError("Input string has an odd length, cannot unhexlify.")
    
    ciphertext = binascii.unhexlify(text)

    cipher = AES.new(key, AES.MODE_CBC, iv)
    plaintext = unpad(cipher.decrypt(ciphertext), AES.block_size)

    return plaintext.decode()

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

@app.route('/version', methods=['POST'])
def version():
    return jsonify({'version': config_data["version"]})

if __name__ == "__main__":
    CORS(app)
    app.run(debug=True)