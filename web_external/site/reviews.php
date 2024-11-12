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
    <div class="bfievfdwere">

    </div>

    <input type="text" id="myInput" placeholder="Enter title...">
    <textarea id="myTextarea" placeholder="Enter some data..."></textarea>
    <button id="sendButton">Send Data</button>

    <?php require 'page/footer.php'; ?>
    <br><br>
    <script src="<?php echo $domen_name; ?>/lib/js/jquery"></script>
    <script src="<?php echo $domen_name; ?>/js/main"></script>

    <script>
        select_button(3);
    </script>

<script>
              $(document).ready(function() {
            $('#sendButton').click(function() {
                const inputData = $('#myInput').val();
                const textData = $('#myTextarea').val();

                $.ajax({
                    url: 'http://localhost/save_data.php',
                    type: 'POST',
                    data: { title: inputData, text: textData },
                    success: function(response) {
                        alert('Data saved successfully!');
                    },
                    error: function() {
                        alert('Error saving data.');
                    }
                });
            });

            function fetchData() {
                $.ajax({
                    url: 'http://localhost/db.json',
                    type: 'GET',
                    dataType: 'json',
                    success: function(data) {
                        console.log(data);
                    },
                    error: function() {
                        console.error('Error fetching data from db.json');
                    }
                });
            }

            fetchData();

        });
    </script>
</body>
</html>