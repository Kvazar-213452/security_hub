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
    <?php require 'page/menu.php'; ?>
    <br><br><br><br><br><br><br>
    <div class="w3fwe">
        <div class="style_dqw">
            <p class="dqd313qe3qdw22">white style</p>
            <p class="dq3233rrffff">Вага: 1 6.97 КБ</p>
            <p class="dq3233rrffff">Створив: Kvazar-213452</p>
            <br>
            <a download href="<?php echo $domen_name; ?>/static/css/main_jr.css"><button>Завантажити</button></a>
        </div>
    </div>
    <?php require 'page/footer.php'; ?>
    <br><br>
    <script src="<?php echo $domen_name; ?>/lib/js/jquery"></script>
    <script src="<?php echo $domen_name; ?>/js/main"></script>

    <script>
        select_button(4);
    </script>
</body>
</html>