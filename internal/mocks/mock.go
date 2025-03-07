// Package mocks provides mock implementations of interfaces.
package mocks

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/stretchr/testify/mock"

	"emailform/internal/types"
)

// MockSesClient is a mock of the SesClientInterface
type MockSesClient struct {
	types.SesClientInterface
	mock.Mock
}

// SendEmail mocks the SendEmail method
func (m *MockSesClient) SendEmail(ctx context.Context, input *ses.SendEmailInput, _ ...func(*ses.Options)) (*ses.SendEmailOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*ses.SendEmailOutput), args.Error(1)
}

// MockFormParser is a mock of the FormParserInterface
type MockFormParser struct {
	mock.Mock
	types.FormParserInterface
}

// ParseFormData mocks the ParseFormData method
func (m *MockFormParser) ParseFormData(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, *types.Message, error) {
	args := m.Called(ctx, req)
	return args.Get(1).(events.LambdaFunctionURLResponse), args.Get(2).(*types.Message), args.Error(0)
}

// MockSender is a mock of the SenderInterface
type MockSender struct {
	types.SenderInterface
	mock.Mock
}

// SendEmail mocks the SendEmail method
func (m *MockSender) SendEmail(ctx context.Context, msg *types.Message) error {
	args := m.Called(ctx, msg)
	return args.Error(0)
}
