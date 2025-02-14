package page_func

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// app_back_end/main_com/page_func/register.go

type RequestData_xxx struct {
	Receiver string `json:"receiver"`
	Code     string `json:"code"`
}

type Config_reg struct {
	Name   string `json:"name"`
	Pasw   string `json:"pasw"`
	Gmail  string `json:"gmail"`
	Code   string `json:"code"`
	Acsses string `json:"acsses"`
}

type ServerResponse struct {
	Message string `json:"message"`
}

func SendPostRequest_xxx(url string, data RequestData_xxx) {
	jsonData, _ := json.Marshal(data)

	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Email sent successfully!")
	} else {
		fmt.Printf("Failed to send email. Status code: %d\n", resp.StatusCode)
	}
}

func GenerateRandomDigits() string {
	randomBytes := make([]byte, 8)
	rand.Read(randomBytes)

	result := ""
	for _, b := range randomBytes {
		result += fmt.Sprintf("%d", b%10)
	}

	return result
}

func Cripter_xxx(text string) string {
	exePath := "library/aes_encryption.exe"
	args := []string{text}

	cmd := exec.Command(exePath, args...)

	output, _ := cmd.Output()

	outputStr := string(output)

	if strings.HasPrefix(outputStr, "Encrypted:") {
		encryptedText := strings.TrimPrefix(outputStr, "Encrypted: ")
		encryptedText = strings.ReplaceAll(encryptedText, " ", "")
		return encryptedText
	} else {
		return ""
	}
}

func Save_data_reg(config Config_reg) {
	fileData, _ := ioutil.ReadFile(config_main.Data_user)

	var fullConfig Config_reg
	json.Unmarshal(fileData, &fullConfig)

	fullConfig.Name = config.Name
	fullConfig.Pasw = config.Pasw
	fullConfig.Gmail = config.Gmail
	fullConfig.Code = config.Code
	fullConfig.Acsses = config.Acsses

	file, _ := os.Create(config_main.Data_user)
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(fullConfig)
}

func Encrypt_code_reg_save(plainText string) string {
	block, _ := aes.NewCipher([]byte(config_main.Key_post))

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	io.ReadFull(rand.Reader, iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))

	return base64.URLEncoding.EncodeToString(cipherText)
}

func Decrypt_code_reg_save(cipherText string) string {
	cipherBytes, _ := base64.URLEncoding.DecodeString(cipherText)

	block, _ := aes.NewCipher([]byte(config_main.Key_post))

	if len(cipherBytes) < aes.BlockSize {
		return ""
	}

	iv := cipherBytes[:aes.BlockSize]
	cipherBytes = cipherBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(cipherBytes, cipherBytes)

	return string(cipherBytes)
}

func Get_config_user() *Config_reg {
	file, _ := os.Open("../data/user.json")
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var config Config_reg
	json.Unmarshal(byteValue, &config)

	return &config
}

func Send_user_data_server(config Config_reg) string {
	jsonData, _ := json.Marshal(config)

	resp, _ := http.Post(config_main.Server_register_and_data_url+config_main.Server_register_and_data_url_save_user, "application/json", bytes.NewBuffer(jsonData))
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var serverResponse ServerResponse
	json.Unmarshal(body, &serverResponse)

	return serverResponse.Message
}
