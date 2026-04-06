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
func ListEnvelopes(n int, cfg *config.Config) ([]mail.Envelope, error) {

	// TLS handshake with IMAP
	log.Println("Connecting to the server.")
	conn, err := client.DialTLS(cfg.ImapConfig.Host+":"+strconv.Itoa(cfg.ImapConfig.Port), nil)

	if err != nil {
		fmt.Println("Couldn't connect to the IMAP server")
		return nil, err
	}
	fmt.Println("Connection succesfully")

	fmt.Println("Logging in...")
	if err = conn.Login(cfg.ImapConfig.Username, cfg.ImapConfig.Password); err != nil {
		fmt.Println("Error logging in")
		return nil, err
	}
	fmt.Println("Successfully logged in")

	defer conn.Logout()

	fmt.Println("Fetching mails from INBOX")

	//select mailbox
	selectedMB, err := conn.Select("INBOX", true)
	if err != nil {
		fmt.Println("Error selecting mailbox")
		return nil, err
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

	envelopes := []mail.Envelope{}

	for msg := range message {

		tempFrom := []*mail.Address{}

		for _, address := range msg.Envelope.From {
			temp := mail.Address{
				HostName:     address.HostName,
				MailboxName:  address.MailboxName,
				PersonalName: address.PersonalName,
			}
			tempFrom = append(tempFrom, &temp)
		}

		temp := mail.Envelope{
			Date:    msg.Envelope.Date,
			Subject: msg.Envelope.Subject,
			From:    tempFrom,
		}

		envelopes = append(envelopes, temp)

	}
	if err = <-done; err != nil {
		fmt.Println("Some error occured reading mails from the IMAP server.")
		return nil, err
	}

	return envelopes, nil
}
