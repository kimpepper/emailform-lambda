package internal

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"

	"emailform/internal/types"
)

type Handler struct {
	formParser types.FormParserInterface
	sender     types.SenderInterface
}

func NewHandler(formParser types.FormParserInterface, sender types.SenderInterface) *Handler {
	return &Handler{
		formParser: formParser,
		sender:     sender,
	}
}

func (h *Handler) HandleRequest(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {

	if req.RequestContext.HTTP.Method != "POST" {
		return events.LambdaFunctionURLResponse{
			StatusCode: 405,
			Body:       "Method not allowed",
		}, fmt.Errorf("invalid method: %s, only POST is allowed", req.RequestContext.HTTP.Method)
	}

	contentType, ok := req.Headers["Content-Type"]
	if !ok || contentType != "application/x-www-form-urlencoded" {
		return events.LambdaFunctionURLResponse{
			StatusCode: 400,
			Body:       "Bad request. Invalid content type.",
		}, fmt.Errorf("invalid content type: %s, only application/x-www-form-urlencoded is allowed", contentType)
	}

	err, response, msg := h.formParser.ParseFormData(ctx, req)
	if err != nil {
		return response, err
	}

	err = h.sender.SendEmail(ctx, msg)
	if err != nil {
		return events.LambdaFunctionURLResponse{StatusCode: 500, Body: "Unable to send email."}, err
	}

	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       "Form submitted",
	}, nil

}
