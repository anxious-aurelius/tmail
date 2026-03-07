package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	SamlConfig Saml `toml:"saml"`
}
type Saml struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

func LoadConfig() (*Config, error) {
	var config Config
	var path = "config.toml"
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
