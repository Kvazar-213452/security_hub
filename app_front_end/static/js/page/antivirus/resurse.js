// app_front_end/static/js/page/antivirus/resurse.js

class Resurse {
    get_resurse() {
        $.ajax({
            url: '/antivirus_resurse',
            method: 'POST',
            data: null,
            processData: false,
            contentType: false,
            success: function (response) {
                let arr = JSON.parse(response.status);
                render_resurse(arr);
            }
        });
    }
    
    render_resurse(response) {
        let roundedCpu = response["total_cpu"].toFixed(2);
        let total_memory = response["total_memory"].toFixed(2);
        
        $("#total_cpu").html(roundedCpu);
        $("#total_memory").html(total_memory);
    
        let text_1 = `
            <div class="info_proc_u">
                <p class="iii1">Name</p>
                <p class="iii2">Pid</p>
                <p class="iii3">CPU</p>
                <p class="iii4">Memory</p>
                <br>
            </div>
        `;
    
        $("#proc_div").html(text_1);
    
        for (let i = 0; i < response["processes"].length; i++) {
            let cpu = response["processes"][i]["cpu"].toFixed(2);
            let memory = response["processes"][i]["memory"].toFixed(2);
    
            let text = `
                <div class="info_proc">
                    <p class="iii1">${response["processes"][i]["name"]}</p>
                    <p class="iii2">${response["processes"][i]["pid"]}</p>
                    <p class="iii3">${cpu}</p>
                    <p class="iii4">${memory}</p>
                    <br>
                </div>
            `;
    
            $("#proc_div").append(text);
        }
    }
}
