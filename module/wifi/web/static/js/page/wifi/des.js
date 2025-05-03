function speed_test() {
    openModal("modal1");
    animation = true;

    $.ajax({
        url: "/get_speed",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(),
        success: function (response) {
            clos("modal1");
            animation = false;
            
            $("#f32p32fvbvve").html(response[0]);
            $("#fpf2op3nvnvnnv").html(response[1]);
        }
    });
}
