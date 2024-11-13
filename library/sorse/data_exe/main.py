import pefile
import json
import sys

file = sys.argv[1]

def analyze_exe(file_path):
    analysis_data = {}

    try:
        pe = pefile.PE(file_path)

        analysis_data['file'] = file_path
        analysis_data['magic_number'] = hex(pe.DOS_HEADER.e_magic)
        analysis_data['file_type'] = pe.FILE_HEADER.Machine
        analysis_data['creation_time'] = pe.FILE_HEADER.TimeDateStamp
        analysis_data['section_count'] = len(pe.sections)

        sections = []
        for section in pe.sections:
            section_data = {
                'section_name': section.Name.decode().strip(),
                'section_size': section.SizeOfRawData,
                'section_virtual_address': hex(section.VirtualAddress),
                'section_characteristics': section.Characteristics
            }
            sections.append(section_data)
        analysis_data['sections'] = sections

        imports = []
        for entry in pe.DIRECTORY_ENTRY_IMPORT:
            library_data = {
                'library_name': entry.dll.decode(),
                'functions': [imp.name.decode() if imp.name else 'unknown' for imp in entry.imports]
            }
            imports.append(library_data)
        analysis_data['imports'] = imports

        with open('data/data_exe.json', 'w', encoding='utf-8') as json_file:
            json.dump(analysis_data, json_file, ensure_ascii=False, indent=4)

        print(f"Аналіз завершено. Результати збережено в 'data_exe.json'.")

    except Exception as e:
        print(f"Помилка під час аналізу файлу: {e}")

analyze_exe(file)