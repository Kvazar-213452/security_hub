package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	config_main "head/main_com/config"
	"io/ioutil"
	"path/filepath"
)

// app_back_end/main_com/page_func/encryption/decipher.go

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
