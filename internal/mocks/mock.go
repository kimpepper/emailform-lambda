package mocks

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/stretchr/testify/mock"

	"emailform/internal/types"
)

type MockSesClient struct {
	types.SesClientInterface
	mock.Mock
}

func (m *MockSesClient) SendEmail(ctx context.Context, input *ses.SendEmailInput, opts ...func(*ses.Options)) (*ses.SendEmailOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*ses.SendEmailOutput), args.Error(1)
}

type MockFormParser struct {
	mock.Mock
	types.FormParserInterface
}

func (m *MockFormParser) ParseFormData(ctx context.Context, req events.LambdaFunctionURLRequest) (error, events.LambdaFunctionURLResponse, *types.Message) {
	args := m.Called(ctx, req)
	return args.Error(0), args.Get(1).(events.LambdaFunctionURLResponse), args.Get(2).(*types.Message)
}

type MockSender struct {
	types.SenderInterface
	mock.Mock
}

func (m *MockSender) SendEmail(ctx context.Context, msg *types.Message) error {
	args := m.Called(ctx, msg)
	return args.Error(0)
}
