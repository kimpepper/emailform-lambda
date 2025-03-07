package internal

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"emailform/internal/mocks"
	"emailform/internal/types"
)

// TestHandleRequest tests the HandleRequest function
func TestHandleRequest(t *testing.T) {
	tests := []struct {
		name           string
		req            events.LambdaFunctionURLRequest
		parseFormError error
		sendEmailError error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Invalid method",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Method: "GET",
					},
				},
			},
			expectedStatus: 405,
			expectedBody:   "Method not allowed",
		},
		{
			name: "Invalid content type",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Method: "POST",
					},
				},
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			expectedStatus: 400,
			expectedBody:   "Bad request. Invalid content type.",
		},
		{
			name: "Form parsing error",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Method: "POST",
					},
				},
				Headers: map[string]string{
					"Content-Type": "application/x-www-form-urlencoded",
				},
			},
			parseFormError: fmt.Errorf("form parsing error"),
			expectedStatus: 400,
			expectedBody:   "Bad request. Invalid form data.",
		},
		{
			name: "Send email error",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Method: "POST",
					},
				},
				Headers: map[string]string{
					"Content-Type": "application/x-www-form-urlencoded",
				},
			},
			sendEmailError: fmt.Errorf("send email error"),
			expectedStatus: 500,
			expectedBody:   "Unable to send email.",
		},
		{
			name: "Successful form submission",
			req: events.LambdaFunctionURLRequest{
				RequestContext: events.LambdaFunctionURLRequestContext{
					HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
						Method: "POST",
					},
				},
				Headers: map[string]string{
					"Content-Type": "application/x-www-form-urlencoded",
				},
			},
			expectedStatus: 200,
			expectedBody:   "Form submitted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formParser := new(mocks.MockFormParser)
			sender := new(mocks.MockSender)

			formParser.On("ParseFormData", mock.Anything, tt.req).Return(tt.parseFormError, events.LambdaFunctionURLResponse{StatusCode: 400, Body: "Bad request. Invalid form data."}, &types.Message{})
			sender.On("SendEmail", mock.Anything, mock.Anything).Return(tt.sendEmailError)

			handler := NewHandler(formParser, sender)
			resp, err := handler.HandleRequest(context.Background(), tt.req)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			assert.Equal(t, tt.expectedBody, resp.Body)
			if tt.expectedStatus != 200 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
