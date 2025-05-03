function config_flash() {
    $.ajax({
        url: "/get_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: "data/bg_script.json"}),
        success: function (response) {
            let obj = JSON.parse(response);

            if (obj["flash_drive"]) {
                $("#bg_dqwderfd").css("background-color", "#68ff9d");
            } else {
                $("#bg_dqwderfd").css("background-color", "#252831");
            }
        }
    });
}

function flash_run() {
    $.ajax({
        url: "/flash_run",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: antivirus_flash_drive, data1: $("#bg_input").val()}),
        success: function (response) {
            if (lang_global === "uk") {
                message_window('Значення встановлено');
            } else if (lang_global === "en") {
                message_window('The value is set');
            }

            config_flash();
        }
    });
}