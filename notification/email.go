package notification

import (
	"log"

	"github.com/Favoree-Team/server-user-api/config"
	"gopkg.in/gomail.v2"
)

type EmailNotification interface {
	SendVerification(to string, link string) error
}

type emailNotification struct {
	Dialer *gomail.Dialer
}

func NewEmailNotification() *emailNotification {
	email, pass := config.GetEmailCredential()

	dialer := gomail.NewDialer("smtp.gmail.com", 587, email, pass)

	return &emailNotification{
		Dialer: dialer,
	}
}

func (n *emailNotification) SendVerification(to string, link string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", "Favoree.id Team <noreply@favoree.id>")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Email register verification")
	m.SetBody("text/html", SendVerificationTemplate(link))

	err := n.Dialer.DialAndSend(m)

	if err != nil {
		log.Printf("ERROR send email verification : %s\n%s\n", to, err.Error())
		return err
	}

	log.Printf("SUCCESS email verification %s sent!\n", to)

	return nil
}
