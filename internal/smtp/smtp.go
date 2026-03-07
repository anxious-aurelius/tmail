package smtp

import (
	"fmt"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/anxious-aurelius/tmail/config"
	"github.com/anxious-aurelius/tmail/internal/mail"
)

func Send(cfg *config.Config, msg mail.Message) error {
	message := fmt.Sprintf(
		"From: " + cfg.SmtpConfig.Username + "\r\n" +
			"To: " + strings.Join(msg.To, ",") + "\r\n" +
			"Subject: Updated Headers.\r\n" + "\r\n" + msg.Body,
	)
	auth := smtp.PlainAuth("", cfg.SmtpConfig.Username, cfg.SmtpConfig.Password, cfg.SmtpConfig.Host)
	err := smtp.SendMail(cfg.SmtpConfig.Host+":"+strconv.Itoa(cfg.SmtpConfig.Port), auth, cfg.SmtpConfig.Username, msg.To, []byte(message))
	return err
}
