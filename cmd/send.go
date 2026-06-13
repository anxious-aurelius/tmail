package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

		//function call, collects message from stdin
		msg, err := collectMessage(cmd.InOrStdin(), cmd.OutOrStderr(), to, subject, body)

		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "%s", err)
			os.Exit(1)
		}

		//sends mail
		err = smtp.SendMail(fetchedConfig, msg)

		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "%s", err)
			os.Exit(1)
		}

	},
}

func init() {
	sendCmd.Flags().StringSliceVar(&to, "to", nil, "recipient email addresses")
	sendCmd.Flags().StringVar(&subject, "subj", "", "email subject")
	sendCmd.Flags().StringVar(&body, "body", "", "email body")

	rootCmd.AddCommand(sendCmd)
}

// function collects input from stdin. takes input and output stream as param, along with the command flags
func collectMessage(r io.Reader, w io.Writer, to []string, subj string, body string) (mail.Message, error) {

	scanner := bufio.NewScanner(r)

	if to == nil {
		fmt.Fprint(w, "Enter recipient email id's.\n")
		fmt.Fprint(w, "Press return to continue to the next id and 'q/Q' to exit the loop.\n")
		for scanner.Scan() {
			temp := scanner.Text()
			if temp == "q" || temp == "Q" {
				break
			}
			to = append(to, temp)
		}
	}

	if subject == "" {
		fmt.Fprint(w, "Enter the email subject:\n")
		scanner.Scan()
		subj = scanner.Text()
	}

	if body == "" {
		var lines []string
		fmt.Fprint(w, "Enter the email body:\n")
		for scanner.Scan() {
			line := scanner.Text()
			lines = append(lines, line)
		}

		body = strings.Join(lines, "\n")
	}

	if err := scanner.Err(); err != nil {
		return mail.Message{}, fmt.Errorf("collecting send message : %w", err)
	}

	msg := mail.Message{
		To:      to,
		Subject: subj,
		Body:    body,
	}

	return msg, nil
}
