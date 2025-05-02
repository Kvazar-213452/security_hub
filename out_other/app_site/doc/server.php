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
        <p class="rrrr23r2vvvvv">Server module</p>
        <br>
        <p>
        Модуль server дозволяє зберігати файли на віддаленому сервері з тимчасовим терміном зберігання до 24 годин. Кожен завантажений файл отримує унікальний ідентифікатор або ключ доступу. За допомогою цього ключа користувач може завантажити свій файл протягом встановленого часу. Після закінчення 24 годин файл автоматично видаляється з сервера. Модуль забезпечує базовий рівень безпеки та конфіденційності, дозволяючи доступ лише тим, хто має ключ. Він підтримує роботу з файлами різних форматів і розмірів у межах дозволеного ліміту. Такий підхід зручний для тимчасового обміну файлами без реєстрації.
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