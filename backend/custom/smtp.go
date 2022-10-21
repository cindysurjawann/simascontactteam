package custom

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "SIMAS CONTACT"
// const CONFIG_AUTH_EMAIL = os.Getenv("email_sender")
// const CONFIG_AUTH_PASSWORD = os.Getenv("pass_sender")

func SendMail(to []string, cc []string, subject, message string) error {
	// "Cc: " + strings.Join(cc, ",") + "\n" +
	authEmail := os.Getenv("email_sender")
	authPass := os.Getenv("pass_sender")
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", authEmail, authPass, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, authEmail, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
func SendEmail2(to []string, subject, message, typeBody string) error{
	if os.Getenv("testing") != "y"{
		authEmail := os.Getenv("email_sender")
		authPass := os.Getenv("pass_sender")
		fmt.Println("send email", to)
		m := gomail.NewMessage()
		m.SetAddressHeader("From", "simaskominfo@gmail.com", CONFIG_SENDER_NAME)
		m.SetHeader("To", to...)
		m.SetHeader("Subject", subject)
		m.SetBody(typeBody, message)
		dialer := gomail.Dialer{Host: CONFIG_SMTP_HOST, Port: CONFIG_SMTP_PORT, Username: authEmail, Password: authPass}
		// d := gomail.NewDialer(CONFIG_SMTP_HOST, CONFIG_SMTP_PORT, CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD)
		err := dialer.DialAndSend(m)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}