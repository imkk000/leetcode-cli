package config

import (
	"os"
	"path/filepath"
)

func GetConfigDir() (string, error) {
	return prepareConfigDir()
}

func prepareConfigDir() (string, error) {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".lcc-data", "config")
	_, err := os.ReadDir(configDir)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return "", err
		}
	}
	return configDir, nil
}
