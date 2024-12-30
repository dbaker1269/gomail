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

func NewGoMail() GoMail {
	var goMail GoMail
	frmAddress := os.Getenv("FROMADDRESS")
	toAddress := os.Getenv("TOADDRESS")
	smtpHost := os.Getenv("SMTPHOST")
	smtpPort := os.Getenv("SMTPPORT")
	password := os.Getenv("PASSWORD")

	if frmAddress != "" {
		goMail.fromAddress = frmAddress
	} else {
		goMail.fromAddress = "dbaker1269@gmail.com"
	}
	if toAddress != "" {
		goMail.toAddress = toAddress
	} else {
		//goMail.toAddress = "dbaker1269@gmail.com"
		goMail.toAddress = "4028025057@vtext.com"
	}
	if smtpHost != "" {
		goMail.smtpHost = smtpHost
	} else {
		goMail.smtpHost = "smtp.gmail.com"
	}
	if smtpPort != "" {
		goMail.smtpPort = smtpPort
	} else {
		goMail.smtpPort = "587"
	}
	if password != "" {
		goMail.password = password
	} else {
		goMail.password = "raetpgjrdpysvlzf"
	}
	return goMail
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
