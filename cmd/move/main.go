package main

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap-move"
	"github.com/emersion/go-imap/client"
	"log"
	"os"
)

func main() {
	log.Println("Connecting to server...")

	c, err := client.DialTLS("imap.yandex.ru:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")
	defer c.Logout()

	if err := c.Login(os.Getenv("USERNAME"), os.Getenv("PASSWORD")); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	criteria := imap.NewSearchCriteria()
	_, err = c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	criteria.Header.Add("AAA", "test")
	seqNums, err := c.Search(criteria)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(seqNums)

	if len(seqNums) > 0 {
		seqset := new(imap.SeqSet)
		seqset.AddNum(seqNums...)

		err = move.NewClient(c).Move(seqset, "Удаленные")
		if err != nil {
			log.Fatal(err)
		}
	}
}
