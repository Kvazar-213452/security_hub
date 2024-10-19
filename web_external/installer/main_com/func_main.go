package main_com

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func Decode_Base64_ToFile(base64Data string, outputFilePath string) error {
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return fmt.Errorf("не вдалося декодувати base64: %v", err)
	}

	err = ioutil.WriteFile(outputFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("не вдалося записати файл: %v", err)
	}

	return nil
}
