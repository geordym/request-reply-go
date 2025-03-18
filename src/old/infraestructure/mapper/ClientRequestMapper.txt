package mapper

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/geordy/request-reply-lambda-go/src/domain/models"
)

func AwsRequestToDomainRequest(request events.APIGatewayProxyRequest) models.ClientRequest {
	return models.ClientRequest{
		Body:                  request.Body,
		HTTPMethod:            request.HTTPMethod,
		Headers:               request.Headers,
		QueryStringParameters: request.QueryStringParameters,
		PathParameters:        request.PathParameters,
		Cookies:               request.MultiValueHeaders["Cookie"],
		Authorization:         request.Headers["Authorization"],
		Origin:                request.Headers["Origin"],
		Stage:                 request.RequestContext.Stage,
	}
}

func DomainResponseToAwsResponse(request models.ServerResponse) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:        request.StatusCode,
		Body:              request.Body,
		Headers:           request.Headers,
		MultiValueHeaders: request.MultiValueHeaders,
	}
}
