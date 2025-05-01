<?php
if ($_SERVER['REQUEST_METHOD'] == 'POST') {
    $title = $_POST['title'];
    $text = $_POST['text'];
    $filePath = 'db.json';
    $dataArray = file_exists($filePath) ? json_decode(file_get_contents($filePath), true) : [];
    $newData = [
        'title' => $title,
        'text' => $text
    ];
    $dataArray[] = $newData;
    file_put_contents($filePath, json_encode($dataArray, JSON_PRETTY_PRINT));
    echo json_encode(["status" => "success"]);
} else {
    echo json_encode(["status" => "error", "message" => "Invalid request method"]);
}
?>