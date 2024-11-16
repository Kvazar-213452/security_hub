import re
import sys

file_1 = sys.argv[1]

def remove_function_patterns(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        content = file.read()

    content = re.sub(r'\(function\(\)\{', '', content)
    content = re.sub(r'\}\)\.call\(this\);', '', content)

    with open(file_path, 'w', encoding='utf-8') as file:
        file.write(content)


remove_function_patterns(file_1)