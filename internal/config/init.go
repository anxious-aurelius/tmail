package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const DefaultConfigTemplate = `# tmail Configuration File
# For more information, see the documentation.

[smtp]
host     = "smtp.example.com"
port     = 587
username = "you@example.com"
password = "your-password"

[imap]
host     = "imap.example.com"
port     = 993
username = "you@example.com"
password = "your-password"
`

func InitializeConfig(targetDir string, force bool) (string, error) {
	// Used 0700 since Linux and Mac needs executable permisson
	if err := os.MkdirAll(targetDir, 0700); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	filePath := filepath.Join(targetDir, "config.toml")

	if _, err := os.Stat(filePath); err == nil {
		if !force {
			return filePath, errors.New("configuration files already exists (use --force to overwrite)")
		}
	}

	err := os.WriteFile(filePath, []byte(DefaultConfigTemplate), 0600)
	if err != nil {
		return "", fmt.Errorf("failed to write config file: %v", err)
	}

	return filePath, nil
}
