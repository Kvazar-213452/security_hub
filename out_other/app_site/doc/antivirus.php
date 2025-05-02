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
        <p class="rrrr23r2vvvvv">Antivirus module</p>
        <br>
        <p>
        Модуль antivirus надає можливість сканувати файли на наявність шкідливого програмного забезпечення та шкідливих кодів. Він може перевіряти як окремі файли, так і цілі директорії чи шляхи на комп'ютері.  Інформація про поточні процеси, їхню активність, споживання ресурсів. Модуль допомагає користувачам захищати систему від вірусів, шкідливих програм і іншого небажаного ПО, забезпечуючи детальне сканування всіх компонентів комп'ютера.
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