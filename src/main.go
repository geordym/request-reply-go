package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/configuration"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/service"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	print("PROCESANDO PETICION")
	configuration.LoadConfig()
	err := configuration.InitializeTargets(configuration.TARGETS_FILE_PATH)
	if err != nil {
		log.Fatalf("Error al inicializar los targets: %v", err)
	}

	handler, err2 := service.NewRequestHandler()
	if err2 != nil {
		log.Fatal("Error al obtener NewRequestHandler:", err2)
	}

	handler.HandleRequest(request)

	/* sqsMessagePublisher, err := adapters.NewSQSMessagePublisher()
	if err != nil {
		log.Fatal("Error al obtener JobResult:", err)
	}
	*/

	/*var messageId, err1 = sqsMessagePublisher.PublishMessage("probando", configuration.JOB_QUEUE_DUMPDF_URL)
	if err1 != nil {
		log.Fatal("Error al obtener message:", err)
	}*/

	//print(messageId)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "¡Hola desde Lambda en Go!",
	}, nil
}

func main() {
	lambda.Start(handler)
}
