import json
import binascii
from Crypto.Cipher import AES
from Crypto.Util.Padding import unpad

key = b"3dp4g9DI8h7MzjVz3dp4g9DI8h7MzjVz"
iv = b"1234567890abcdef"

def is_data_exists(data, db_data):
    for existing_data in db_data:
        if existing_data == data:
            return True
    return False

def save_to_db(data):
    try:
        with open('db.json', 'r') as f:
            db_data = json.load(f)
    except (FileNotFoundError, json.JSONDecodeError):
        db_data = []

    if not is_data_exists(data, db_data):
        db_data.append(data)

        with open('db.json', 'w') as f:
            json.dump(db_data, f, indent=4)
        return True
    else:
        return False
    
def decript(text):
    text = text.strip()

    if len(text) % 2 != 0:
        raise ValueError("Input string has an odd length, cannot unhexlify.")
    
    ciphertext = binascii.unhexlify(text)

    cipher = AES.new(key, AES.MODE_CBC, iv)
    plaintext = unpad(cipher.decrypt(ciphertext), AES.block_size)

    return plaintext.decode()