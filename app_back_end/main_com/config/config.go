package config

// struct
type Config_global struct {
	Version       int    `json:"version"`
	Visualization int    `json:"visualization"`
	Log           int    `json:"log"`
	Port          int    `json:"port"`
	Shell         int    `json:"shell"`
	Lang          string `json:"lang"`
	Style         string `json:"style"`
	Antivirus     struct {
		Antivirus_flash_drive     int    `json:"antivirus_flash_drive"`
		Antivirus_flash_drive_cmd string `json:"antivirus_flash_drive_cmd"`
	} `json:"antivirus"`
}

// data
var Main_config string = "../data/main_config.json"
var Log_file string = "../data/main.log"
var Library_folder string = "library"
var Key_post = "3dp4g9DI8h7MzjVz"

// core web
var Core_web string = "./shell_web.exe"
var Core_web_NM2 string = "./main.exe"

var Window_h string = "800"
var Window_w string = "1000"
var Name string = "Security Hub"

// phat
var Global_phat = Get_phat_global()

// dll
var Cleanup_dll = "library/cleanup.dll"

// Frontend
var Frontend_folder = "../app_front_end"

// flows func
var Stop_antivirus_flash_drive = make(chan bool)

// exe lib
var Available_wifi_exe string = "./available_wifi.exe"
var Get_ssid_exe string = "./get_ssid.exe"
var System_info_exe string = "./system_info.exe"
var Antivirus_data_exe string = "./data_exe.exe"

// exe lib data
var File_exe_data string = "get_ssid.xml"
var File_1_exe_data string = "available_wifi.xml"
var File_2_exe_data string = "system_info.xml"
var File_data_exe string = "data_exe.json"

// data json
var Antivirus_flash_drive_cmd string = Get_antivirus_flash_drive_cmd(Main_config)

// virustotal api
const ApiKey_virustotal = "b022e26c03533fcd236535e650661b72c41165db486500bf5877f08184b21099"

const ApiURL_virustotal = "https://www.virustotal.com/api/v3/urls/"
const Url_domains_virustotal = "https://www.virustotal.com/api/v3/domains/"
const Files_virustotal = "https://www.virustotal.com/api/v3/files/"

// solver
const Site_main = "https://spx-security-hub.wuaze.com/"
const Starter_file = "../data/starter.md"
