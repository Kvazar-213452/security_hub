import base64

def convert_to_base64(input_file, output_file):
    with open(input_file, 'rb') as file:
        file_content = file.read()
        
        base64_encoded = base64.b64encode(file_content).decode('utf-8')

    with open(output_file, 'w') as file:
        file.write(base64_encoded)

input_file = 'icon.ico'
output_file = 'main.txt'

convert_to_base64(input_file, output_file)
print(f"Файл {input_file} успішно перетворено в Base64 і записано в {output_file}.")