package ses

// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sesv2

// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ses-example-send-email.html
// https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-email-addresses-procedure.html

// https://docs.aws.amazon.com/sdk-for-go/api/service/ses/#SES.VerifyEmailIdentity
// https://docs.aws.amazon.com/sdk-for-go/api/service/ses/#SES.WaitUntilIdentityExists
// https://docs.aws.amazon.com/sdk-for-go/api/service/ses/#SES.ListIdentities
// https://docs.aws.amazon.com/sdk-for-go/api/service/ses/#SES.DeleteIdentity

// https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html
// https://docs.aws.amazon.com/ses/latest/APIReference/API_CreateCustomVerificationEmailTemplate.html
// https://docs.aws.amazon.com/sdk-for-go/api/service/ses/#SES.CreateCustomVerificationEmailTemplate

// https://us-west-2.console.aws.amazon.com/ses/home?region=us-west-2#smtp-settings:

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	_ "log"
	"net/url"

	"github.com/aaronland/go-aws-auth"
	"github.com/aaronland/gomail-sender"
	"github.com/aaronland/gomail/v2"
	aws_ses "github.com/aws/aws-sdk-go-v2/service/sesv2"	
)

type SESSender struct {
	gomail.Sender
	client *aws_ses.Client
}

func init() {
	ctx := context.Background()
	err := sender.RegisterSender(ctx, "ses", NewSESSender)

	if err != nil {
		panic(err)
	}
}

func NewSESSender(ctx context.Context, uri string) (gomail.Sender, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	q := u.Query()

	config_uri := q.Get("config-uri")
	
	cfg, err := auth.NewConfig(ctx, config_uri)

	if err != nil {
		return nil, err
	}

	cl := aws_ses.NewFromConfig(cfg)

	s := SESSender{
		client: cl,
	}

	// https://docs.aws.amazon.com/sdk-for-go/api/service/ses/#GetSendQuotaOutput

	return &s, nil
}

func (s *SESSender) Send(from string, to []string, msg io.WriterTo) error {

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	_, err := msg.WriteTo(wr)

	if err != nil {
		return err
	}

	wr.Flush()

	raw_msg := &aws_ses.RawMessage{
		Data: buf.Bytes(),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, recipient := range to {

		err := s.sendMessage(ctx, from, recipient, raw_msg)

		// maybe check err here and sometimes continue ?

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SESSender) sendMessage(ctx context.Context, sender string, recipient string, msg *aws_ses.RawMessage) error {

	// throttle send here... (see quota stuff above)

	select {
	case <-ctx.Done():
		return nil
	default:
		// pass
	}

	req := &aws_ses.SendRawEmailInput{
		RawMessage: msg,
	}

	_, err := s.service.SendRawEmailWithContext(ctx, req)

	if err != nil {
		return err
	}

	return nil
}
