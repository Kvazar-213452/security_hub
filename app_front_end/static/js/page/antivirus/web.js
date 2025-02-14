// app_front_end/static/js/page/antivirus/web.js

class ArtivirusWeb {
    constructor(lang) {
        this.lang_global = lang;
    }

    antivirus_web_start() {
        let inputValue = $('#fkwe9203f').val();
        const dataToSend = { url_site: [inputValue] };
    
        $('#dwdefw4f4').html(this.lang_global === "uk" ? 'Перевірка.....' : 'Audit.....');
    
        $.ajax({
            url: '/antivirus_web',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(dataToSend),
            success: (response) => this.antivirus_web_end(response)
        });
    }
    
    antivirus_web_end(response) {
        $('#dqdcew336g').show();
    
        clean_div("dq13892r2323233313");
    
        $('#dwdefw4f4').text(this.lang_global === "uk" ? 'Завершено' : 'Completed');
    
        if (this.lang_global === "uk") {
            $('#we31f3qecsdx13rv1').prepend("SSL сертифікат: ");
            $('#we31f3qecsdx13rv').prepend("Загрози: ");
            $('#we31f3qecsdx13rv2').prepend("DNS зміни: ");
            
            $('#dwdefw4f4ewqe').html(response['ssl'] === 1 ? "працює" : "не працює");
            $('#ewfsdt4w43tgfd321').html(response['url'] === 1 ? "не знайдено" : "знайдені");
        } else {
            $('#we31f3qecsdx13rv1').prepend("SSL Certificate: ");
            $('#we31f3qecsdx13rv').prepend("Threats: ");
            $('#we31f3qecsdx13rv2').prepend("DNS Changes: ");
            
            $('#dwdefw4f4ewqe').html(response['ssl'] === 1 ? "working" : "not working");
            $('#ewfsdt4w43tgfd321').html(response['url'] === 1 ? "Not found" : "Found");
        }
    
        $('#qefwfvd244ttff').html(response['dns']);
    }
}
