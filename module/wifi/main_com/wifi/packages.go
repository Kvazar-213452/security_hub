package wifi

import (
	"encoding/json"
	"encoding/xml"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

// module/wifi/main_com/wifi/packages.go

func Get_info_packages_wifi() []byte {
	cmd := exec.Command(config_main.Wifi_packege_data_exe)
	cmd.Dir = config_main.Library_folder
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	cmd.Run()

	filePath := "./" + config_main.Library_folder + "/" + config_main.File_data_exe_wifi_packege
	xmlFile, _ := os.Open(filePath)

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var networkInterfaces NetworkInterfaces1

	xml.Unmarshal(byteValue, &networkInterfaces)
	jsonData, _ := json.MarshalIndent(networkInterfaces, "", "  ")

	func_all.Clear_file(config_main.Global_phat + "\\" + config_main.Library_folder + "/" + config_main.File_data_exe_wifi_packege)

	return jsonData
}
