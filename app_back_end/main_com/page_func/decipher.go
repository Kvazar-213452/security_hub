package page_func

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	config_main "head/main_com/config"
	"io/ioutil"
	"path/filepath"
)

func DecryptFile(filePath string, keyHex string) error {
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return fmt.Errorf("error %v", err)
	}

	ciphertext, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("error %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("error %v", err)
	}

	if len(ciphertext) < aesGCM.NonceSize() {
		return fmt.Errorf("error")
	}

	nonce, ciphertext := ciphertext[:aesGCM.NonceSize()], ciphertext[aesGCM.NonceSize():]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("error %v", err)
	}

	outputFilePath := config_main.Frontend_folder + "/static/data/" + filepath.Base(filePath[:len(filePath)-4])

	ioutil.WriteFile(outputFilePath, plaintext, 0644)

	return nil
}
