$(document).on('keydown', function(event) {
    if (event.key === ']' || event.key === 'ї') {
        console_open();
    }
});

$(document).ready(function() {
    fetchLogs();
});

function button_hover(name) {
    for (let i = 0; i < mmain_buuton.length; i++) {
        $("#" + mmain_buuton[i]).css("opacity", "0.6");
        $("#" + mmain_buuton[i]).css("background-color", "#ffffff00");
    }

    $("#" + name).css("opacity", "1");
    $("#" + name).css("background-color", "#565574");
}

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

function message_window(content) {
    const $block = $('<div class="animatedBlock hide"></div>').text(content);
    $('body').append($block);

    setTimeout(() => {
        $block.removeClass('hide').addClass('show');
    }, 0);

    setTimeout(() => {
        $block.removeClass('show').addClass('hide');

        setTimeout(() => {
            $block.remove();
        }, 1000);
    }, 3000);
}

function open_site() {
    $.ajax({
        url: "/browser_site_app",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            message_window('Сайт відкрито');
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function change_menu_antivirus(id) {
    for (let i = 0; i < dwdc21e12d.length; i++) {
        $("#" + dwdc21e12d[i]).css("color", "#ffffffd4");
    }

    for (let i = 0; i < frg45th9nd.length; i++) {
        $("#" + frg45th9nd[i]).hide();
    }

    $("#" + dwdc21e12d[id]).css("color", "#766aff");
    $('#' + frg45th9nd[id]).show();
}

function change_lang_now(type) {
    $.ajax({
        url: "/config_global",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            lang_global = response['lang'];

            if (type === 0) {
                lang_change_page(lang_global);
            } else {
                lang_change_main(lang_global);
            }
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}
