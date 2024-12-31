package gomail

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

type GoMail struct {
	fromAddress string
	toAddress   string
	smtpHost    string
	smtpPort    string
	password    string
}

func NewGoMail() (*GoMail, error) {

	fromAddress := os.Getenv("FROMADDRESS")
	toAddress := os.Getenv("TOADDRESS")
	smtpHost := os.Getenv("SMTPHOST")
	smtpPort := os.Getenv("SMTPPORT")
	password := os.Getenv("PASSWORD")

	if fromAddress == "" || toAddress == "" || smtpHost == "" || smtpPort == "" || password == "" {
		return nil, fmt.Errorf("invalid environment variables. ")
	}

	goMail := &GoMail{
		fromAddress: fromAddress,
		toAddress:   toAddress,
		smtpHost:    smtpHost,
		smtpPort:    smtpPort,
		password:    password,
	}

	return goMail, nil
}

func (m *GoMail) SendNotification(message string) error {

	if strings.Contains(m.toAddress, "vtext.com") {
		return m.sendText(message)
	} else {
		return m.sendEmail(message)
	}
}

func (m *GoMail) sendText(message string) error {
	//TODO: add back subject if necessary
	//also could send multiple messages if longer than
	body := message
	if len(message) > 159 {
		body = message[0:159]
	}

	auth := smtp.PlainAuth("", m.fromAddress, m.password, m.smtpHost)

	err := smtp.SendMail(m.smtpHost+":"+m.smtpPort, auth, m.fromAddress, []string{m.toAddress}, []byte(body))
	if err != nil {
		fmt.Println("Error sending email:", err)
		return err
	}

	fmt.Println("text sent successfully!")
	return nil
}

func (m *GoMail) sendEmail(message string) error {
	subject := "Subject: Hello from Go!\n"
	body := []byte(subject + "\n" + message)

	// Authenticate with Gmail SMTP server
	auth := smtp.PlainAuth("", m.fromAddress, m.password, m.smtpHost)

	// Send the email
	err := smtp.SendMail(m.smtpHost+":"+m.smtpPort, auth, m.fromAddress, []string{m.toAddress}, body)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}
