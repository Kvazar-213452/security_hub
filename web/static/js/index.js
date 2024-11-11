function page_iframe(url, btn) {
    $("#iframe").attr("src", url);

    button_hover(btn);
}

$(window).on("message", function(event) {
    const receivedData = event.originalEvent.data;

    if (receivedData === lang_change) {
        change_lang_now(1);
    }
});
