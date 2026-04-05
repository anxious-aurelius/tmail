package imap

import (
	"fmt"
	"log"
	"strconv"

	"github.com/anxious-aurelius/tmail/config"
	"github.com/anxious-aurelius/tmail/internal/mail"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

// TLS Handshake -> Login -> Select Folder -> Fetch Messages -> Logout and close connection
func ListEvelopes(n int, cfg *config.Config) ([]mail.Envelope, error) {

	// TLS handshake with IMAP
	log.Println("Connecting to the server.")
	conn, err := client.DialTLS(cfg.ImapConfig.Host+":"+strconv.Itoa(cfg.ImapConfig.Port), nil)

	if err != nil {
		fmt.Println("Couldn't connect to the IMAP server")
	}
	fmt.Println("Connection succesfully")

	fmt.Println("Logging in...")
	if err = conn.Login(cfg.ImapConfig.Username, cfg.ImapConfig.Password); err != nil {
		fmt.Println("Error when logging in")
	}
	fmt.Println("Successfully logged in")

	fmt.Println("Fetching mails from INBOX")

	//select mailbox
	selectedMB, err := conn.Select("INBOX", true)
	if err != nil {
		fmt.Println("Error selecting mailbox")
	}

	// if mailbox size is greater than the requested list size. just return the last n envelopes
	seqFrom := uint32(1)
	seqTo := selectedMB.Messages

	if selectedMB.Messages > uint32(n) {
		seqFrom = selectedMB.Messages - uint32(n)
	}

	seqset := new(imap.SeqSet)
	seqset.AddRange(seqFrom, seqTo)

	message := make(chan *imap.Message, n)
	done := make(chan error, 1)

	go func() {
		done <- conn.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, message)
	}()

	log.Println("Fetching mails from the mail server.")

	for msg := range message {
		fmt.Println("* " + msg.Envelope.Subject)
	}
	return nil, nil
}
