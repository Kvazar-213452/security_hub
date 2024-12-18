package page_func

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type RequestData_xxx struct {
	Receiver string `json:"receiver"`
	Code     string `json:"code"`
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

func Cripter_xxx() {
	exePath := "library/aes_encryption.exe"
	args := []string{"kvazar382@gmail.com"}

	cmd := exec.Command(exePath, args...)

	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	outputStr := string(output)

	if strings.HasPrefix(outputStr, "Encrypted:") {
		encryptedText := strings.TrimPrefix(outputStr, "Encrypted: ")
		encryptedText = strings.ReplaceAll(encryptedText, " ", "")
		fmt.Println(encryptedText)
	} else {
		fmt.Println("No encrypted text found in the output.")
	}
}
