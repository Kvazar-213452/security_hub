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
        <a href="<?php echo $server_name; ?>/installer_online.exe"><button class="d3232f4ff">Скачати офлайн інсталятор</button></a>
        <a href="<?php echo $server_name; ?>/installer_ofline.exe"><button class="d3232f4ff">Скачати онлайн інсталятор</button></a>
        <a href="<?php echo $server_name; ?>/main.zip"><button class="d3232f4ff">Скачати zip архів</button></a>
    </div>
    <?php require 'page/footer.php'; ?>
    <br><br>

    <script src="<?php echo $domen_name; ?>/lib/js/jquery"></script>
    <script src="<?php echo $domen_name; ?>/js/main"></script>
</body>
</html>