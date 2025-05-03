let lang_global;
let animation = true;

function change_lang_now() {
    $.ajax({
        url: "/get_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({"data": "../../../../core/data/config.json"}),
        success: function (response) {
            let obj = JSON.parse(response);
            lang_global = obj["lang"];
            
            lang_change_page(obj["lang"]);
        }
    });
}

function get_data_config() {
    $.ajax({
        url: "/get_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: "../../../../core/data/config.json"}),
        success: function (response) {
            let obj = JSON.parse(response);

            if (obj['style'] === "main") {
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
        data: JSON.stringify({data: "../../../../core/data/main.css"}),
        success: function (response) {
           $('#style_dudqdc').html(response);
        }
    });
}

function scan_flash() {
    $.ajax({
        url: "/scan_flash",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            setTimeout(function () {
                let text = `
                
                `;
    
                $(".main_div").html(text);
                animation = false;
            }, 2000);
        }
    });
}
