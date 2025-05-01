// module/password/web/static/js/page/password/user.js

class PasswordManager {
    constructor() {
        this.passwordInput = $('#password_d1qwasz');
        this.passwordConfirmInput = $('#password_d1qwasz1');
        this.keyContainer = $("#pasw_pkjnf2qewvsd");
        this.regButton = $("#pasword_1_btn_page1");
    }

    get_key_reg() {
        $.ajax({
            url: "/get_password",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(null),
            success: (response) => {
                console.log(response);
                const parsedJson = JSON.parse(response);
                this.render_key(parsedJson['key']);
            }
        });
    }

    add_key_pasw() {
        let data = {
            key: this.passwordInput.val(),
            pasw: this.passwordConfirmInput.val()
        };

        $.ajax({
            url: "/add_key_pasw",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(data),
            success: () => {
                this.get_key_reg();
            }
        });
    }

    render_key(data) {
        this.keyContainer.html(null);

        for (let i = 0; i < data.length; i++) {
            let text = `
            <div class="div_pasw_user">
                <p class="table_pasw">${data[i][0]}</p>
                <p class="table_pasw1">${data[i][1]}</p>
                <p onclick="passwordManager.del_key('${data[i][0]}')" class="table_pasw2">del</p>
            </div>
            `;

            this.keyContainer.append(text);
        }
    }

    del_key(data) {
        $.ajax({
            url: "/del_key_pasw",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify({data: data}),
            success: () => {
                this.get_key_reg();
            }
        });
    }

    get_status_reg_hide_pasw() {
        $.ajax({
            url: "/reg_status",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(null),
            success: (response) => {
                const reg_login = response['acsses'];

                if (reg_login == 0) {
                    this.regButton.hide();
                }
            }
        });
    }
}
