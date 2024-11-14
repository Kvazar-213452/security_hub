function button_hover(name) {
    for (let i = 0; i < mmain_buuton.length; i++) {
        $("#" + mmain_buuton[i]).removeClass("vw92dy9qccde32122021"); 
        $("#" + mmain_buuton[i]).addClass("vw92dy9qccde3212202"); 
    }

    $("#" + name).removeClass("vw92dy9qccde3212202"); 
    $("#" + name).addClass("vw92dy9qccde32122021"); 
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

function change_menu_antivirus(id) {
    for (let i = 0; i < dwdc21e12d.length; i++) {
        $("#" + dwdc21e12d[i]).addClass("beds12323r4feddfq1");
    }

    for (let i = 0; i < frg45th9nd.length; i++) {
        $("#" + frg45th9nd[i]).hide();
    }

    $("#" + dwdc21e12d[id]).removeClass("beds12323r4feddfq1"); 
    $("#" + dwdc21e12d[id]).addClass("beds12323r4feddfq"); 
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

function get_data_config() {
    $.ajax({
        url: "/config_global",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            if (response['style'] === "main") {
                get_style();
            }
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function get_style() {
    $.ajax({
        url: "/get_style",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
           console.log(response)
           $('#style_dudqdc').html(response);
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

get_data_config();