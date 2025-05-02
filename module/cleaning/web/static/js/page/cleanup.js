// module/cleaning/web/static/js/page/cleanup.js

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
    openModal("modal1");
    animation = true;

    $.ajax({
        url: "/cleanup",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data_cleaning),
        success: function (response) {
            setTimeout(function () {
                clos("modal1");
                animation = false;
            }, 2000);
        }
    });
}

$(".ump_button").click(function() {
    if (lang_global === "uk") {
        message_window("Оброблено");
    } else if (lang_global === "en") {
        message_window("Processed");           
    }

    let data = $(this).data("data");
    let currentColor = $(this).css("background-color");

    if (currentColor === bg_color1) {
        $(this).removeClass("weyiqwd22e3c");
        $(this).addClass("weyiqwd22e3c1"); 
        data_cleaning[data] = 1;
    } else if (currentColor === bg_color2) {
        $(this).removeClass("weyiqwd22e3c1");
        $(this).addClass("weyiqwd22e3c"); 
        data_cleaning[data] = 0;
    }
});
