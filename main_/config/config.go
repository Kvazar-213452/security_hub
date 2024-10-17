package config

// data
var Main_config string = "data/main_config.json"
var Log_file string = "data/main.log"

// core web
var Core_web string = "./shell_web.exe"
var Core_web_config string = "start_conf.log"

func Core_web_config_content(port string) string {
	content := `name = Security Hub
window_h = 800
window_w = 1000
html = <style>iframe{position: fixed;height: 100%;width: 100%;top: 0%;left: 0%;}</style><iframe src="http://127.0.0.1` + port + `/about" frameborder="0"></iframe>`
	return content
}
