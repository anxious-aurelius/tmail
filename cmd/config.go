package cmd

import (
	"fmt"
	"os"

	"github.com/anxious-aurelius/tmail/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show the currently loaded configuration",
	Long: `Print the SMTP and IMAP settings tmail has loaded from ~/.tmail/config.toml.

Use this to confirm your config file is found and parsed before running send or list.`,
	Example: `  tmail config`,

	Run: func(cmd *cobra.Command, args []string) {
		fetchedConfig, err := config.Load()
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
		fmt.Println(fetchedConfig)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
