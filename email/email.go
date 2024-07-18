package email

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

//"02b5f145-6b3e-468c-94d1-3e02bb834f0d"
//"630a75917ce5d93e5f7cb82b6b3e7d8d3e0d64645d49728ae46f308ec36d90b2"
//"s3.ir-thr-at1.arvanstorage.ir"

func SendEmail(subject string, message string) {

	// Sender data.
	from := "mahdi.cpp@gmail.com"
	password := "fleczhadbgbuxpqh"

	// Receiver email address.
	to := []string{
		"mahdi.cpp@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	//message := []byte("Hello Matin. Congratulations! Your withdrawal transaction is successful")

	body := `<b>` + message + `</b>`

	request := Mail{
		Sender:  from,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	msg := BuildMessage(request)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
