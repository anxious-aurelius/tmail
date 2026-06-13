/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/anxious-aurelius/tmail/internal/config"
	"github.com/spf13/cobra"
)

var forceOverwrite = false

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Scaffolds a default configuration file",
	Long:  `Creates a default configuration file inside ~/.tmail/ with commented placeholders.`,
	Example: `  tmail config init
  tmail config init --force`,
	Run: func(cmd *cobra.Command, args []string) {
		// setting target path (~/.tmail/)
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error finding home directory: %v\n", err)
			os.Exit(1)
		}
		targetDir := filepath.Join(homeDir, ".tmail")

		// Execute internal logic
		filePath, err := config.InitializeConfig(targetDir, forceOverwrite)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully initialized configuration at: %s\n\n", filePath)
		fmt.Println("Next Steps:")
		fmt.Println("  1. Open the config.toml file in your favorite text editor.")
		fmt.Println("  2. Fill out your credentials.")
		fmt.Println("  3. Run 'tmail config' to verify your connection settings.")
	},
}

func init() {
	configInitCmd.Flags().BoolVarP(&forceOverwrite, "force", "f", false, "Overwrite existing config.toml if it exists")
	configCmd.AddCommand(configInitCmd)
}
