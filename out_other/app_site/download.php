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
        <button onclick="openModal('modal1')" class="d3232f4ff">Скачати інсталятор</button>
        <a href="/script/count_zip.php"><button class="d3232f4ff">Скачати zip архів</button></a>

        <!-- modal -->

        <div id="modal1"> 
            <div id="modal-content">
                <p class="load_title">Повідомлення</p>
                <br>
                <p>Є 2 інсталятора на вибір. Версія інсталятора FISV більш стабільна</p>
                <br><br>
                <a href="/script/count_FISV.php"><button class="d3232f4ff fwevgrevf">FISV інсталятор</button></a>
                <a href="/script/count_NISV.php"><button class="d3232f4ff fwevgrevf">NISV інсталятор</button></a>
                <br><br><br><br>
                <p class="f34g45h43ew">Програма знаходиться в БЕТА тесті можливі проблеми з сервером</p>
                <br>
                <p>Дані для входу в бета-акаунт</p>
                <p>name = beta_test</p>
                <p>pasw = 123456_beta</p>
                <br>
                <button onclick="clos('modal1')" class="d3232f4ffddq">Закрити</button>
                <br>
            </div>
        </div>
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