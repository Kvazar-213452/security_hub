import os

ignored_extensions = ['.h', '.exe', '.txt', '.syso', '.zip', '.i', '.cc', '.css', 'dll', 'png', 'jpg', 'ico', 'pyc', 'toc'
                      'pkg,' 'pyz', ]


def count_lines_in_files(directory_path):
    total_lines = 0

    for root, dirs, files in os.walk(directory_path):
        for file in files:
            file_extension = os.path.splitext(file)[1].lower()

            if file_extension not in ignored_extensions:
                file_path = os.path.join(root, file)
                try:
                    with open(file_path, 'r', encoding='utf-8', errors='ignore') as f:
                        lines = f.readlines()
                        total_lines += len(lines)
                except Exception as e:
                    print(f"Не вдалося відкрити файл {file_path}: {e}")

    return total_lines

directory_path = r'C:\Users\god19\Desktop\security_hub'
result = count_lines_in_files(directory_path)

print(f"Загальна кількість рядків в усіх файлах: {result}")