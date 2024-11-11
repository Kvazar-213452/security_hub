package page_func

import (
	"bufio"
	config_main "head/main_com/config"
	"net/http"
	"strings"
)

func CheckUrlInFile(url string) int {
	resp, err := http.Get(config_main.Server_data)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == url {
			return 1
		}
	}

	if err := scanner.Err(); err != nil {
		return 0
	}

	return 0
}
