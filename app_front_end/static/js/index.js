function page_iframe(url, btn) {
    $("#iframe").attr("src", url);

    button_hover(btn);
}

$(window).on("message", function(event) {
    const receivedData = event.originalEvent.data;

    if (receivedData === "lang_change") {
        change_lang_now(1);
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

$(document).ready(function() {
    fetchLogs();
});

function render_main_start() {
    $('.ump_textw').html(html_1);

    page_iframe("/system", "btn1");
    change_lang_now(1);
}
