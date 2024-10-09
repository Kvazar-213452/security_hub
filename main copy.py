import psutil
from tabulate import tabulate

# Отримання списку всіх процесів
def get_processes_info():
    processes = []
    for proc in psutil.process_iter(['pid', 'name', 'cpu_percent', 'memory_info']):
        try:
            # Отримуємо інформацію про процес
            process_info = proc.info
            processes.append([
                process_info['pid'],
                process_info['name'],
                f"{process_info['cpu_percent']}%",
                f"{process_info['memory_info'].rss / 1024 / 1024:.2f} MB"
            ])
        except (psutil.NoSuchProcess, psutil.AccessDenied, psutil.ZombieProcess):
            pass  # Ігноруємо процеси, які не можемо отримати

    return processes

# Виведення інформації про процеси в таблиці
def display_task_manager():
    processes = get_processes_info()
    headers = ["PID", "Process Name", "CPU Usage", "Memory Usage"]

    print(tabulate(processes, headers, tablefmt="pretty"))

if __name__ == '__main__':
    display_task_manager()
