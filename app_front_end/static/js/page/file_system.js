let data = {
    labels: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Jul"],
    datasets: [{
        label: "Dataset #1",
        backgroundColor: "rgba(99, 172, 255, 0.2)",
        borderColor: "rgb(115, 172, 236)",
        borderWidth: 2,
        hoverBackgroundColor: "rgba(99, 172, 255, 0.2)",
        hoverBorderColor: "rgb(115, 172, 236)",
        data: [65, 59, 20, 81, 56, 55, 40, 21],
    }]
};

let options = {
    maintainAspectRatio: false,
    scales: {
        y: {
        stacked: true,
        grid: {
            display: true,
            color: "rgba(99, 146, 255, 0.2)"
        }
        },
        x: {
        grid: {
            display: false
        }
        }
    }
};

new Chart('chart', {
    type: 'bar',
    options: options,
    data: data
});


// setings

let mas_file_extension = [];

function render_div() {
    $('#extension_select').html(null);
    
    for (let i = 0; i < mas_file_extension.length; i++) {

        let text = `
        <div class="file_system_div_select">
            <p>${mas_file_extension[i]}</p>
            <div onclick="del_div('${mas_file_extension[i]}')" class="file_system_del_div">del</div>
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

    console.log(mas_file_extension)
    render_div()
}

function inoxwd() {
    let dir_ = "C:\\Users\\god19\\Desktop\\Нова папка (3)";
    let mas1 = mas_file_extension;
    let mas2 = ["py"];

    console.log("Дані, які відправляються на сервер:", { dir: dir_, mas1: mas1, mas2: mas2 });

    $.ajax({
        url: '/scan_dir',
        type: 'POST',
        data: JSON.stringify({ dir: dir_, mas1: mas1, mas2: mas2 }),
        processData: false,
        contentType: 'application/json',
        success: function (response) {
            console.log("Відповідь від сервера:", response);
        },
        error: function (xhr, status, error) {
            console.error("Помилка при запиті:", { status: status, error: error });
        }
    });
}