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
    <div class="bfievfdwere">
        <div id="weg4gwevds" class="wedwffcsd"></div>
    </div>
    <br><br>
    <div>
        <input type="text" id="myInput" placeholder="Ведіть імя">
        <br>
        <textarea id="myTextarea" placeholder="Ведіть повідмлення"></textarea>
        <br><br>
        <button id="sendButton">Send Data</button>
    </div>
    </div>
    <?php require 'page/footer.php'; ?>
    <br><br>
    <script src="<?php echo $domen_name; ?>/lib/js/jquery"></script>
    <script src="<?php echo $domen_name; ?>/js/main"></script>
    <script src="<?php echo $domen_name; ?>/js/massage"></script>
    <script>
        select_button(3);
        fetchData();
    </script>
</body>
</html>