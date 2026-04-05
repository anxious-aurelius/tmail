/*
Copyright © 2026 Kripal Parsekar kripalparsekar@gmail.com
*/
package cmd

import (
	"log"

	"github.com/anxious-aurelius/tmail/config"
	"github.com/anxious-aurelius/tmail/internal/imap"
	"github.com/spf13/cobra"
)

var n int

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatal(err)
		}
		_, err = imap.ListEvelopes(n, cfg)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	listCmd.Flags().IntVar(&n, "n", 10, "list length")
	rootCmd.AddCommand(listCmd)
}
