package ses

import (
	"context"
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/aaronland/gomail-sender"
	"github.com/aaronland/gomail/v2"
)

var sender_uri = flag.String("sender-uri", "", "...")
var from = flag.String("from", "", "...")
var to = flag.String("to", "", "...")

func TestSESSender(t *testing.T) {

	if *sender_uri == "" {
		t.Log("-sender-uri is empty, skipping test.")
		t.Skip()
	}

	ctx := context.Background()

	s, err := sender.NewSender(ctx, *sender_uri)

	if err != nil {
		t.Fatalf("Failed to create new sender, %v", err)
	}

	if *from == "" {
		t.Log("-from is empty, skipping test")
		t.Skip()
	}

	if *to == "" {
		t.Log("-to is empty, skipping test")
		t.Skip()
	}

	now := time.Now()

	subject := fmt.Sprintf("This is a test (%v)", now)

	msg := gomail.NewMessage()
	msg.SetHeader("From", *from)
	msg.SetHeader("To", *to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", "This is a test")

	err = gomail.Send(s, msg)

	if err != nil {
		t.Fatalf("Failed to send messsage, %v", err)
	}
}
