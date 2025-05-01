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
	Style         string `json:"style"`
}

// data
var Main_config string = "../data/config.json"
var Log_file string = "../data/main.log"
var Data_user string = "../data/user.json"
var Library_folder string = "library"
var Key_post = "3dp4g9DI8h7MzjVz"

// Frontend
var Frontend_folder = "web"

// solver
const Site_main = "https://spx-security-hub.wuaze.com/"
const Site_server = "http://localhost:3000/main"
const Starter_file = "starter.md"

//
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

// dep
const Url_dep = "http://fi3.bot-hosting.net:23113/app_back_end.zip"
