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
    $.post('/get_logs', function(data) {
        const logsArray = data.logs.split('\n');

        $('.console').html(logsArray.join('<br>'));
    });
}

function console_open() {
    fetchLogs();
    $('.console').toggle();
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
            level_wifi_render(response['signal_strength'])
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
            get_network_now(function(ssid) {
                render_all_network_wifi(response, ssid);
            });
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function data_wifi_render_now(response) {
    $('#info_wifi_div').html(null);
    
    let text = `
    <p class="info_wifi_">SSID: <span id="name_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Authentication: <span id="authentication_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Description: <span id="description_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">GUID: <span id="guid_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Physical Address: <span id="physical_address_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Radio Type: <span id="radio_type_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Signal Strength: <span id="signal_strength_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">State: <span id="state_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Cost: <span id="cost_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Ciphers: <span id="ciphers_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Key Content: <span id="key_content_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Profile Name: <span id="profile_name_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Cost Source: <span id="cost_source_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Approaching Limit: <span id="approaching_limit_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Congested: <span id="congested_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Over Limit: <span id="over_limit_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Roaming: <span id="roaming_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Vendor Extension: <span id="vendor_extension_wifi" class="info_graan"></span></p>
    <p class="info_wifi_">Version: <span id="version_wifi" class="info_graan"></span></p>
    `;

    $('#info_wifi_div').append(text);
    
    $('#name_wifi').text(response['ssid_name'] || 'N/A');
    $('#authentication_wifi').text(response['authentication'] || 'N/A');
    $('#description_wifi').text(response['description'] || 'N/A');
    $('#guid_wifi').text(response['guid'] || 'N/A');
    $('#physical_address_wifi').text(response['physical_address'] || 'N/A');
    $('#radio_type_wifi').text(response['radio_type'] || 'N/A');
    $('#signal_strength_wifi').text(response['signal_strength'] || 'N/A');
    $('#state_wifi').text(response['state'] || 'N/A');
    $('#cost_wifi').text(response['cost'] || 'N/A');
    $('#ciphers_wifi').text(response['ciphers'] ? response['ciphers'].join(', ') : 'N/A');
    $('#key_content_wifi').text(response['key_content'] || 'N/A');
    $('#profile_name_wifi').text(response['profile_name'] || 'N/A');
    $('#cost_source_wifi').text(response['cost_source'] || 'N/A');
    $('#approaching_limit_wifi').text(response['approaching_limit'] || 'N/A');
    $('#congested_wifi').text(response['congested'] || 'N/A');
    $('#over_limit_wifi').text(response['over_limit'] || 'N/A');
    $('#roaming_wifi').text(response['roaming'] || 'N/A');
    $('#vendor_extension_wifi').text(response['vendor_extension'] || 'N/A');
    $('#version_wifi').text(response['version'] || 'N/A');
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
    if (level > 90) {
        max_wifi();

    } else if (level > 75) {
        max_wifi();
        $("#signal_1").removeClass("curveOne1");
        $("#signal_1").addClass("curveOne2");
    } else if (level > 50) {
        max_wifi();
        $("#signal_1").removeClass("curveOne1");
        $("#signal_1").addClass("curveOne2");
        
        $("#signal_2").removeClass("curveTwo1");
        $("#signal_2").addClass("curveTwo2");
    } else if (level > 25) {
        max_wifi();
        $("#signal_1").removeClass("curveOne1");
        $("#signal_1").addClass("curveOne2");
        
        $("#signal_2").removeClass("curveTwo1");
        $("#signal_2").addClass("curveTwo2");
        
        $("#signal_3").removeClass("curveThree1");
        $("#signal_3").addClass("curveThree2");
    } else if (level === 0) {
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

function render_all_network_wifi(response, ssid) {
    for (let i = 0; i < response.length; i++) {
        let text = `<div class="div_wifi_all">
        <p class="name_wifi_div_all">${response[i]['ssid']}</p>
        <p class="right_left_signal">${response[i]['signal']}</p>
        </div>`;
        $('#render_all_wifi').append(text);
    }

    $('.name_wifi_div_all').each(function() {
        if ($(this).text() === ssid) {
            $(this).closest('.div_wifi_all').addClass("border_active_grran");
        }
    });
}

function get_network_now(callback) {
    $.ajax({
        url: "/network_now",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function(response) {
            if (callback) {
                callback(response['ssid']);
            }
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}