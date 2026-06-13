package cmd

import (
	"bufio"
	"fmt"
	"strings"

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
	Use:   "send [flags]",
	Short: "Send an email over SMTP",
	Long: `Send an email using the SMTP account in your config file.

Provide the recipients, subject, and body as flags. Credentials and
server settings are read from ~/.tmail/config.toml.`,
	Example: `  tmail send --to alice@example.com --subj "Lunch" --body "12:30 work?"
  tmail send --to a@example.com --to b@example.com --subj "Hi" --body "Hello both"`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sending Email")
		fetchedConfig, err := config.Load()
		if err != nil {
			fmt.Print(err)
			return
		}

		scanner := bufio.NewScanner(cmd.InOrStdin())

		if to == nil {
			fmt.Println("Enter recipient email id's.")
			fmt.Println("Press return to continue to the next id and 'q/Q' to exit the loop.")
			for scanner.Scan() {
				temp := scanner.Text()
				if temp == "q" || temp == "Q" {
					break
				}
				to = append(to, temp)
			}
		}

		if subject == "" {
			fmt.Println("Enter the email subject:")
			scanner.Scan()
			subject = scanner.Text()
		}

		if body == "" {
			var lines []string
			fmt.Println("Enter the email body:")
			for scanner.Scan() {
				line := scanner.Text()
				lines = append(lines, line)
			}

			body = strings.Join(lines, "\n")
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
