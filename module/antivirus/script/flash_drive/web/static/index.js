let lang_global;
let animation = true;
let response1;

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
            response1 = response;

            setTimeout(function () {
                let text = `
                    <p class="mmtj576yrt34" id="f23f34trtyhy566hp"></p>
                    <br><br>
                    <p class="f232f9hwovvv4" id="f32pol24vnza"></p>
                    <br>
                    <div class="gh5j6juthgbfe" id="d32dq3fweaev_ew"></div>
                    <br><br>
                    <p class="f232f9hwovvv4" id="fpkpzj021r23ovew"></p>
                    <br>
                    <div class="gh5j6juthgbfe" id="d32r3fgnifwefwe"></div>
                `;
    
                $(".lod3br32fvvvv2_3").html(text);

                change_lang_now();
                render_file();

                animation = false;
            }, 2000);
        }
    });
}

function render_file() {
    $("#d32dq3fweaev_ew").html(null);
    $("#d32r3fgnifwefwe").html(null);

    for (let i = 0; i < response1["all_exe_files"].length; i++) {
        text = `
            <div class="pfk230344fvvnvn2">${response1["all_exe_files"][i]}</div>
        `;

        $("#d32dq3fweaev_ew").append(text);
    }

    for (let i = 0; i < response1["malicious_files"].length; i++) {
        const filePath = response1["malicious_files"][i];
        const safePath = JSON.stringify(filePath)
        
        text = `
            <div class="pfk230344fvvnvn21">${response1["malicious_files"][i]} 
            <button onclick='del_file(${safePath})'>del</button>
            </div>
        `;

        $("#d32r3fgnifwefwe").append(text);
    }
}

function del_file(name) {
    $.ajax({
        url: "/del_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: name}),
        success: function (response) {
            for (let i = 0; i < response1["malicious_files"].length; i++) {
                if (response1["malicious_files"][i] == name) {
                    response1["malicious_files"].splice(i, 1);
                    break;
                }
            }

            render_file();
        }
    });
}
