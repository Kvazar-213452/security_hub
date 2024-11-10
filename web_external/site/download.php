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
    <div class="rrggef">
        <br><br><br><br>
        <br><br><br>
        <div class="fwf32rggg atvImg">
            <div class="atvImg-layer" data-img="<?php echo $domen_name; ?>/static/img/9.png"></div>
        </div>
        <br><br><br><br>
        <a href="https://github.com/Kvazar-213452/security_hub/raw/refs/heads/main/web_external/data/installer.exe"><button class="d3232f4ff">Скачати інсталятор</button></a>
        <a href="https://github.com/Kvazar-213452/security_hub/raw/refs/heads/main/web_external/data/main.zip"><button class="d3232f4ff">Скачати zip архів</button></a>
    </div>
    <?php require 'page/footer.php'; ?>
    <br><br>

    <script src="<?php echo $domen_name; ?>/lib/js/jquery"></script>
    <script src="<?php echo $domen_name; ?>/js/main"></script>

    <script>
        select_button(2);
    </script>
</body>
</html>