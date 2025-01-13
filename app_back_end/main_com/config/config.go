package config

// ⣿⣿⣿⠟⠛⠛⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⢋⣩⣉⢻⣿⡇
// ⣿⣿⣿⠀⣿⣶⣕⣈⠹⠿⠿⠿⠿⠟⠛⣛⢋⣰⠣⣿⣿⠀⣿⣿
// ⣿⣿⣿⡀⣿⣿⣿⣧⢻⣿⣶⣷⣿⣿⣿⣿⣿⣿⠿⠶⡝⠀⣿⣿
// ⣿⣿⣿⣷⠘⣿⣿⣿⢏⣿⣿⣋⣀⣈⣻⣿⣿⣷⣤⣤⣿⡐⢿⣿
// ⣿⣿⣿⣿⣆⢩⣝⣫⣾⣿⣿⣿⣿⡟⠿⠿⠦⠀⠸⠿⣻⣿⡄⢻
// ⣿⣿⣿⣿⣿⡄⢻⣿⣿⣿⣿⣿⣿⣿⣿⣶⣶⣾⣿⣿⣿⣿⠇⣼
// ⣿⣿⣿⣿⣿⣿⡄⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⣰⣿
// ⣿⣿⣿⣿⣿⣿⠇⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢀⣿⣿
// ⣿⣿⣿⣿⣿⠏⢰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿
// ⣿⣿⣿⣿⠟⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⣿⣿
// ⣿⣿⣿⠋⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⣿⣿
// ⣿⣿⠋⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⢸⣿
// ⣿⠏⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡯⢸⣿

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
var Data_user string = "../data/user.json"
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
var Antivirus_data_exe string = "./data_exe.exe"
var Wifi_packege_data_exe string = "./packages_wifi.exe"

// exe lib data
var File_exe_data string = "get_ssid.xml"
var File_1_exe_data string = "available_wifi.xml"
var File_data_exe string = "data_exe.json"
var File_data_exe_wifi_packege string = "packages_wifi.xml"

// data json
var Antivirus_flash_drive_cmd string = Get_antivirus_flash_drive_cmd(Main_config)

// virustotal api
const ApiKey_virustotal = "b022e26c03533fcd236535e650661b72c41165db486500bf5877f08184b21099"

const ApiURL_virustotal = "https://www.virustotal.com/api/v3/urls/"
const Url_domains_virustotal = "https://www.virustotal.com/api/v3/domains/"
const Files_virustotal = "https://www.virustotal.com/api/v3/files/"

// solver
const Site_main = "https://spx-security-hub.wuaze.com/"
const Site_server = "http://localhost:3000/main"
const Starter_file = "../data/starter.md"

// server
// data
const Server_data_file_url = "http://localhost:3000/"
const Server_data_file_url_search = "search"
const Server_data_file_url_upload = "upload"
const Server_data_file_url_get_how_many = "get_how_many"
const Server_data_file_url_server_unix = "server_unix"

// login and data
const Server_register_and_data_url = "http://127.0.0.1:5000/"
const Server_register_and_data_url_send_email = "send_email"
const Server_register_and_data_url_login = "login"
const Server_register_and_data_url_version = "version"
const Server_register_and_data_url_save_user = "save_user"
const Server_register_and_data_url_get_password = "get_password"
const Server_register_and_data_url_add_key_pasw = "add_key_pasw"
const Server_register_and_data_url_del_key_pasw = "del_key_pasw"
const Server_register_and_data_url_check = "check"
