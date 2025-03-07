package internal

import (
	"context"
	"net/url"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

// TestParseFormData tests the ParseFormData method of the FormParser struct
func TestParseFormData(t *testing.T) {
	parser := NewFormParser()

	// Create a sample Lambda request with form data
	form := url.Values{}
	form.Add("name", "John Doe")
	form.Add("email", "john.doe@example.com")
	form.Add("content", "Hello, this is a test message.")
	body := form.Encode()

	lambdaReq := events.LambdaFunctionURLRequest{
		RequestContext: events.LambdaFunctionURLRequestContext{
			HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
				Method: "POST",
				Path:   "/",
			},
		},
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		Body: body,
	}

	ctx := context.Background()
	resp, message, err := parser.ParseFormData(ctx, lambdaReq)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.StatusCode != 0 {
		t.Fatalf("expected status code 0, got %d", resp.StatusCode)
	}

	if message.Name != "John Doe" {
		t.Errorf("expected name 'John Doe', got %s", message.Name)
	}

	if message.ReplyTo != "john.doe@example.com" {
		t.Errorf("expected email 'john.doe@example.com', got %s", message.ReplyTo)
	}

	expectedBody := "Hello, this is a test message."
	if message.Body != expectedBody {
		t.Errorf("expected body '%s', got '%s'", expectedBody, message.Body)
	}
}
