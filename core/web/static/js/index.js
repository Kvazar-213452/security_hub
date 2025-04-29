// app_front_end/static/js/index.js

function page_iframe(name, btn) {
    $.ajax({
        url: "/api/get_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: `../module/${name}/starter.md`}),
        success: function (response) {
            $("#iframe").attr("src", response["val"]);

            button_hover(btn);
        }
    });
}

$(document).ready(function() {
    fetchLogs();
});

function get_status_reg() {
    $.ajax({
        url: "/reg_status",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            reg_login = response['acsses']

            if (reg_login == 1) {
                $("#btn10").hide();
            }
        }
    });
}

function get_module() {
    $.ajax({
        url: "/api/get_json_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: "data/config_module.json"}),
        success: function (response) {
            render_module(response);
        }
    });
}

function render_module(response) {
    for (let i = 0; i < response["val"]["module_uinstall"].length; i++) {
        let name = response["val"]["module_uinstall"][i];

        $.ajax({
            url: "/api/get_module_for_render",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify({data: name}),
            success: function (response) {
                let text = `
                    <div onclick="page_iframe('${name}', 'btn${i}')" id="btn${i}" class="button">
                    <img src="data:image/png;base64,${response["icon"]}"><p></p>
                    </div>
                `;

                $(".menu_div").prepend(text);
                db_lang.push([`btn${i}`, response["data"]["lang"]]);
            }
        });
    }

    if (response["val"]["module_uinstall"].length <= 1) {
        // lang
        $.ajax({
            url: "/api/get_json_file",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify({data: "data/config.json"}),
            success: function (response) {
                let type = response["val"]["lang"];

                $("#relon p").html(lang_db[type]["relon"]);

                for (let i = 0; i < db_lang.length; i++) {
                    $(`#${db_lang[i][0]} p`).html(db_lang[i][1][type]);
                }
            }
        });
    }
}

$(window).on("message", function(event) {
    const receivedData = event.originalEvent.data;

    if (receivedData === "lang_change") {
        change_lang_now(1);
    } else if (receivedData === "reload") {
        location.reload();
    }
});

function fetchLogs() {
    $.post('/get_logs', function (data) {
        const logsArray = data.log.split('\n');

        $('.console').html(logsArray.join('<br>'));
    });
}

function console_open() {
    fetchLogs();
    $('.console').toggle();
}

$(document).on('keydown', function(event) {
    if (event.key === ']' || event.key === 'ї') {
        console_open();
    }
});

function render_main_start() {
    $('.ump_textw').html(html_1);

    get_module();
}
