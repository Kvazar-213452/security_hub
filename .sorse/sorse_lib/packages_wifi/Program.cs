using System;
using System.Linq;
using System.Net.NetworkInformation;
using System.Xml.Linq;

class Program {
    static void Main() {
        var xmlDocument = new XElement("NetworkInterfaces");

        var interfaces = NetworkInterface.GetAllNetworkInterfaces();

        foreach (var netInterface in interfaces) {
            if (netInterface.NetworkInterfaceType == NetworkInterfaceType.Wireless80211) {
                var stats = netInterface.GetIPv4Statistics();

                var interfaceElement = new XElement("Interface",
                    new XElement("Name", netInterface.Name),
                    new XElement("Description", netInterface.Description),
                    new XElement("Status", netInterface.OperationalStatus.ToString()),
                    new XElement("BytesSent", stats.BytesSent / 1024),
                    new XElement("BytesReceived", stats.BytesReceived / 1024),
                    new XElement("PacketsSent", stats.UnicastPacketsSent),
                    new XElement("PacketsReceived", stats.UnicastPacketsReceived)
                );

                xmlDocument.Add(interfaceElement);
            }
        }

        string filePath = "packages_wifi.xml";
        xmlDocument.Save(filePath);

        Console.WriteLine($"Дані записано у файл {filePath}");
    }
}
