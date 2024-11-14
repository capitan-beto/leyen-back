package tools

import (
	"net/smtp"
	"os"
)

func SendEmail(to []string, subj, body string) error {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("FROM_EMAIL"),
		os.Getenv("FROM_EMAIL_PASSWORD"),
		os.Getenv("FROM_EMAIL_SMTP"),
	)

	msg := "Subject: " + subj + "\n" + body

	return smtp.SendMail(
		os.Getenv("STMP_ADDR"),
		auth,
		os.Getenv("FROM_EMAIL"),
		to,
		[]byte(msg),
	)
}
