import os

os.system('go build -ldflags="-H windowsgui"')
os.system(r"head.exe")