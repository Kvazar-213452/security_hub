import psutil
import tkinter as tk
from tkinter import ttk

# Функція для отримання інформації про процеси
def get_processes_info():
    processes = []
    for proc in psutil.process_iter(['pid', 'name', 'cpu_percent', 'memory_info']):
        try:
            process_info = proc.info
            processes.append((
                process_info['pid'],
                process_info['name'],
                f"{process_info['cpu_percent']}%",
                f"{process_info['memory_info'].rss / 1024 / 1024:.2f} MB"
            ))
        except (psutil.NoSuchProcess, psutil.AccessDenied, psutil.ZombieProcess):
            pass  # Ігноруємо процеси, до яких немає доступу
    return processes

# Функція для оновлення таблиці
def update_table():
    # Очистити таблицю перед оновленням
    for row in tree.get_children():
        tree.delete(row)

    # Отримати нову інформацію про процеси і додати їх у таблицю
    processes = get_processes_info()
    for process in processes:
        tree.insert("", "end", values=process)

    # Запустити оновлення кожні 3 секунди
    root.after(3000, update_table)

# Створення основного вікна програми
root = tk.Tk()
root.title("Диспетчер задач")

# Створення таблиці для відображення процесів
tree = ttk.Treeview(root, columns=("PID", "Process Name", "CPU Usage", "Memory Usage"), show="headings")
tree.heading("PID", text="PID")
tree.heading("Process Name", text="Назва процесу")
tree.heading("CPU Usage", text="Використання CPU")
tree.heading("Memory Usage", text="Використання пам'яті")

# Відображення таблиці на вікні
tree.pack(fill=tk.BOTH, expand=True)

# Запуск першого оновлення таблиці
update_table()

# Запуск графічного інтерфейсу
root.mainloop()
