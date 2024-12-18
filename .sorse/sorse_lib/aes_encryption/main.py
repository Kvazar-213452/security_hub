from Crypto.Cipher import AES
from Crypto.Util.Padding import unpad
from Crypto.Protocol.KDF import scrypt
import binascii

# Текст, який потрібно розшифрувати (ваш зашифрований текст)
ciphertext = binascii.unhexlify("0e4c3de92f2fefdeea147b5c3fdffb50db68491006218e64dbff5f5af6bac29e")

# Ваш ключ та IV
key = b"3dp4g9DI8h7MzjVz3dp4g9DI8h7MzjVz"  # 32 байти
iv = b"1234567890abcdef"  # 16 байт

# Створення AES шифратора з режимом CBC
cipher = AES.new(key, AES.MODE_CBC, iv)

# Розшифровка та видалення відступів (PKCS7 padding)
plaintext = unpad(cipher.decrypt(ciphertext), AES.block_size)

# Перетворення в текстовий формат
print("Decrypted text:", plaintext.decode())
