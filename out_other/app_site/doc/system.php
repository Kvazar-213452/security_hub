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
        <p class="rrrr23r2vvvvv">System module</p>
        <br>
        <p>
        Модуль system надає загальну інформацію про систему Windows, включаючи дані про операційну систему, апаратне забезпечення та ресурси. Він дозволяє дізнатись назву комп’ютера, версію Windows та архітектуру процесора. Також відображається обсяг оперативної пам’яті та її поточне використання. Можна отримати дані про процесор, кількість ядер і логічних потоків. 
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