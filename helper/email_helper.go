package helper

import (
	"bytes"
	"html/template"
	"log"

	"github.com/Mr-Malomz/skintech_be/config"
	gomail "gopkg.in/gomail.v2"
)

type infos struct {
	Name string
	OTP  string
}

func SendEmail(firstname string, lastname string, email string) bool {

	t, err := template.ParseFiles("template.html")

	if err != nil {
		log.Panic(err)
	}

	var info infos = infos{Name: firstname + lastname, OTP: GenerateOTP().String()}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, info); err != nil {
		log.Println(err)
	}

	//setting up gomail
	result := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", "no-reply@skintech.io")
	m.SetHeader("To", email)
	m.SetAddressHeader("Cc", "demlabz@gmail.com", "Demola")
	m.SetHeader("Subject", "Skintech - Verify your Account")
	m.SetBody("text/html", result)
	m.Attach("template.html")

	d := gomail.NewDialer("smtp.gmail.com", 587, "no-reply@skintech.io", config.EnvEmailPassword())

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return true

}
