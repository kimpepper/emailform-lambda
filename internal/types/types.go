// Package types provides types and interfaces.
package types

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

// Message represents an email message.
type Message struct {
	Name    string
	ReplyTo string
	Body    string
}

// SesClientInterface is an interface for the AWS SES client.
type SesClientInterface interface {
	SendEmail(ctx context.Context, params *ses.SendEmailInput, optFns ...func(*ses.Options)) (*ses.SendEmailOutput, error)
}

// FormParserInterface is an interface for the form parser.
type FormParserInterface interface {
	ParseFormData(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, *Message, error)
}

// SenderInterface is an interface for sending emails.
type SenderInterface interface {
	SendEmail(ctx context.Context, msg *Message) error
}
