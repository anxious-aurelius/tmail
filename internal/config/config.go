package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	SmtpConfig Smtp `toml:"smtp"`
	ImapConfig Imap `toml:"imap"`
}
type Smtp struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}
type Imap struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

// Loads config into memory
func Load() (*Config, error) {
	var config Config
	var ErrConfigNotFound = errors.New("config file not found")

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	configFilePath := filepath.Join(homeDir, ".tmail", "config.toml")

	if _, err := os.Stat(configFilePath); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("loading config file from %s: %w\n", configFilePath, ErrConfigNotFound)
		} else {
			return nil, fmt.Errorf("loading config file from %s: %w\n", configFilePath, err)
		}
	}

	_, err = toml.DecodeFile(configFilePath, &config)

	if err != nil {
		return nil, fmt.Errorf("loading config file from %s: %w\n", configFilePath, ErrConfigNotFound)
	}
	return &config, nil
}
