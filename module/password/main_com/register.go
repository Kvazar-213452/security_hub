package main_com

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// module/password/main_com/register.go

type Config_reg struct {
	Name   string `json:"name"`
	Pasw   string `json:"pasw"`
	Gmail  string `json:"gmail"`
	Code   string `json:"code"`
	Acsses string `json:"acsses"`
}

func Cripter_xxx(plainText string) string {
	key := []byte("3dp4g9DI8h7MzjVz3dp4g9DI8h7MzjVz")
	iv := []byte("1234567890abcdef")

	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}

	paddedText := pkcs7Pad([]byte(plainText), aes.BlockSize)
	ciphertext := make([]byte, len(paddedText))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedText)

	return hex.EncodeToString(ciphertext)
}

func Get_config_user() *Config_reg {
	file, err := os.Open("../../core/data/user.json")
	if err != nil {
		fmt.Println("error read file user.json:", err)
		return nil
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error read file user.json:", err)
		return nil
	}

	var config Config_reg
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		fmt.Println("erorr JSON:", err)
		return nil
	}

	return &config
}

func Decrypter_AES256(encryptedHex string) string {
	key := []byte("3dp4g9DI8h7MzjVz3dp4g9DI8h7MzjVz")
	iv := []byte("1234567890abcdef")

	ciphertext, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return ""
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}

	if len(ciphertext) < aes.BlockSize {
		return ""
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext, err = pkcs7Unpad(plaintext)
	if err != nil {
		return ""
	}

	return string(plaintext)
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("pkcs7: empty data")
	}

	padLength := int(data[len(data)-1])
	if padLength > len(data) {
		return nil, fmt.Errorf("pkcs7: invalid padding")
	}

	for i := len(data) - padLength; i < len(data); i++ {
		if int(data[i]) != padLength {
			return nil, fmt.Errorf("pkcs7: invalid padding")
		}
	}

	return data[:len(data)-padLength], nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}
