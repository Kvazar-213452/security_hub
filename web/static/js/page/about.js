function get_massage_info() {
    $.ajax({
        url: "http://localhost:3000/data",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            $('#r4e23efvd_').html(response['message']);
            $('#textfrwefwf_').html(response['desc']);
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}
