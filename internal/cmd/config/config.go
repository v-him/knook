package config

import (
	"os"
	"path/filepath"
	"encoding/json"
)

type Config struct {
	Token string
}

func Read() (Config, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return Config{}, err
	}

	configDirPath := filepath.Join(userConfigDir, "knook")
	err = os.MkdirAll(configDirPath, 0750)
	if err != nil {
		return Config{}, err
	}

	configFile := filepath.Join(configDirPath, "config.json")
	configContents, err := os.ReadFile(configFile)
	switch err {
	case nil:
	case os.ErrNotExist:
		return Config{}, nil
	default:
		return Config{}, err
	}

	config := &Config{}
	err = json.Unmarshal(configContents, config)
	if err != nil {
		return Config{}, err
	}

	return *config, nil
}
