package func_all

import (
	"bytes"
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// app_back_end/main_com/func_all/func_server.go

type LogContent struct {
	Log string `json:"log"`
}

func LoadLogFile() []byte {
	content, _ := ioutil.ReadFile(config_main.Log_file)

	logContent := LogContent{
		Log: string(content),
	}

	jsonData, _ := json.Marshal(logContent)

	return jsonData
}

type Config struct {
	Port int `json:"port"`
}

func PrintPortFromConfig() int {
	file, err := os.Open("../data/config.json")
	if err != nil {
		fmt.Println("Помилка відкриття файлу:", err)
		return 0
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Println("Помилка декодування JSON:", err)
		return 0
	}

	return config.Port
}

func RemoveNewlines(s string) string {
	return strings.ReplaceAll(s, "\n", "")
}

func splitLines(content string) []string {
	return strings.Split(strings.TrimSpace(content), "\n")
}

func Check_server_Status(url string) int {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	if url == "" {
		return 0
	}

	resp, err := client.Post(url, "application/json", nil)
	if err != nil {
		fmt.Println("error mikroserver " + url)
		return 0
	}

	if resp != nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			return 1
		}
	}

	return 0
}

func Get_server_version() int {
	req, err := http.NewRequest("POST", config_main.Server_register_and_data_url+config_main.Server_register_and_data_url_version, bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Println("ERROR: Failed to create request:", err)
		return 0
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("ERROR: Failed to send request: Post_version_get_server")
		return 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR: Failed to read response body:")
		return 0
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("ERROR: Failed to parse JSON:")
		return 0
	}

	versionStr, ok := response["version"].(string)
	if !ok {
		log.Println("ERROR: Invalid version format in response:", response)
		return 0
	}

	version, err := strconv.Atoi(versionStr)
	if err != nil {
		log.Println("ERROR: Failed to convert version to int:")
		return 0
	}

	return version
}
