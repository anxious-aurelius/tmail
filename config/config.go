package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	SamlConfig Saml `toml:"saml"`
}
type Saml struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

func LoadConfig() {
	var config Config
	var path = "config.toml"
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Print(config)
}
