import os

let = input("Type ")

os.system(f"gendef {let}.dll")
os.system(f"dlltool -v --dllname {let}.dll --output-lib {let}.a --input-def {let}.def")