import subprocess
import threading
import os

# start.py

def run_process(command, cwd):
    subprocess.run(command, shell=True, cwd=cwd)

def main():
    try:
        thread1 = threading.Thread(target=run_process, args=("python main.py", "out_other/server/register_and_data"))
        thread1.start()

        thread2 = threading.Thread(target=run_process, args=("npm i && npm start", "out_other/server/data_file"))
        thread2.start()

        thread1.join()
        thread2.join()

    finally:
        os.chdir(os.getcwd())

if __name__ == "__main__":
    main()
