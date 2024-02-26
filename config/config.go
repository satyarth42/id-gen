package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	Port   int  `json:"port"`
	DC     int8 `json:"dc"`
	Server int8 `json:"server"`
}

const (
	configFilePath = "CONF_PATH"
	DC_ID          = "DC_ID"
	SERVER_ID      = "SERVER_ID"
)

var conf Config

func getDC() int {
	dc := os.Getenv(DC_ID)
	if dc == "" {
		return 0
	}

	id, err := strconv.Atoi(dc)
	if err != nil {
		return 0
	}
	return id
}

func getServerID() int {
	serverIDEnv := os.Getenv(SERVER_ID)
	if serverIDEnv == "" {
		return 0
	}

	id, err := strconv.Atoi(serverIDEnv)
	if err != nil {
		return 0
	}
	return id
}

func LoadConfig() {

	filePath := os.Getenv(configFilePath)
	if filePath == "" {
		filePath, _ = filepath.Abs("config.json")
	}

	bytes, readErr := os.ReadFile(filePath)
	if readErr != nil {
		log.Fatalf("failed to read file: %s, err:%+v", filePath, readErr)
	}

	config := Config{}
	unmarshalErr := json.Unmarshal(bytes, &config)
	if unmarshalErr != nil {
		log.Fatalf("failed to unmarshal config file bytes: %s, err: %+v", string(bytes), unmarshalErr)
	}

	config.DC = int8(getDC())
	config.Server = int8(getServerID())

	conf = config
}

func GetConfig() Config {
	return conf
}
