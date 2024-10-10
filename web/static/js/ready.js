$(document).ready(function() {
    fetchLogs();

    $.ajax({
        url: "/endpoint",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({"key1": "11", "key2": "s22"}),
        success: function(response) {
            console.log("Відповідь сервера:", response);
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
});