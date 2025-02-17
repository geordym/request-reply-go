package main

import (
    "fmt"
    "context"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    name := request.QueryStringParameters["name"]
    if name == "" {
        name = "Mundo"
    }

    message := fmt.Sprintf("¡Hola, %s!", name)

    return events.APIGatewayProxyResponse{
        StatusCode: 200,
        Body:       message,
    }, nil
}

func main() {
    lambda.Start(handler)
}