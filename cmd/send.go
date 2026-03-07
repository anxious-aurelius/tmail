/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/anxious-aurelius/tmail/config"
	"github.com/anxious-aurelius/tmail/internal/mail"
	"github.com/anxious-aurelius/tmail/internal/smtp"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("send command - implementation inprogess")
		fetchedConfig, err := config.LoadConfig()
		if err != nil {
			fmt.Print(err)
			return
		}
		to := []string{"krupalparsekar3@gmail.com"}
		err = smtp.Send(fetchedConfig, mail.Message{
			To:      to,
			Subject: "test 123",
			Body:    "this is a test message 1234",
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
