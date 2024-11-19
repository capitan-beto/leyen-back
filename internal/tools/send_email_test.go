package tools

import (
	"testing"
)

func TestSendEmailBadCredentials(t *testing.T) {

	t.Setenv("FROM_EMAIL", "test@gmail.com")
	t.Setenv("FROM_EMAIL_PASSWORD", "passtest1234")
	t.Setenv("FROM_EMAIL_SMTP", "smtp.gmail.com")
	t.Setenv("SMTP_ADDR", "smtp.gmail.com:587")

	if err := SendEmail([]string{"crnana98@gmail.com"}, "TestMsg", "Hello motafucja"); err == nil {
		t.Fatalf("expected error, got: %v", err)
	}
}
