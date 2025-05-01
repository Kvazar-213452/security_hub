package encryption

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	config_main "head/main_com/config"
	"io"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/crypto/chacha20poly1305"
)

func EncryptFileChaCha20(filename string, key []byte) ([]byte, error) {
	plaintext, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, chacha20poly1305.NonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aead.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil
}

func DecryptFileChaCha20(filePath string, keyHex string) error {
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return fmt.Errorf("неправильний HEX ключ: %v", err)
	}

	ciphertext, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("помилка читання файлу: %v", err)
	}

	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return fmt.Errorf("помилка створення ChaCha20: %v", err)
	}

	if len(ciphertext) < chacha20poly1305.NonceSize {
		return fmt.Errorf("занадто короткий cipher текст")
	}

	nonce, ciphertext := ciphertext[:chacha20poly1305.NonceSize], ciphertext[chacha20poly1305.NonceSize:]

	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("помилка дешифрування: %v", err)
	}

	outputFilePath := config_main.Frontend_folder + "/static/data/" + filepath.Base(filePath[:len(filePath)-4])

	err = ioutil.WriteFile(outputFilePath, plaintext, 0644)
	if err != nil {
		return fmt.Errorf("помилка запису файлу: %v", err)
	}

	return nil
}
