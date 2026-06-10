package config

import (
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

func Load() (*Config, error) {
	var config Config
	//TODO: need to change the path to home directory once packaged into a executable
	var path = "config.toml"
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
