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
        <p class="rrrr23r2vvvvv">Cleaning module</p>
        <br>
        <p>
        Модуль cleaning забезпечує ефективне очищення комп'ютера від різних тимчасових файлів та даних, що можуть накопичуватися з часом. Очищення DNS та ARP-таблиць дозволяє видаляти застарілі записи, що можуть спричиняти проблеми з мережевим з'єднанням. Він також очищає журнали файлової системи, що зберігають інформацію про операції з файлами, знижуючи ризик витоку конфіденційних даних.
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