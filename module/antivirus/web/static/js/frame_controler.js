// app_front_end/static/js/frame_controler.js

$(document).on("keydown", function (e) {
    if (e.key === "]" || e.key === 'ї') {
        window.parent.postMessage("console", "*");
    }
});
