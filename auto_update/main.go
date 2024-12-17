package main

import (
	"fmt"
	"head/main_com"
)

func main() {
	version := main_com.Get_version()
	version_conf := main_com.File_config_get_version()

	if version == version_conf {
		fmt.Println(version, version_conf)
	}
}
