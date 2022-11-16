package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func handleRequest(ctx context.Context, event events.SQSEvent) {
	// event
	eventJson, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("EVENT: %s", eventJson)
	// environment variables
	log.Printf("REGION: %s", os.Getenv("AWS_REGION"))
	log.Println("ALL ENV VARS:")
	for _, element := range os.Environ() {
		log.Println(element)
	}
	// request context
	lc, _ := lambdacontext.FromContext(ctx)
	log.Printf("REQUEST ID: %s", lc.AwsRequestID)
	// global variable
	log.Printf("FUNCTION NAME: %s", lambdacontext.FunctionName)
	// context method
	deadline, _ := ctx.Deadline()
	log.Printf("DEADLINE: %s", deadline)
}

func main() {
	runtime.Start(handleRequest)
}
