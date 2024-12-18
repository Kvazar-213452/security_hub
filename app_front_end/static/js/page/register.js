function send_data() {
    let data = {
        name: $('.register_name').val(),
        gmail: $('.register_gmail').val(),
        password: $('.register_password').val()
    }
    
    $.ajax({
        url: "/send_email",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function (response) {
            console.log("Ddd")
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}