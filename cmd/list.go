package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/anxious-aurelius/tmail/internal/config"
	"github.com/anxious-aurelius/tmail/internal/imap"
	"github.com/spf13/cobra"
)

var n int

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "List recent messages from your inbox",
	Long: `Fetch and print the most recent messages in your INBOX over IMAP.

Each line shows the date, sender, and subject. By default the last 10
messages are shown; use --n to change the count.`,
	Example: `  tmail list
  	tmail list --n 25`,

	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			log.Fatal(err)
		}
		envelopes, err := imap.ListEnvelopes(n, cfg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		for _, envelope := range envelopes {
			addressString := ""
			for _, address := range envelope.From {
				addressString += address.PersonalName + " <" + address.MailboxName + "@" + address.HostName + ">,"
			}
			if len(addressString) == 0 {
				addressString = "unknown,"
			}
			addressString = addressString[:len(addressString)-1]
			fmt.Printf("%v : %v : %v\n", envelope.Date.Format(time.DateTime), addressString, envelope.Subject)
		}
		fmt.Println()
	},
}

func init() {
	listCmd.Flags().IntVar(&n, "n", 10, "list length")
	rootCmd.AddCommand(listCmd)
}
