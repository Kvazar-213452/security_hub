using System;
using System.Diagnostics;
using System.IO;
using System.Xml;
using System.Text.RegularExpressions;

// .sorse/sorse_lib/get_netsh/Program.cs

class Program {
    static void Main() {
        string ssid = GetActiveSSID();

        if (string.IsNullOrEmpty(ssid)) {
            Console.WriteLine("error");
            return;
        }

        Process process = new Process();
        process.StartInfo.FileName = "netsh";
        process.StartInfo.Arguments = $"wlan show profile name=\"{ssid}\" key=clear";
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

        XmlElement root = xmlDoc.CreateElement("WiFiProfile");
        xmlDoc.AppendChild(root);

        AddElement(root, "SSID", ssid);
        AddElement(root, "Authentication", GetWiFiDetail(output, "Authentication"));
        AddElement(root, "Cipher", GetWiFiDetail(output, "Cipher"));
        AddElement(root, "NetworkType", GetWiFiDetail(output, "Network type"));
        AddElement(root, "RadioType", GetWiFiDetail(output, "Radio type"));
        AddElement(root, "ConnectionMode", GetWiFiDetail(output, "Connection mode"));
        AddElement(root, "MACRandomization", GetWiFiDetail(output, "MAC Randomization"));
        AddElement(root, "NetworkBroadcast", GetWiFiDetail(output, "Network broadcast"));
        AddElement(root, "AutoSwitch", GetWiFiDetail(output, "AutoSwitch"));
        AddElement(root, "Cost", GetWiFiDetail(output, "Cost"));
        AddElement(root, "Congested", GetWiFiDetail(output, "Congested"));
        AddElement(root, "ApproachingDataLimit", GetWiFiDetail(output, "Approaching Data Limit"));
        AddElement(root, "OverDataLimit", GetWiFiDetail(output, "Over Data Limit"));
        AddElement(root, "Roaming", GetWiFiDetail(output, "Roaming"));
        AddElement(root, "CostSource", GetWiFiDetail(output, "Cost Source"));

        xmlDoc.Save(filePath);

        Console.WriteLine($"error {filePath}");
    }

    static string GetActiveSSID() {
        Process process = new Process();
        process.StartInfo.FileName = "netsh";
        process.StartInfo.Arguments = "wlan show interfaces";
        process.StartInfo.RedirectStandardOutput = true;
        process.StartInfo.UseShellExecute = false;
        process.StartInfo.CreateNoWindow = true;

        process.Start();

        string output = process.StandardOutput.ReadToEnd();
        process.WaitForExit();

        return GetWiFiDetail(output, "SSID");
    }

    static string GetWiFiDetail(string output, string parameter) {
        string pattern = $@"{parameter}\s*:\s*(.*)";
        Match match = Regex.Match(output, pattern);

        return match.Success ? match.Groups[1].Value.Trim() : "error";
    }

    static void AddElement(XmlElement parent, string name, string value) {
        XmlElement element = parent.OwnerDocument.CreateElement(name);
        element.InnerText = value;
        parent.AppendChild(element);
    }
}
