package internal

import (
	"context"
	"fmt"
	"html"
	"net/mail"
	"net/url"

	"github.com/aws/aws-lambda-go/events"

	"emailform/internal/types"
)

type FormParser struct {
}

func NewFormParser() *FormParser {
	return &FormParser{}
}

// ParseFormData parses the form data from the request.
func (p *FormParser) ParseFormData(ctx context.Context, req events.LambdaFunctionURLRequest) (error, events.LambdaFunctionURLResponse, *types.Message) {
	method := req.RequestContext.HTTP.Method
	if method != "POST" {
		return fmt.Errorf("invalid method"), events.LambdaFunctionURLResponse{StatusCode: 405, Body: "Method not allowed"}, &types.Message{}
	}

	// Get form formData
	formData, err := url.ParseQuery(req.Body)
	if err != nil {
		return fmt.Errorf("failed to convert to http request: %w", err), events.LambdaFunctionURLResponse{StatusCode: 400, Body: "Bad request. Unable to parse form."}, &types.Message{}
	}

	// Sanitize name
	name := html.EscapeString(formData.Get("name"))

	// Validate email
	email := formData.Get("email")
	if err := validateEmail(email); err != nil {
		return fmt.Errorf("invalid email: %w", err), events.LambdaFunctionURLResponse{StatusCode: 400, Body: "Bad request. Invalid email."}, &types.Message{}
	}

	// Sanitize content
	content := html.EscapeString(formData.Get("content"))

	return nil, events.LambdaFunctionURLResponse{}, &types.Message{
		Name:    name,
		ReplyTo: email,
		Body:    content,
	}
}

// validateEmail checks if the email address is valid.
func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
