function button_hover(name) {
    $("#" + name).css("border-bottom", "2px solid #ff8a2a");
    $("#" + name).css("color", "#ff8a2a");
}

function clos(name) {
    $('#' + name).hide(); 
}

function openModal(name) {
    $('#' + name).show(); 
}

function fetchLogs() {
    $.get('/get_logs', function(data) {
        $('.console').html(data.join('<br>'));
    });
}

function change_shell(name, button) {
    button_active(button, visualization_mas);
    const dataToSend = {
        message: name
    };

    $.ajax({
        url: '/visualization',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(dataToSend),
        success: function(response) {
            console.log("Server response: " + response);
        },
        error: function(error) {
            console.log("Error: " + error);
        }
    });
}

function button_active(name, mas) {
    mas.forEach(function(item) {
        $("#" + item).css({
            "border": "none",
            "color": "#ffffffd4"
        });
    });

    $("#" + name).css({
        "border": "2px solid #ff8a2a",
        "color": "#ff8a2a"
    });
}

function get_data_wifi_now() {
    $.ajax({
        url: "/get_wifi_now",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function(response) {
            data_wifi_render_now(response)
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function get_data_wifi_all() {
    $.ajax({
        url: "/get_wifi",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function(response) {
            console.log("Відповідь сервера:", response);
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function data_wifi_render_now(response) {
    console.log(response)
    $('#name_wifi').text(response['ssid']);
    $('#authentication_wifi').text(response['authentication']);
    $('#description_wifi').text(response['description']);
    $('#guid_wifi').text(response['guid']);
    $('#physical_address_wifi').text(response['physical_address']);
    $('#radio_type_wifi').text(response['radio_type']);
    $('#signal_strength_wifi').text(response['signal_strength']);
    $('#state_wifi').text(response['state']);
}








function max_wifi() {
    $("#signal_1").removeClass("curveOne2");
    $("#signal_1").addClass("curveOne1");

    $("#signal_2").removeClass("curveTwo2");
    $("#signal_2").addClass("curveTwo1");

    $("#signal_3").removeClass("curveThree2");
    $("#signal_3").addClass("curveThree1");

    $("#signal_4").removeClass("curveFour2");
    $("#signal_4").addClass("curveFour1");
}


function level_wifi_render(level) {
    if (level === 0) {
        max_wifi();

    } else if (level === 1) {
        max_wifi();
        $("#signal_1").removeClass("curveOne1");
        $("#signal_1").addClass("curveOne2");
    } else if (level === 2) {
        max_wifi();
        $("#signal_1").removeClass("curveOne1");
        $("#signal_1").addClass("curveOne2");
        
        $("#signal_2").removeClass("curveTwo1");
        $("#signal_2").addClass("curveTwo2");
    } else if (level === 3) {
        max_wifi();
        $("#signal_1").removeClass("curveOne1");
        $("#signal_1").addClass("curveOne2");
        
        $("#signal_2").removeClass("curveTwo1");
        $("#signal_2").addClass("curveTwo2");
        
        $("#signal_3").removeClass("curveThree1");
        $("#signal_3").addClass("curveThree2");
    } else if (level === 4) {
        max_wifi();
        $("#signal_1").removeClass("curveOne1");
        $("#signal_1").addClass("curveOne2");
        
        $("#signal_2").removeClass("curveTwo1");
        $("#signal_2").addClass("curveTwo2");
        
        $("#signal_3").removeClass("curveThree1");
        $("#signal_3").addClass("curveThree2");

        $("#signal_").removeClass("curveFour1");
        $("#signal_").addClass("curveFour2");
    }
}

level_wifi_render(2)