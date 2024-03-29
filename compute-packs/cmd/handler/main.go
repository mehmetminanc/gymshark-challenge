package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"

	"github.com/mehmetminanc/gymshark-challenge/compute-packs/handlers"
)

func main() {
	//this lambda setup is used for testing/running locally without deployment
	//if specific events are required we can read in json event samples or
	//command line flags to help build the events

	lambdaMaxRuntime := time.Now().Add(15 * time.Minute)

	ctx, cancel := context.WithDeadline(context.Background(), lambdaMaxRuntime)
	defer cancel()

	handler := handlers.New()

	// setup event from cli args or reading from file
	event := events.APIGatewayProxyRequest{
		Body: `{"order":1002}`,
	}

	resp, err := handler(ctx, event)

	log.Printf(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	os.Exit(0)
}
