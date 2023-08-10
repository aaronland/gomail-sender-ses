package main

import (
	_ "github.com/aaronland/gomail-sender-ses"
)

import (
	"context"
	"flag"
	"log"

	"github.com/aaronland/gomail-sender"
	"github.com/aaronland/gomail/v2"
)

func main() {

	sender_uri := flag.String("sender-uri", "", "A valid aaronland/gomail-sender URI")

	from := flag.String("from", "", "A valid From: address (that has been registered with SES)")
	to := flag.String("to", "", "A valid To: address")
	subject := flag.String("subject", "", "A valid email subject")

	flag.Parse()

	ctx := context.Background()

	mail_sender, err := sender.NewSender(ctx, *sender_uri)

	if err != nil {
		log.Fatalf("Failed to create mail sender, %v", err)
	}

	msg := gomail.NewMessage()

	msg.SetHeader("Subject", *subject)
	msg.SetHeader("From", *from)
	msg.SetHeader("To", *to)

	msg.SetBody("text/plain", "This message left intentionally blank.")

	err = gomail.Send(mail_sender, msg)

	if err != nil {
		log.Fatalf("Failed to send message, %v", err)
	}

}
