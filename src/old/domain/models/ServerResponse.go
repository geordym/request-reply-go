package models

import (
	"github.com/aws/aws-lambda-go/events"
)

type ServerResponse struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
	Body              string              `json:"body"`
}

func toServerResponse(response events.APIGatewayProxyResponse) ServerResponse {
	return ServerResponse{
		StatusCode:        response.StatusCode,
		Headers:           response.Headers,
		MultiValueHeaders: response.MultiValueHeaders,
		Body:              response.Body,
	}
}

func toApiGatewayProxyResponse(response ServerResponse) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:        response.StatusCode,
		Headers:           response.Headers,
		MultiValueHeaders: response.MultiValueHeaders,
		Body:              response.Body,
	}
}
