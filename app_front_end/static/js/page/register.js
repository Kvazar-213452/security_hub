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

$('.server_0312edcccc').html(div_1);

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
            $('.server_0312edcccc').html(div_1);
        }
    });
}

function send_code() {
    let data = {
        code: $('#register_code').val(),
    };
    
    $.ajax({
        url: "/send_email",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function (response) {
            $('.server_0312edcccc').html(div_1);
        }
    });
}
