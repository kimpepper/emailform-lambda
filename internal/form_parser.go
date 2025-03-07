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

// FormParser is a struct that parses form data from the request.
type FormParser struct {
}

// NewFormParser creates a new FormParser.
func NewFormParser() *FormParser {
	return &FormParser{}
}

// ParseFormData parses the form data from the request.
func (p *FormParser) ParseFormData(_ context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, *types.Message, error) {
	method := req.RequestContext.HTTP.Method
	if method != "POST" {
		return events.LambdaFunctionURLResponse{StatusCode: 405, Body: "Method not allowed"}, &types.Message{}, fmt.Errorf("invalid method")
	}

	// Get form formData
	formData, err := url.ParseQuery(req.Body)
	if err != nil {
		return events.LambdaFunctionURLResponse{StatusCode: 400, Body: "Bad request. Unable to parse form."}, &types.Message{}, fmt.Errorf("failed to convert to http request: %w", err)
	}

	// Sanitize name
	name := html.EscapeString(formData.Get("name"))

	// Validate email
	email := formData.Get("email")
	if err := validateEmail(email); err != nil {
		return events.LambdaFunctionURLResponse{StatusCode: 400, Body: "Bad request. Invalid email."}, &types.Message{}, fmt.Errorf("invalid email: %w", err)
	}

	// Sanitize content
	content := html.EscapeString(formData.Get("content"))

	return events.LambdaFunctionURLResponse{}, &types.Message{
		Name:    name,
		ReplyTo: email,
		Body:    content,
	}, nil
}

// validateEmail checks if the email address is valid.
func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
