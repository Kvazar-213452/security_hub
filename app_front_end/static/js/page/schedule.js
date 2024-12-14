let data = [];
let len = data.length;

function schedule_render(x, y) {
    y = String(y).replace('%', '');
    len += 1;

    data.push({
        x: len,
        y: y,
        z: x
    });

    
    if (data.length > 10) {
        data.shift();
    }

    if (len >= 99) {
        len = 1;
    }

    $chartContainer.innerHTML = '';
    new LineChart(data, $chartContainer).create();
}