package main_

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port int `toml:"port"`
}

func LoadConfig(configPath string) Config {
	var config Config

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("none: %v", err)
	}
	defer configFile.Close()

	toml.NewDecoder(configFile).Decode(&config)

	return config
}
