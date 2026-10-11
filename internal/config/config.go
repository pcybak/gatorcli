package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	configPath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath = filepath.Join(configPath, configFileName)
	return configPath, nil
}

func Read() (Config, error) {
	config := Config{}
	configPath, err := getConfigFilePath()
	if err != nil {
		return config, err
	}
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user

	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, bytes, 0o600)
	if err != nil {
		return err
	}
	return nil
}
