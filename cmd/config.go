/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/anxious-aurelius/tmail/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
