package internal

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	sestypes "github.com/aws/aws-sdk-go-v2/service/ses/types"

	"emailform/internal/types"
)

type Sender struct {
	sesClient types.SesClientInterface
	fromEmail string
	toEmail   string
}

func NewSender(sesClient types.SesClientInterface, fromEmail, toEmail string) *Sender {
	return &Sender{
		sesClient: sesClient,
		fromEmail: fromEmail,
		toEmail:   toEmail,
	}
}

// SendEmail sends the form data to the specified email address.
func (s *Sender) SendEmail(ctx context.Context, msg *types.Message) error {
	_, err := s.sesClient.SendEmail(ctx, &ses.SendEmailInput{
		Source:           aws.String(s.fromEmail),
		ReplyToAddresses: []string{msg.ReplyTo},
		Destination: &sestypes.Destination{
			ToAddresses: []string{s.toEmail},
		},

		Message: &sestypes.Message{
			Body: &sestypes.Body{
				Text: &sestypes.Content{
					Data: &msg.Body,
				},
			},
		},
	})

	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}

	return nil
}
