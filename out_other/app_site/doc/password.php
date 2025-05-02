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
        <p class="rrrr23r2vvvvv">Password module</p>
        <br>
        <p>
        Модуль password дозволяє генерувати надійні та випадкові паролі з урахуванням різної довжини та складності. Користувач може створити пароль із використанням великих і малих літер, цифр та спеціальних символів. Після реєстрації в системі відкривається доступ до вбудованого менеджера паролів. Менеджер дозволяє зберігати згенеровані паролі у захищеному вигляді. 
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