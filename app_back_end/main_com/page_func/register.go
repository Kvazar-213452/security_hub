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
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type RequestData_xxx struct {
	Receiver string `json:"receiver"`
	Code     string `json:"code"`
}

type Config_reg struct {
	Name  string `json:"name"`
	Pasw  string `json:"pasw"`
	Gmail string `json:"gmail"`
	Code  string `json:"code"`
}

func SendPostRequest_xxx(url string, data RequestData_xxx) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error marshalling data:", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error sending POST request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Email sent successfully!")
	} else {
		fmt.Printf("Failed to send email. Status code: %d\n", resp.StatusCode)
	}
}

func GenerateRandomDigits() string {
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Error generating random bytes:", err)
	}

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

	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

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
	fileData, err := ioutil.ReadFile("../data/user.json")
	if err != nil {
		fmt.Println("Помилка при читанні файлу:", err)
		return
	}

	var fullConfig Config_reg
	err = json.Unmarshal(fileData, &fullConfig)
	if err != nil {
		fmt.Println("Помилка при розпарсуванні JSON:", err)
		return
	}

	fullConfig.Name = config.Name
	fullConfig.Pasw = config.Pasw
	fullConfig.Gmail = config.Gmail
	fullConfig.Code = config.Code

	file, err := os.Create("../data/user.json")
	if err != nil {
		fmt.Println("Помилка при створенні файлу:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(fullConfig)
	if err != nil {
		fmt.Println("Помилка при запису в файл:", err)
		return
	}

	fmt.Println("Дані успішно записано в файл main_config.json")
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
