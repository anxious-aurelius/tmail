/*
Copyright © 2026 Kripal Parsekar kripalparsekar@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"time"

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
