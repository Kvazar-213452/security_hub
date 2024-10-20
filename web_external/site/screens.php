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
    <div class="rrggef">
        <img class="dqwd23d3ff" src="<?php echo $domen_name; ?>/static/img/4.jpg">
        <img class="dqwd23d3ff" src="<?php echo $domen_name; ?>/static/img/2.jpg">
        <img class="dqwd23d3ff" src="<?php echo $domen_name; ?>/static/img/3.jpg">
        <img class="dqwd23d3ff" src="<?php echo $domen_name; ?>/static/img/4.jpg">
        <img class="dqwd23d3ff" src="<?php echo $domen_name; ?>/static/img/5.jpg">
        <img class="dqwd23d3ff" src="<?php echo $domen_name; ?>/static/img/2.jpg">
    </div>
    <?php require 'page/footer.php'; ?>
    <br><br>
    <script src="<?php echo $domen_name; ?>/lib/js/jquery"></script>
    <script src="<?php echo $domen_name; ?>/js/main"></script>
</body>
</html>