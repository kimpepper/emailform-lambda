package internal

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"emailform/internal/mocks"
	"emailform/internal/types"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
)

func TestSendEmail(t *testing.T) {
	mockSesClient := new(mocks.MockSesClient)
	sender := NewSender(mockSesClient, "foo", "bar")

	msg := &types.Message{
		ReplyTo: "reply@example.com",
		Body:    "Test email body",
	}

	mockSesClient.On("SendEmail", mock.Anything, mock.AnythingOfType("*ses.SendEmailInput")).
		Return(&ses.SendEmailOutput{}, nil)

	err := sender.SendEmail(context.Background(), msg)

	assert.NoError(t, err)
	mockSesClient.AssertExpectations(t)
}

func TestSendEmail_ErrorHandling(t *testing.T) {
	mockSesClient := new(mocks.MockSesClient)
	sender := NewSender(mockSesClient, "foo", "bar")

	msg := &types.Message{
		Name:    "Test Name",
		ReplyTo: "reply@example.com",
		Body:    "Test email body",
	}

	mockSesClient.On("SendEmail", mock.Anything, mock.AnythingOfType("*ses.SendEmailInput")).
		Return(&ses.SendEmailOutput{}, errors.New("failed to send email"))

	err := sender.SendEmail(context.Background(), msg)

	assert.Error(t, err)
	assert.Equal(t, "failed to send email", err.Error())
	mockSesClient.AssertExpectations(t)
}
