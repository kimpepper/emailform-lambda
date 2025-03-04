package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"

	"emailform/internal"
)

var handler *internal.Handler

func init() {
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("unable to load SDK awsconfig, %v", err)
	}

	formParser := internal.NewFormParser()

	sesClient := ses.NewFromConfig(cfg)
	sender := internal.NewSender(sesClient, os.Getenv("FROM_EMAIL"), os.Getenv("TO_EMAIL"))
	handler = internal.NewHandler(formParser, sender)
}

func main() {
	lambda.Start(handler.HandleRequest)
}
