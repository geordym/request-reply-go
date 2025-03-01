package models

import (
	"github.com/aws/aws-lambda-go/events"
)

type ClientRequest struct {
	Body string `json:"body"`

	HTTPMethod string `json:"httpMethod"`

	Headers map[string]string `json:"headers"`

	QueryStringParameters map[string]string `json:"queryStringParameters"`

	PathParameters map[string]string `json:"pathParameters"`

	Cookies []string `json:"cookies"`

	Authorization string `json:"authorization"`

	Origin string `json:"origin"`

	Stage string `json:"stage"`
}

func ToDomain(request events.APIGatewayProxyRequest) ClientRequest {
	return ClientRequest{
		Body:                  request.Body,
		HTTPMethod:            request.HTTPMethod,
		Headers:               request.Headers,
		QueryStringParameters: request.QueryStringParameters,
		PathParameters:        request.PathParameters,
		Authorization:         request.Headers["Authorization"],
		Origin:                request.Headers["Origin"],
		Stage:                 request.RequestContext.Stage,
	}
}
