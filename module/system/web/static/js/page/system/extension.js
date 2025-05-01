// module/system/web/static/js/page/system/extension.js

let mas_file_extension = [];

function render_div() {
    $('#extension_select').html(null);

    for (let i = 0; i < mas_file_extension.length; i++) {

        let text = `
        <div class="file_system_div_select">
            <p>${mas_file_extension[i]}</p>
            <div onclick="del_div('${mas_file_extension[i]}', 0)" class="file_system_del_div">del</div>
        </div>
    `;

    $('#extension_select').append(text);
    }
    
}

function add_div() {
    let text_file_extension = $('#file_stystem_dwqdasz').val();

    mas_file_extension.push(text_file_extension);

    render_div()
}

function del_div(text_file_extension) {
    mas_file_extension.splice(text_file_extension, 1);

    render_div()
}

function render_top_10(data) {
    $('#data_bate').html(data.Rootsize);

    for (let i = 0; i < 9; i++) {
        let text = `
        <div class="top_10_extension"><span1>${data.Top[i][0]}</span1><span2>${data.Top[i][1]}</span2><span3>${data.Top[i][2]}</span3></div>
        `;
        
        $('#out_data_extension').append(text);
    }
}

function inoxwd() {
    let dir_ =  $('#file_stystem_dwqdasz_2').val();
    let mas1 = mas_file_extension;

    $.ajax({
        url: '/scan_dir',
        type: 'POST',
        data: JSON.stringify({ dir: dir_, mas1: mas1 }),
        processData: false,
        contentType: 'application/json',
        success: function (response) {
            render_top_10(response);
        }
    });
}
