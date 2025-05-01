package func_all

import (
	config_main "head/main_com/config"
	"os"
)

// module/server/main_com/func_all/func_shell.go

func Clear_file(filePath string) {
	file, _ := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0644)
	defer file.Close()
}

func Config_port(data string) {
	Clear_file(config_main.Starter_file)

	file, _ := os.OpenFile(config_main.Starter_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()

	file.WriteString("http://localhost:" + data + "/")
}
