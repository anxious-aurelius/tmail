package imap

import (
	"log"
	"strconv"

	"github.com/anxious-aurelius/tmail/config"
	"github.com/anxious-aurelius/tmail/internal/mail"
	"github.com/emersion/go-imap/client"
)

// TLS Handshake -> Login -> Select Folder -> Fetch Messages -> Logout and close connection
func ListEvelopes(n int, cfg *config.Config) ([]mail.Envelope, error) {
	// TLS handshake with IMAP
	log.Println("Connecting to the server.")
	conn, err := client.Dial(cfg.ImapConfig.Host + ":" + strconv.Itoa(cfg.ImapConfig.Port))
	if err != nil {
		return nil, err
	}

	log.Println("Connected to IMAP")

	// Login
	if err = conn.Login(cfg.ImapConfig.Username, cfg.ImapConfig.Password); err != nil {
		return nil, err
	}

	defer conn.Logout()

	mbox, err := conn.Select("INBOX", true)
	if err != nil {
		return nil, err
	}

	log.Println("Flags for INBOX:", mbox.Flags)

	return nil, nil
}
