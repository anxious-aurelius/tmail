/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/anxious-aurelius/tmail/internal/config"
	"github.com/anxious-aurelius/tmail/internal/mail"
	"github.com/anxious-aurelius/tmail/internal/smtp"
	"github.com/spf13/cobra"
)

var to []string
var subject string
var body string

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
		fmt.Println("Sending Email")
		fetchedConfig, err := config.Load()
		if err != nil {
			fmt.Print(err)
			return
		}

		if to == nil {
			fmt.Println("Enter the email address. Enter q to quit")
			for true {
				var temp string
				fmt.Scanf("Email Address : %s", temp)
				if temp == "q" {
					break
				}
				to = append(to, temp)
			}
		}

		if subject == "" {
			fmt.Println("Enter the email subject")
			fmt.Scanf("Subject : %s", subject)
		}

		if body == "" {
			fmt.Println("Enter the email body")
			fmt.Scanf("Body : \n %s ", body)
		}

		err = smtp.SendMail(fetchedConfig, mail.Message{
			To:      to,
			Subject: subject,
			Body:    body,
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	sendCmd.Flags().StringSliceVar(&to, "to", nil, "recipient email addresses")
	sendCmd.Flags().StringVar(&subject, "subj", "", "email subject")
	sendCmd.Flags().StringVar(&body, "body", "", "email body")

	rootCmd.AddCommand(sendCmd)
}
