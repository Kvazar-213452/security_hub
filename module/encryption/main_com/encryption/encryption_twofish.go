package encryption

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	config_main "head/main_com/config"
	"io"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/crypto/twofish"
)

// module/encryption/main_com/encryption/encryption_twofish.go

func EncryptFileTwofish(filename string, key []byte) ([]byte, error) {
	plaintext, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, err := twofish.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

func DecryptFileTwofish(filePath string, keyHex string) error {
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return fmt.Errorf("erro key: %v", err)
	}

	ciphertext, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	block, err := twofish.NewCipher(key)
	if err != nil {
		return fmt.Errorf("error %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("error GCM: %v", err)
	}

	if len(ciphertext) < gcm.NonceSize() {
		return fmt.Errorf("error nonce")
	}

	nonce, ciphertext := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	outputFilePath := config_main.Frontend_folder + "/static/data/" + filepath.Base(filePath[:len(filePath)-4])

	err = ioutil.WriteFile(outputFilePath, plaintext, 0644)
	if err != nil {
		return fmt.Errorf("error in file: %v", err)
	}

	return nil
}
