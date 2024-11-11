$(".toggle-btn").click(function() {
    let nextDropdown = $(this).next(".dropdown-content");
    
    if (nextDropdown.css("display") === "none") {
        nextDropdown.show();
        $(this).css("color", "#766aff");
    } else if (nextDropdown.css("display") === "block") {
        nextDropdown.hide();
        $(this).css("color", "#fff");
    }
});

function cleanup() {
    message_window("Очищення комп'ютера");

    $.ajax({
        url: "/cleanup",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data_cleaning),
        success: function (response) {
            message_window('Завершено');
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

$(".ump_button").click(function() {
    let data = $(this).data("data");
    let currentColor = $(this).css("background-color");

    if (currentColor === "rgb(50, 52, 77)") {
        $(this).css("background-color", "#565574");
        data_cleaning[data] = 1;
    } else if (currentColor === "rgb(86, 85, 116)") {
        $(this).css("background-color", "#32344d");
        data_cleaning[data] = 0;
    }
});