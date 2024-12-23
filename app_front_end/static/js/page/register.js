let div_1 = `
    <p class="register_top_text main_color" id="segister_Poif83g"></p>
    <br>
    <p id="register_pfj3fv"></p>
    <br>
    <input type="text" class="saert45trgf bottom" id="register_name" placeholder="Name">
    <br>
    <input type="text" class="saert45trgf bottom" id="register_gmail" placeholder="Gimail">
    <br>
    <input type="text" class="saert45trgf bottom" id="register_password" placeholder="Password">
    <br><br>
    <button class="zxbnmewd1" onclick="send_data()" id="register_f2qewds"></button>
    <button class="zxbnmewd1" onclick="login_page()" id="register_rfrbbbb4"></button>
`;

let div_2 = `
    <p class="register_top_text main_color" id="segister_Poif83g"></p>
    <br>
    <p id="register_pwfew4fv"></p>
    <br>
    <input type="text" class="saert45trgf bottom" id="register_code" placeholder="Gimail">
    <br><br>
    <button class="zxbnmewd1" onclick="send_code()" id="register_f2qewds"></button>
    <button class="zxbnmewd1" onclick="login_page()" id="register_rfrbbbb4"></button>
`;

let div_3 = `
    <p class="register_top_text main_color" id="register_o92ufff"></p>
    <br>
    <p id="register_pfj3fv1"></p>
    <br>
    <input type="text" class="saert45trgf bottom" id="register_name" placeholder="Name">
    <br>
    <input type="text" class="saert45trgf bottom" id="register_password" placeholder="Password">
    <br><br>
    <button class="zxbnmewd1" onclick="login()" id="register_f2qewds"></button>
    <button class="zxbnmewd1" onclick="reg_page()" id="register_pfi02ifv"></button>
`;

$('#sect_1').html(div_1);

function login_page() {
    $('#sect_1').html(div_3);
    change_lang_now(0);
}

function reg_page() {
    $('#sect_1').html(div_1);
    change_lang_now(0);
}

function send_data() {
    let data = {
        name: $('#register_name').val(),
        gmail: $('#register_gmail').val(),
        password: $('#register_password').val()
    };
    
    $.ajax({
        url: "/send_email",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function (response) {
            $('#sect_1').html(div_2);
        }
    });
}

function send_code() {
    let data = {
        code: $('#register_code').val(),
    };
    
    $.ajax({
        url: "/code_verefic",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function (response) {
            if (response == 1) {
                window.parent.postMessage("reload", "*");
            } else {
                if (lang_global === "uk") {
                    message_window('Невірний пароль');
                } else if (lang_global === "en") {
                    message_window('Invalid password');
                }
            }
        }
    });
}

function login() {
    let data = {
        name: $('#register_name').val(),
        password: $('#register_password').val()
    };
    
    $.ajax({
        url: "/login_acaunt",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function (response) {
            if (response == 1) {
                window.parent.postMessage("reload", "*");
            } else {
                if (lang_global === "uk") {
                    message_window('Невірний пароль');
                } else if (lang_global === "en") {
                    message_window('Invalid password');
                }
            }
        }
    });
}
