import json
import binascii
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad

# server/register_and_data/main_com/func.py

key = b"3dp4g9DI8h7MzjVz3dp4g9DI8h7MzjVz"
iv = b"1234567890abcdef"

def is_data_exists(data, db_data):
    for existing_data in db_data:
        if existing_data == data:
            return True
    return False

def save_to_db(data):
    try:
        with open('db.cos', 'r') as f:
            db_data = json.load(f)
    except (FileNotFoundError, json.JSONDecodeError):
        db_data = []

    data['key'] = []

    if not is_data_exists(data, db_data):
        db_data.append(data)

        with open('db.cos', 'w') as f:
            json.dump(db_data, f, indent=4)
        return True
    else:
        return False
    
def decript(text):
    text = text.strip()

    if len(text) % 2 != 0:
        raise ValueError("error")
    
    ciphertext = binascii.unhexlify(text)

    cipher = AES.new(key, AES.MODE_CBC, iv)
    plaintext = unpad(cipher.decrypt(ciphertext), AES.block_size)

    return plaintext.decode()

def encrypt_file():
    try:
        with open("db.cos", 'r', encoding='utf-8') as file:
            data = json.load(file) 
    except UnicodeDecodeError as e:
        print(f"Error{e}")
        return
    except json.JSONDecodeError as e:
        print(f"Error parsing JSON: {e}")
        return

    cipher = AES.new(key, AES.MODE_CBC, iv)
    json_data = json.dumps(data)
    padded_data = pad(json_data.encode('utf-8'), AES.block_size)

    encrypted_data = cipher.encrypt(padded_data)

    with open("db.cos", 'wb') as file:
        file.write(encrypted_data)
    print(f"File db.cos encrypted successfully.")

def decrypt_file():
    try:
        with open("db.cos", 'rb') as file:
            encrypted_data = file.read()
    except Exception as e:
        print(f"Error reading file: {e}")
        return

    cipher = AES.new(key, AES.MODE_CBC, iv)
    decrypted_data = unpad(cipher.decrypt(encrypted_data), AES.block_size)

    try:
        decrypted_json = decrypted_data.decode('utf-8')
        data = json.loads(decrypted_json)
    except UnicodeDecodeError as e:
        print(f"Error decoding JSON: {e}")
        return
    except json.JSONDecodeError as e:
        print(f"Error parsing decrypted JSON: {e}")
        return

    with open("db.cos", 'w', encoding='utf-8') as file:
        json.dump(data, file, ensure_ascii=False, indent=4)
    print(f"File db.cos decrypted successfully.")
