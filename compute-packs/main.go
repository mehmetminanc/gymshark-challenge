package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/mehmetminanc/gymshark-challenge/compute-packs/handlers"
)

func main() {
	// https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html
	// New returns handler function this is useful for passing
	// in environment configurations and dependencies

	handler := handlers.New()
	lambda.Start(handler)
}
