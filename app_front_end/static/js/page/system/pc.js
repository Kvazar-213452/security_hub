class OSData {
    constructor() {
        // Bind functions to the class instance
        this.get_os_data = this.get_os_data.bind(this);
        this.window_open = this.window_open.bind(this);
        this.write_os_data = this.write_os_data.bind(this);
        this.render_data_window_open = this.render_data_window_open.bind(this);
    }

    get_os_data() {
        $.ajax({
            url: "/get_os_data",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(null),
            success: (response) => {
                this.write_os_data(response);
            }
        });
    }

    window_open() {
        $.ajax({
            url: "/window_open",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(null),
            success: (response) => {
                this.render_data_window_open(response);
            }
        });
    }

    write_os_data(response) {
        let bios_info = response['bios_info'];
        let host_info = response['host_info'];
        let operating_system_info = response['operating_system_info'];

        $('#ddcbnxcew33333').html(null);
        $('#ndwe8rfier').html(null);
        $('#bfgtey65yt').html(null);

        let text = `
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Виробник" : (lang_global === "en" ? "Manu facturer" : "")}</p>
                <p class="desc_o">${bios_info[0]['manufacturer']}</p>
                <div class="hr_div"></div>
            </div>
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Імя" : (lang_global === "en" ? "Тame" : "")}</p>
                <p class="desc_o">${bios_info[0]['name']}</p>
                <div class="hr_div"></div>
            </div>
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Дата випуску" : (lang_global === "en" ? "Release date" : "")}</p>
                <p class="desc_o">${bios_info[0]['release_date']}</p>
                <div class="hr_div"></div>
            </div>
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Серійний номер" : (lang_global === "en" ? "Serial number" : "")}</p>
                <p class="desc_o">${bios_info[0]['serial_number']}</p>
                <div class="hr_div"></div>
            </div>
        `;

        $('#ddcbnxcew33333').append(text);

        let data_text = `
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Імя хоста" : (lang_global === "en" ? "Host name" : "")}</p>
                <p class="desc_o">${host_info['hostname']}</p>
                <div class="hr_div"></div>
            </div>
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Арка ядра" : (lang_global === "en" ? "kernel arch" : "")}</p>
                <p class="desc_o">${host_info['kernel_arch']}</p>
                <div class="hr_div"></div>
            </div>
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Платформа" : (lang_global === "en" ? "Platform" : "")}</p>
                <p class="desc_o">${host_info['platform']}</p>
                <div class="hr_div"></div>
            </div>
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Сімейство платформ" : (lang_global === "en" ? "Platform family" : "")}</p>
                <p class="desc_o">${host_info['platform_family']}</p>
                <div class="hr_div"></div>
            </div>
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Версія платформи" : (lang_global === "en" ? "Platform version" : "")}</p>
                <p class="desc_o">${host_info['platform_version']}</p>
                <div class="hr_div"></div>
            </div>
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Час роботи" : (lang_global === "en" ? "Uptime" : "")}</p>
                <p class="desc_o">${host_info['uptime']} [s]</p>
                <div class="hr_div"></div>
            </div>
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Номер складання" : (lang_global === "en" ? "Build number" : "")}</p>
                <p class="desc_o">${operating_system_info[0]['build_number']}</p>
                <div class="hr_div"></div>
            </div>
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Дата встановлення" : (lang_global === "en" ? "Install date" : "")}</p>
                <p class="desc_o">${operating_system_info[0]['install_date']}</p>
                <div class="hr_div"></div>
            </div>
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Час останнього завантаження" : (lang_global === "en" ? "Last boot up time" : "")}</p>
                <p class="desc_o">${operating_system_info[0]['last_boot_up_time']}</p>
                <div class="hr_div"></div>
            </div>
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Виробник" : (lang_global === "en" ? "Manufacturer" : "")}</p>
                <p class="desc_o">${operating_system_info[0]['manufacturer']}</p>
                <div class="hr_div"></div>
            </div>
            <div class="div_info_os">
                <p class="name_o">${lang_global === "uk" ? "Серійний номер" : (lang_global === "en" ? "Serial number" : "")}</p>
                <p class="desc_o">${operating_system_info[0]['serial_number']}</p>
                <div class="hr_div"></div>
            </div>
        `;

        $('#bfgtey65yt').append(data_text);
    }

    render_data_window_open(response) {
        $('#sifewfewx').html(null);
        response = response['devices'];

        for (let i = 0; i < response.length; i++) {
            if (i + 1 !== response.length) {
                let text = `<div class="div_wifi_all">
                <p class="name_wifi_div_all">${response[i]}</p>
                </div>`;
                $('#sifewfewx').append(text);
            }
        }
    }
}
