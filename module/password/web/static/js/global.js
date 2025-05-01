// app_front_end/static/js/global.js

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
    for (let i = 0; i < frg45th9nd.length; i++) {
        $("#" + frg45th9nd[i]).hide();
    } 

    $('#' + frg45th9nd[id]).show();
}

function change_menu_page(id_, id) {
    for (let i = 0; i < mas_sonar[id_].length; i++) {
        $("#" + mas_sonar[id_][i]).addClass("beds12323r4feddfq1");
    }
    
    $("#" + mas_sonar[id_][id]).removeClass("beds12323r4feddfq1"); 
    $("#" + mas_sonar[id_][id]).addClass("beds12323r4feddfq");

    for (let i = 1; i < mas_sonar[id_].length + 1; i++) {
        $('#section_' + i).hide();
    }

    $('#section_' + (id + 1)).show();
}

function change_lang_now() {
    $.ajax({
        url: "/get_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: "../../core/data/config.json"}),
        success: function (response) {
            let obj = JSON.parse(response);
            lang_change_page(obj["lang"]);
        }
    });
}

function get_data_config() {
    $.ajax({
        url: "/get_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: "../../core/data/config.json"}),
        success: function (response) {
            let obj = JSON.parse(response);

            if (obj['style'] === true) {
                get_style();
            }
        }
    });
}

function get_style() {
    $.ajax({
        url: "/get_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: "../../core/data/main.css"}),
        success: function (response) {
           $('#style_dudqdc').html(response);
        }
    });
}

function clean_div(name) {
    $('#' + name).find('*').contents().filter(function () {
        return this.nodeType === Node.TEXT_NODE;
    }).each(function () {
        this.nodeValue = '';
    });
}

get_data_config();