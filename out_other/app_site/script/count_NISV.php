<?php
$filename = "data/count_NISV.txt";

if (!file_exists($filename)) {
    file_put_contents($filename, "0");
}

$count = (int)file_get_contents($filename);
$count++;
file_put_contents($filename, $count);

header("Location: https://github.com/Kvazar-213452/data/raw/refs/heads/main/installer_app.exe");
exit;
?>
