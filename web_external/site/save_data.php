<?php
if ($_SERVER['REQUEST_METHOD'] == 'POST') {
    $title = $_POST['title'];
    $text = $_POST['text'];
    $filePath = 'db.json';

    // Retrieve existing data and decode it
    $dataArray = file_exists($filePath) ? json_decode(file_get_contents($filePath), true) : [];

    // Create an associative array with input and textarea data
    $newData = [
        'title' => $title,
        'text' => $text
    ];

    // Append new data to the array
    $dataArray[] = $newData;

    // Encode the data back to JSON and save it
    file_put_contents($filePath, json_encode($dataArray, JSON_PRETTY_PRINT));

    echo json_encode(["status" => "success"]);
} else {
    echo json_encode(["status" => "error", "message" => "Invalid request method"]);
}
?>
