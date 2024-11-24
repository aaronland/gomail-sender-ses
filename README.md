# gomail-ses

Go package to implement the `gomail.Sender` interface using the AWS Simple Email Service (sesv2).

## Documentation

Documentation is incomplete.

## Example

```
import(
	"context"
	"fmt"
	"net/url"

	"github.com/aaronland/gomail-sender"
	_ "github.com/aaronland/gomail-sender-ses/v2"	
	"github.com/aaronland/gomail/v2"	
)

func main() {

     	config_uri := "aws://?region={REGION}&credentials={CREDENTIALS}"
	sender_uri := fmt.Sprintf("ses://?config-uri=%s", url.QueryEscape)

	from := "bob@bob.com"
	to := "alice@alice.com"

	ctx := context.Background()
	
	s, _ := sender.NewSender(ctx, sender_uri)

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)	
	msg.SetHeader("This is a test")
	msg.SetBody("text/plain", "This is a test")

	gomail.Send(s, msg)
```

## Credentials

Credentials strings (as `?credentials={CREDENTIALS}`) are expected to be a valid [aaronland/go-aws-auth](https://github.com/aaronland/go-aws-auth?tab=readme-ov-file#credentials) credentials "label":

| Label | Description |
| --- | --- |
| `anon:` | Empty or anonymous credentials. |
| `env:` | Read credentials from AWS defined environment variables. |
| `iam:` | Assume AWS IAM credentials are in effect. |
| `sts:{ARN}` | Assume the role defined by `{ARN}` using STS credentials. |
| `{AWS_PROFILE_NAME}` | This this profile from the default AWS credentials location. |
| `{AWS_CREDENTIALS_PATH}:{AWS_PROFILE_NAME}` | This this profile from a user-defined AWS credentials location. |

For example:

```
aws:///us-east-1?credentials=iam:
```

## See also

* https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sesv2
* https://github.com/aaronland/gomail-sender
* https://github.com/aaronland/gomail/v2
* https://github.com/aaronland/go-aws-auth