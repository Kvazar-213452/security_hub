package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	config_main "head/main_com/config"
	"io"
	"io/ioutil"
	"path/filepath"
)

// module/encryption/main_com/encryption/encryption_aec-CBC.go

func EncryptFileAES_CBC(filename string, key []byte) ([]byte, error) {
	plaintext, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	plaintext = pad(plaintext, aes.BlockSize)

	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return append(iv, ciphertext...), nil
}

func DecryptFileAES_CBC(filePath string, keyHex string) error {
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return fmt.Errorf("error key: %v", err)
	}

	ciphertext, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error file open: %v", err)
	}

	if len(ciphertext) < aes.BlockSize {
		return fmt.Errorf("error IV")
	}

	iv, ciphertext := ciphertext[:aes.BlockSize], ciphertext[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext = unpad(plaintext)

	outputFilePath := config_main.Frontend_folder + "/static/data/" + filepath.Base(filePath[:len(filePath)-4])
	err = ioutil.WriteFile(outputFilePath, plaintext, 0644)
	if err != nil {
		return fmt.Errorf("error in file: %v", err)
	}

	return nil
}

func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := make([]byte, padding)
	for i := 0; i < padding; i++ {
		padText[i] = byte(padding)
	}
	return append(src, padText...)
}

func unpad(src []byte) []byte {
	padding := int(src[len(src)-1])
	return src[:len(src)-padding]
}
