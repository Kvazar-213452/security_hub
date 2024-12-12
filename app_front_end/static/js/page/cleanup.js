$(".toggle-btn").click(function() {
    let nextDropdown = $(this).next(".dropdown-content");
    
    if (nextDropdown.css("display") === "none") {
        nextDropdown.show();
        $(this).removeClass("wefopq2e2e2ccc1");
        $(this).addClass("wefopq2e2e2ccc"); 
    } else if (nextDropdown.css("display") === "block") {
        nextDropdown.hide();
        $(this).removeClass("wefopq2e2e2ccc");
        $(this).addClass("wefopq2e2e2ccc1"); 
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

    if (currentColor === "rgb(24, 24, 34)") {
        $(this).removeClass("weyiqwd22e3c");
        $(this).addClass("weyiqwd22e3c1"); 
        data_cleaning[data] = 1;
    } else if (currentColor === "rgb(55, 55, 69)") {
        $(this).removeClass("weyiqwd22e3c1");
        $(this).addClass("weyiqwd22e3c"); 
        data_cleaning[data] = 0;
    }
});