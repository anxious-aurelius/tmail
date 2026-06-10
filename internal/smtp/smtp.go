package smtp

import (
	"fmt"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/anxious-aurelius/tmail/internal/config"
	"github.com/anxious-aurelius/tmail/internal/mail"
)

func SendMail(cfg *config.Config, msg mail.Message) error {
	message := fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n\r\n%s",
		cfg.SmtpConfig.Username, strings.Join(msg.To, ", "), msg.Subject, msg.Body)
	auth := smtp.PlainAuth("", cfg.SmtpConfig.Username, cfg.SmtpConfig.Password, cfg.SmtpConfig.Host)
	err := smtp.SendMail(cfg.SmtpConfig.Host+":"+strconv.Itoa(cfg.SmtpConfig.Port), auth, cfg.SmtpConfig.Username, msg.To, []byte(message))
	return err
}
