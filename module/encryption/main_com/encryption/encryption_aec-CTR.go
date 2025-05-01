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

func EncryptFileAES_CTR(filename string, key []byte) ([]byte, error) {
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

	stream := cipher.NewCTR(block, iv)

	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	return append(iv, ciphertext...), nil
}

func DecryptFileAES_CTR(filePath string, keyHex string) error {
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return fmt.Errorf("помилка декодування ключа: %v", err)
	}

	ciphertext, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("помилка відкриття файлу: %v", err)
	}

	if len(ciphertext) < aes.BlockSize {
		return fmt.Errorf("дані занадто короткі для IV")
	}

	iv, ciphertext := ciphertext[:aes.BlockSize], ciphertext[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("помилка створення шифратора: %v", err)
	}

	stream := cipher.NewCTR(block, iv)

	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	outputFilePath := config_main.Frontend_folder + "/static/data/" + filepath.Base(filePath[:len(filePath)-4]) // Заміни на потрібний шлях
	err = ioutil.WriteFile(outputFilePath, plaintext, 0644)
	if err != nil {
		return fmt.Errorf("помилка запису файлу: %v", err)
	}

	return nil
}
