import GPUtil

# Отримання списку GPU
gpus = GPUtil.getGPUs()

# Перебір кожного GPU та вивід температури
for gpu in gpus:
    gpu_temperature = f"{gpu.temperature} °C"
    print(gpu_temperature)