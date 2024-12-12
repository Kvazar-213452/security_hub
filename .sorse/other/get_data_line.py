import os
import json

def count_lines_in_file(file_path):
    try:
        with open(file_path, 'r', encoding='utf-8') as file:
            return sum(1 for line in file if line.strip())
    except Exception as e:
        print(f"Не вдалося відкрити файл {file_path}: {e}")
        return 0

def scan_directory(directory, extensions_to_check, ignore_files):
    total_lines = {key: 0 for key in extensions_to_check}
    total_count = 0

    for root, dirs, files in os.walk(directory):
        for file in files:
            file_name = os.path.basename(file)

            if file_name in ignore_files:
                continue

            file_extension = file.split('.')[-1]
            
            for language, extensions in extensions_to_check.items():
                if file_extension in extensions:
                    file_path = os.path.join(root, file)
                    lines_in_file = count_lines_in_file(file_path)
                    total_lines[language] += lines_in_file
                    total_count += lines_in_file

    return total_lines, total_count

def load_config(config_path):
    try:
        with open(config_path, 'r', encoding='utf-8') as config_file:
            return json.load(config_file)
    except Exception as e:
        print(f"Не вдалося завантажити конфігураційний файл: {e}")
        return {}

def main():
    config_path = 'config.json'
    config = load_config(config_path)

    if config:
        directory = input("Введіть шлях до каталогу для сканування: ")
        ignore_files = config.get("ignire", [])
        result, total_count = scan_directory(directory, config, ignore_files)

        if total_count > 0:
            print("Статистика по рядках коду:")
            for language, lines in result.items():
                if lines > 0:
                    percentage = (lines / total_count) * 100
                    print(f"{language.capitalize()}: {lines} рядків ({percentage:.2f}%)")

        else:
            print("У проекті немає рядків коду.")
    else:
        print("Не вдалося завантажити конфігурацію.")

if __name__ == "__main__":
    main()
