function button_hover(name) {
    $("#" + name).css("background-color", "#17191f");
    $("#" + name).css("color", "#506aff");
}

function clos(name) {
    $('#' + name).hide(); 
}

function openModal(name) {
    $('#' + name).show(); 
}

function fetchLogs() {
    $.post('/get_logs', function(data) {
        const logsArray = data.log.split('\n');

        $('.console').html(logsArray.join('<br>'));
    });
}

function console_open() {
    fetchLogs();
    $('.console').toggle();
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
            checkUnsafeProtocols();
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

function checkUnsafeProtocols() {
    const text = $("#ciphers_wifi").text();

    if (text) {
        for (const protocol of unsafeProtocols) {
            if (text.includes(protocol)) {
                $("#wifi_protection").html(`<p class="wifi_3_div_red">Незахищено</p>`);
                break;
            }
        }
    }

    $("#wifi_protection").html(`<p class="wifi_3_div">Захищено</p>`);
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

function getConfig() {
    return fetch('/config_global', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(null)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        console.log(data);
        let visualization_button = data['visualization'];
        
        if (visualization_button === 1) {
            button_active('visualization1', visualization_mas);
        } else {
            button_active('visualization2', visualization_mas);
        }
    })
    .catch(error => {
        console.error("Помилка при запиті:", error);
        throw error;
    });
}

function change_shell(name, button) {
    button_active(button, visualization_mas);
    let fff = null
    if (name === true) {
        fff = 1
    } else {
        fff = 0
    }
    const dataToSend = {
        message: fff
    };

    console.log("Data to send:", dataToSend);

    $.ajax({
        url: '/visualization',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(dataToSend),
        success: function(response) {
            console.log("Server response: " + response);
        },
        error: function(xhr, status, error) {
            console.log("Error: " + error);
            console.log("Response text:", xhr.responseText);
        }
    });
}

function dw3fw3() {
    $.ajax({
        url: "/get_os_data",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function(response) {
            dwqdqdxx(response)
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function dw3fw31() {
    $.ajax({
        url: "/usb_info",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function(response) {
            render_usb(response);
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function dwqdqdxx(response) {
    $('#adqwdwxxxx').text(response['computer_name'] || 'N/A');
    $('#qwefsbsfb').text(response['os_version'] || 'N/A');
    $('#efwfwefwegwecc').text(response['processor_info'] || 'N/A');
    $('#fewffefevrewew').text(response['system_memory'] || 'N/A');
    $('#gwghyjyyjy').text(response['system_uptime'] || 'N/A');
    $('#retrevvvvde').text(response['user_name'] || 'N/A');
}

function render_usb(response) {
    response = response['devices']

    for (let i = 0; i < response.length; i++) {
        let text = `<div class="div_wifi_all">
        <p class="name_wifi_div_all">${response[i]}</p>
        </div>`;
        $('#sifewfewx').append(text);
    }
}

function redsgdff() {
    console.log("dddd")
    $.ajax({
        url: "/resource_info",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function(response) {
            dwqdqw2wwww(response);
        },
        error: function(xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

setInterval(redsgdff, 1000);

function dwqdqw2wwww(response) {
    response = response['data'];

    $('#dwqwdwfcfff44').text(response[0] || 'N/A');
    $('#rggwiovnewcee').text(response[1] || 'N/A');
}