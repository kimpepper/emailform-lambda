package types

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

type Message struct {
	Name    string
	ReplyTo string
	Body    string
}

type SesClientInterface interface {
	SendEmail(ctx context.Context, params *ses.SendEmailInput, optFns ...func(*ses.Options)) (*ses.SendEmailOutput, error)
}

type FormParserInterface interface {
	ParseFormData(ctx context.Context, req events.LambdaFunctionURLRequest) (error, events.LambdaFunctionURLResponse, *Message)
}

type SenderInterface interface {
	SendEmail(ctx context.Context, msg *Message) error
}
