// app_front_end/static/js/page/wifi/info.js

function get_data_wifi_now() {
    $.ajax({
        url: "/get_wifi_now",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            data_wifi_render_now(response)
            checkUnsafeProtocols();
        }
    });
}

function get_data_wifi_all() {
    $.ajax({
        url: "/get_wifi",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            get_network_now(function (ssid) {
                render_all_network_wifi(response, ssid);
            });
        }
    });
}

function get_data_for_schedule() {
    $.ajax({
        url: "/network_now",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            let wifi_now = response['ssid'];

            const now = new Date();

            const hours = now.getHours();
            const minutes = now.getMinutes();
            const seconds = now.getSeconds();

            let data = `${hours}:${minutes}:${seconds}`;

            get_wifi_info_level(data, wifi_now)
        }
    });
}

function get_wifi_info_level(data, wifi_now) {
    $.ajax({
        url: "/get_wifi",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            let i = 0

            while (i < response.length) {
                wifi_now = wifi_now.replace(/^"|"$/g, '');
                if (wifi_now === response[i]['SSID']) {
                    schedule_render(data, response[i]['SignalQuality']);
                }
                i++;
            }
        }
    });
}

function data_wifi_render_now(response) {
    $('#fweetrfgcvweee').html(null);

    for (let key in response) {

        if (response.hasOwnProperty(key)) {
            let formattedKey = key.charAt(0).toUpperCase() + key.slice(1);

            if (response[key] == "") {
                let text = `
                <div class="div_info_os">
                    <p class="name_o">${formattedKey}</p>
                    <p class="desc_o"><span>=</span>   N/A</p>
                    <div class="hr_div"></div>
                </div>
                `;

                $('#fweetrfgcvweee').append(text);
            } else {
                let text = `
                <div class="div_info_os">
                    <p class="name_o">${formattedKey}</p>
                    <p class="desc_o"><span>=</span>   ${response[key]}</p>
                    <div class="hr_div"></div>
                </div>
                `;

                $('#fweetrfgcvweee').append(text);
            }
        }
    }
}

function checkUnsafeProtocols() {
    const text = $("#ciphers_wifi").text();

    if (text) {
        for (const protocol of unsafeProtocols) {
            if (text.includes(protocol)) {
                if (lang_global === "uk") {
                    $("#wifi_protection").html(`<p class="wifi_3_div_red">Незахищено</p>`);
                } else if (lang_global === "en") {
                    $("#wifi_protection").html(`<p class="wifi_3_div_red">Not protected</p>`);
                }
                break;
            }
        }
    }

    if (lang_global === "uk") {
        $("#wifi_protection").html(`<p class="wifi_3_div">Захищено</p>`);
    } else if (lang_global === "en") {
        $("#wifi_protection").html(`<p class="wifi_3_div">Protected</p>`);
    }
}

function render_all_network_wifi(response, ssid) {
    $('#render_all_wifi').html(null);

    for (let i = 0; i < response.length; i++) {
        for (let r = i + 1; r < response.length; r++) {
            if (response[i]['SSID'] === response[r]['SSID']) {
                response.splice(r, 1);
            }
        }
    }

    for (let i = 0; i < response.length; i++) {
        let text = `<div class="div_wifi_all">
        <p class="name_wifi_div_all">${response[i]['SSID']}</p>
        <p class="right_left_signal">${response[i]['SignalQuality']}%</p>
        </div>`;
        $('#render_all_wifi').append(text);
    }

    $('.name_wifi_div_all').each(function () {
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
        success: function (response) {
            if (callback) {
                callback(response['ssid']);
            }
        }
    });
}
