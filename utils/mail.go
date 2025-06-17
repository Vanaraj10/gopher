package utils

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendVerificationEmail(to string, token string) error {
	from := os.Getenv("MAIL_USER")
    password := os.Getenv("MAIL_PASS")
    smtpHost := os.Getenv("MAIL_HOST")
    smtpPort := os.Getenv("MAIL_PORT")

	verificationLink := fmt.Sprintf("http://localhost:8080/api/verify?token=%s",token)

	subject := "Subject: Verify your email for gopher\n"
	body := fmt.Sprintf("Click the link to verify: %s", verificationLink)
	message := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("",from,password,smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from,[]string{to},message)
	if err != nil {
		log.Println("Failed to send email")
		return err
	}
	log.Println("Verification email sent to",to)
	return nil
}