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

