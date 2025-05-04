// core/web/static/js/global.js

function button_hover(name) {
    for (let i = 0; i < mmain_buuton.length; i++) {
        $("#" + mmain_buuton[i]).removeClass("vw92dy9qccde32122021"); 
        $("#" + mmain_buuton[i]).addClass("vw92dy9qccde3212202"); 
    }

    $("#" + name).removeClass("vw92dy9qccde3212202"); 
    $("#" + name).addClass("vw92dy9qccde32122021"); 
}

function message_window(content) {
    const $block = $('<div class="animatedBlock hide"></div>').text(content);
    $('body').append($block);

    setTimeout(() => {
        $block.removeClass('hide').addClass('show');
    }, 0);

    setTimeout(() => {
        $block.removeClass('show').addClass('hide');

        setTimeout(() => {
            $block.remove();
        }, 1000);
    }, 3000);
}

function change_menu_antivirus(id) {
    for (let i = 0; i < frg45th9nd.length; i++) {
        $("#" + frg45th9nd[i]).hide();
    } 

    $('#' + frg45th9nd[id]).show();
}

function change_menu_page(id_, id) {
    for (let i = 0; i < mas_sonar[id_].length; i++) {
        $("#" + mas_sonar[id_][i]).addClass("beds12323r4feddfq1");
    }
    
    $("#" + mas_sonar[id_][id]).removeClass("beds12323r4feddfq1"); 
    $("#" + mas_sonar[id_][id]).addClass("beds12323r4feddfq");

    for (let i = 1; i < mas_sonar[id_].length + 1; i++) {
        $('#section_' + i).hide();
    }

    $('#section_' + (id + 1)).show();
}

function get_data_config() {
    $.ajax({
        url: "/api/get_json_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: "data/config.json"}),
        success: function (response) {
            if (response["val"]['style']) {
                get_style();
            }
        }
    });
}

function get_style() {
    $.ajax({
        url: "/api/get_file",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify({data: "data/main.css"}),
        success: function (response) {
           $('#style_dudqdc').html(response["val"]);
        }
    });
}

class TypingEffect {
    constructor(elementId, revealSpeed = 50, finalDelay = 50) {
        this.element = document.getElementById(elementId);
        this.targetText = this.element.textContent.trim();
        this.element.textContent = "";
        this.revealSpeed = revealSpeed;
        this.finalDelay = finalDelay;
        this.randomChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()";
        this.currentText = Array(this.targetText.length).fill(" ");
        this.revealedIndexes = [];
    }

    getRandomChar() {
        return this.randomChars[Math.floor(Math.random() * this.randomChars.length)];
    }

    updateText() {
        let newText = "";
        for (let i = 0; i < this.targetText.length; i++) {
            if (this.revealedIndexes.includes(i)) {
                newText += this.targetText[i];
            } else {
                newText += this.getRandomChar();
            }
        }
        this.element.textContent = newText;
    }

    revealNextChar() {
        if (this.revealedIndexes.length < this.targetText.length) {
            let index;
            do {
                index = Math.floor(Math.random() * this.targetText.length);
            } while (this.revealedIndexes.includes(index));

            this.revealedIndexes.push(index);
            setTimeout(() => this.revealNextChar(), this.finalDelay);
        }
    }

    startAnimation() {
        setInterval(() => this.updateText(), this.revealSpeed);
        setTimeout(() => this.revealNextChar(), 500);
    }
}

get_data_config();