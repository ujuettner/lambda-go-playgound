package main

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrUnexpectedBody indicates that the HTTP body does not conform our expectations.
	ErrUnexpectedBody = errors.New("Unexpected HTTP body")
)

// RequestHandler handles the request arriving from API Gateway.
func RequestHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrUnexpectedBody
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hi there, " + request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(RequestHandler)
}
