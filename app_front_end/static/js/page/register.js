let div_1 = `
    <p class="register_top_text main_color">Регестрація</p>
    <br>
    <input type="text" class="saert45trgf bottom" id="register_name" placeholder="Name">
    <br>
    <input type="text" class="saert45trgf bottom" id="register_gmail" placeholder="Gimail">
    <br>
    <input type="text" class="saert45trgf bottom" id="register_password" placeholder="Password">
    <br><br>
    <button class="zxbnmewd1" onclick="send_data()">Send data</button>
    <button class="zxbnmewd1" onclick="add_div()">Залогінитись</button>
`;

let div_2 = `
    <p class="register_top_text main_color">Регестрація</p>
    <br>
    <input type="text" class="saert45trgf bottom" id="register_code" placeholder="Gimail">
    <br><br>
    <button class="zxbnmewd1" onclick="send_code()">Send data</button>
    <button class="zxbnmewd1" onclick="add_div()">Залогінитись</button>
`;

let div_3 = `
    <p class="register_top_text main_color">Регестрація</p>
    <br>
    <input type="text" class="saert45trgf bottom" id="register_name" placeholder="Name">
    <br>
    <input type="text" class="saert45trgf bottom" id="register_password" placeholder="Password">
    <br><br>
    <button class="zxbnmewd1" onclick="login()">Send data</button>
    <button class="zxbnmewd1" onclick="add_div()">Залогінитись</button>
`;

$('#sect_1').html(div_3);

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
            $('.server_0312edcccc').html(div_2);
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
            console.log(response)
            $('.server_0312edcccc').html(div_2);
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
            console.log(response)
        }
    });
}
