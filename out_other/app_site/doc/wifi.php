<?php
include $_SERVER['DOCUMENT_ROOT'] . '/config.php'; 
?>
<!DOCTYPE html>
<html lang="en">
<head>
    <?php echo $content; ?>
    <link rel="stylesheet" href="<?php echo $domen_name; ?>/style/main">
</head>
<body>
    <?php require '../page/menu.php'; ?>
    <br><br><br><br><br><br><br><br><br><br>

    <div class="w3fwe">
        <p class="rrrr23r2vvvvv">Wi-Fi module</p>
        <br>
        <p>
        Модуль Wi-Fi для Windows дозволяє отримувати інформацію про доступні бездротові мережі, їхні SSID, рівень сигналу, типи шифрування та канали. Він може взаємодіяти з мережевим інтерфейсом через системні API або утиліти, як-от netsh або PowerShell-команди. За допомогою цього модуля можна дізнатись, до якої мережі підключений комп’ютер, а також отримати MAC-адресу точки доступу. Деякі реалізації модуля дозволяють моніторити трафік або перехоплювати пакети за допомогою додаткових бібліотек, як-от WinPcap або Npcap. Також можна аналізувати параметри бездротового інтерфейсу в реальному часі. Модуль може бути корисним для діагностики з'єднання або аналізу безпеки мережі.
        </p>
    </div>

    <?php require '../page/footer.php'; ?>
<br><br><br>
    <script src="<?php echo $domen_name; ?>/lib/js/jquery"></script>
    <script src="<?php echo $domen_name; ?>/js/main"></script>

    <script>
        select_button(0);
    </script>
</body>
</html>