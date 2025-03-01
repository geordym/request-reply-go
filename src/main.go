package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/geordy/request-reply-lambda-go/src/bootstrap"
	"github.com/geordy/request-reply-lambda-go/src/domain/models"

	"github.com/geordy/request-reply-lambda-go/src/infraestructure/mapper"
)

func domainHandler(ctx context.Context, request models.ClientRequest) models.ServerResponse {

	return models.ServerResponse{Body: "asd"}
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	handleRequestUseCase := bootstrap.HandleRequestUseCase()

	modelRequest := mapper.AwsRequestToDomainRequest(request)

	modelResponse, err := handleRequestUseCase.HandleRequest(modelRequest)

	if err != nil {
		fmt.Print("No se pudo hacer esto")
	}

	awsResponse := mapper.DomainResponseToAwsResponse(modelResponse)

	return awsResponse, nil
}

func main() {
	lambda.Start(handler)
}
