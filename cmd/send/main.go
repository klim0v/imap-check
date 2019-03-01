package main

import (
	"crypto/tls"
	"github.com/go-gomail/gomail"
	"log"
	"os"
)

func main() {
	d := gomail.NewDialer("smtp.yandex.ru", 465, os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("AAA", "test")
	m.SetHeader("From", os.Getenv("USERNAME"))
	m.SetHeader("To", os.Getenv("USERNAME"))
	m.SetAddressHeader("Cc", os.Getenv("USERNAME"), "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	if err := d.DialAndSend(m); err != nil {
		log.Fatalln(err)
	}
}
