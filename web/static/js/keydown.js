$(document).on('keydown', function(event) {
    if (event.key === ']' || event.key === 'ї') {
        $('.console').toggle();
    }
});