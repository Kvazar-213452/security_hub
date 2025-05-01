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

// module/encryption/main_com/encryption/encryption_AES-GCM.go

func GenerateKey() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return key
}

func EncryptFile(filename string, key []byte) ([]byte, error) {
	plaintext, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil
}

func DecryptFile(filePath string, keyHex string) error {
	key, _ := hex.DecodeString(keyHex)

	ciphertext, _ := ioutil.ReadFile(filePath)

	block, _ := aes.NewCipher(key)

	aesGCM, _ := cipher.NewGCM(block)

	if len(ciphertext) < aesGCM.NonceSize() {
		return fmt.Errorf("error")
	}

	nonce, ciphertext := ciphertext[:aesGCM.NonceSize()], ciphertext[aesGCM.NonceSize():]

	plaintext, _ := aesGCM.Open(nil, nonce, ciphertext, nil)

	outputFilePath := config_main.Frontend_folder + "/static/data/" + filepath.Base(filePath[:len(filePath)-4])

	ioutil.WriteFile(outputFilePath, plaintext, 0644)

	return nil
}
