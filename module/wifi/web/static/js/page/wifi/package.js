// module/wifi/web/static/js/page/wifi/package.js

class PackageData {
    constructor() {
        this.wifiDataPackageElement = $("#wifi_data_packege");
    }

    getPackageData() {
        $.ajax({
            url: "/get_pacage_info",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(null),
            success: (response) => {
                this.writeDataAdapter(response);
            }
        });
    }

    writeDataAdapter(response) {
        const jsonObject = JSON.parse(response);
        this.wifiDataPackageElement.html(null);

        let data = jsonObject["Interfaces"];

        for (let i = 0; i < data.length; i++) {
            let text = `
                <div class="info_adapter_unix">
                    <p><unix>Name = <span>${data[i]["Name"]}</span></unix></p>
                    <p>Description = <span>${data[i]["Description"]}</span></p>
                    <p>Status = <span>${data[i]["Status"]}</span></p>
                    <p>BytesSent = <span>${data[i]["BytesSent"]}</span></p>
                    <p>BytesReceived = <span>${data[i]["BytesReceived"]}</span></p>
                    <p>PacketsSent = <span>${data[i]["PacketsSent"]}</span></p>
                    <p>PacketsReceived = <span>${data[i]["PacketsReceived"]}</span></p>
                </div>
            `;

            this.wifiDataPackageElement.append(text);
        }
    }
}
