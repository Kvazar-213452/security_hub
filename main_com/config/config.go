package config

// data
var Main_config string = "data/main_config.json"
var Log_file string = "data/main.log"

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
var Frontend_folder = "web"

// Server
var Server_url = Get_server_url(Main_config)

var Server_data = Server_url + "site_virus.txt"
var Server_data_sha1_hashes_1 = Server_url + "sha1_hashes_1.txt"
var Server_data_sha1_hashes_2 = Server_url + "sha1_hashes_2.txt"

var Stop_antivirus_flash_drive = make(chan bool)

var Antivirus_flash_drive_cmd string = Get_antivirus_flash_drive_cmd(Main_config)

type Config_global struct {
	Visualization int    `json:"visualization"`
	Log           int    `json:"log"`
	URL           string `json:"url"`
	Port          int    `json:"port"`
	Shell         int    `json:"shell"`
	Lang          string `json:"lang"`
	Antivirus     struct {
		Antivirus_flash_drive     int    `json:"antivirus_flash_drive"`
		Antivirus_flash_drive_cmd string `json:"antivirus_flash_drive_cmd"`
	} `json:"antivirus"`
}
