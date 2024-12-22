using System;
using System.Diagnostics;
using System.IO;
using System.Xml;
using System.Text.RegularExpressions;

class Program {
    static void Main() {
        Process process = new Process();
        process.StartInfo.FileName = "netsh";
        process.StartInfo.Arguments = "wlan show interfaces";
        process.StartInfo.RedirectStandardOutput = true;
        process.StartInfo.UseShellExecute = false;
        process.StartInfo.CreateNoWindow = true;

        process.Start();

        string output = process.StandardOutput.ReadToEnd();
        process.WaitForExit();

        string filePath = "main.xml";
        XmlDocument xmlDoc = new XmlDocument();

        XmlDeclaration xmlDeclaration = xmlDoc.CreateXmlDeclaration("1.0", "UTF-8", null);
        xmlDoc.AppendChild(xmlDeclaration);

        XmlElement root = xmlDoc.CreateElement("WiFiData");
        xmlDoc.AppendChild(root);

        AddElement(root, "Name", GetWiFiDetail(output, "Name"));
        AddElement(root, "Description", GetWiFiDetail(output, "Description"));
        AddElement(root, "GUID", GetWiFiDetail(output, "GUID"));
        AddElement(root, "State", GetWiFiDetail(output, "State"));
        AddElement(root, "SignalStrength", GetWiFiDetail(output, "Signal"));
        AddElement(root, "RadioType", GetWiFiDetail(output, "Radio type"));
        AddElement(root, "BSSID", GetWiFiDetail(output, "BSSID"));
        AddElement(root, "Frequency", GetWiFiDetail(output, "Frequency"));
        AddElement(root, "Channel", GetWiFiDetail(output, "Channel"));
        AddElement(root, "SSID", GetWiFiDetail(output, "SSID"));
        AddElement(root, "Authentication", GetWiFiDetail(output, "Authentication"));
        AddElement(root, "Cipher", GetWiFiDetail(output, "Cipher"));
        AddElement(root, "ConnectionMode", GetWiFiDetail(output, "Connection mode"));
        AddElement(root, "ProfileType", GetWiFiDetail(output, "Profile type"));

        xmlDoc.Save(filePath);

        Console.WriteLine($"Інформація записана в файл {filePath}");
    }

    static string GetWiFiDetail(string output, string parameter) {
        string pattern = $@"{parameter}\s*:\s*(.*)";
        Match match = Regex.Match(output, pattern);

        return match.Success ? match.Groups[1].Value.Trim() : "Не знайдено";
    }

    static void AddElement(XmlElement parent, string name, string value) {
        XmlElement element = parent.OwnerDocument.CreateElement(name);
        element.InnerText = value;
        parent.AppendChild(element);
    }
}
