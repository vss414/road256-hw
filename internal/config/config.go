package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

type Config struct {
	TelegramApiKey string `json:"telegram_api_key"`
	DbHost         string `json:"db_host"`
	DbPort         int    `json:"db_port"`
	DbUser         string `json:"db_user"`
	DbPassword     string `json:"db_password"`
	DbName         string `json:"db_name"`
}

func GetConfig() (*Config, error) {
	var config *Config

	configFile, err := os.Open("../internal/config/config.json")
	if err != nil {
		return nil, errors.Wrap(err, "read config file")
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "decode config file")
	}

	return config, nil
}
