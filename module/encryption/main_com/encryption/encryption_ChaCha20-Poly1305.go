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

// module/encryption/main_com/encryption/encryption_ChaCha20-Poly1305.go

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
		return fmt.Errorf("error HEX key: %v", err)
	}

	ciphertext, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error read: %v", err)
	}

	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return fmt.Errorf("error ChaCha20: %v", err)
	}

	if len(ciphertext) < chacha20poly1305.NonceSize {
		return fmt.Errorf("error short cipher")
	}

	nonce, ciphertext := ciphertext[:chacha20poly1305.NonceSize], ciphertext[chacha20poly1305.NonceSize:]

	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	outputFilePath := config_main.Frontend_folder + "/static/data/" + filepath.Base(filePath[:len(filePath)-4])

	err = ioutil.WriteFile(outputFilePath, plaintext, 0644)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}
