package mail

import (
	"errors"
	"fmt"
	"net/mail"
	"time"
)

type Message struct {
	// Recipient Emails
	To []string
	// Mail Subject
	Subject string
	// Mail Body
	Body string
}

// Envelop struct taken from emersion/imap-client
type Envelope struct {
	// The message date.
	Date time.Time
	// The message subject.
	Subject string
	// The From header addresses.
	From []*Address
	// The message senders.
	// Mail flags
	Flags []string
}

type Address struct {
	// The personal name.
	PersonalName string
	// The mailbox name.
	MailboxName string
	// The host name.
	HostName string
}

func (msg Message) ValidateMessage() error {

	if len(msg.To) == 0 {
		return errors.New("validate message: there should be atleast one recipient")
	}

	for _, to := range msg.To {
		_, err := mail.ParseAddress(to)
		if err != nil {
			return fmt.Errorf("input %s is not a valid email address : %w", to, err)
		}
	}

	return nil

}
