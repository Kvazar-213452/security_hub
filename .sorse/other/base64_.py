# .sorse/other/base64_.py

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
print(f"file {input_file} in base Base64 {output_file}.")
