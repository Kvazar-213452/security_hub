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

// app_back_end/main_com/config/config.go

// struct
type Config_global struct {
	Version       int    `json:"version"`
	Visualization int    `json:"visualization"`
	Log           int    `json:"log"`
	Port          int    `json:"port"`
	Shell         int    `json:"shell"`
	Lang          string `json:"lang"`
	Style         int    `json:"style"`
}

// data
var Main_config string = "../data/config.json"
var Log_file string = "../data/main.log"
var Data_user string = "../data/user.json"
var Library_folder string = "library"
var Key_post string = "3dp4g9DI8h7MzjVz"

// Frontend
var Frontend_folder string = "web"

// solver
const Site_main string = "https://spx-security-hub.wuaze.com/"
const Site_server string = "http://localhost:3000/main"
const Starter_file string = "starter.md"

//
// data
const Server_data_file_url string = "http://localhost:3000/"
const Server_data_file_url_search string = "search"
const Server_data_file_url_upload string = "upload"
const Server_data_file_url_get_how_many string = "get_how_many"
const Server_data_file_url_server_unix string = "server_unix"

// login and data
const Server_register_and_data_url string = "http://127.0.0.1:5000/"
const Server_register_and_data_url_send_email string = "send_email"
const Server_register_and_data_url_login string = "login"
const Server_register_and_data_url_version string = "version"
const Server_register_and_data_url_save_user string = "save_user"
const Server_register_and_data_url_get_password string = "get_password"
const Server_register_and_data_url_add_key_pasw string = "add_key_pasw"
const Server_register_and_data_url_del_key_pasw string = "del_key_pasw"
const Server_register_and_data_url_check string = "check"

// dep
const Url_dep string = "https://github.com/Kvazar-213452/data/raw/refs/heads/main/dwn_dep.zip"
