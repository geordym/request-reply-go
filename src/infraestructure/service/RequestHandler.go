package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/adapters"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/configuration"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/models"
	"github.com/google/uuid"
)

type RequestHandler struct {
}

func NewRequestHandler() (*RequestHandler, error) {
	return &RequestHandler{}, nil
}

func (h *RequestHandler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var pdfRequest PDFRequest

	err := json.Unmarshal([]byte(request.Body), &pdfRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       fmt.Sprintf("Error al deserializar el cuerpo de la solicitud: %s", err.Error()),
		}, nil
	}

	log.Println(pdfRequest)
	pdfRequestJSON, err := json.Marshal(pdfRequest)
	if err != nil {
		log.Fatal("Error al convertir PDFRequest a JSON:", err)
	}

	var messageRequest = models.MessageModel{
		JobId:   uuid.New().String(),
		Payload: string(pdfRequestJSON),
	}

	messageRequestJson, err := json.Marshal(messageRequest)
	if err != nil {
		log.Fatal("Error al convertir messageRequest a JSON:", err)
	}

	sqsMessagePublisher, err := adapters.NewSQSMessagePublisher()
	if err != nil {
		log.Fatal("Error al obtener JobResult:", err)
	}

	messageRequestJsonString := string(messageRequestJson)
	log.Println(messageRequestJsonString)

	var messageId, err1 = sqsMessagePublisher.PublishMessage(messageRequestJsonString, h.findSQSUrlToProccessRequest(pdfRequest.Target))
	if err1 != nil {
		log.Fatal("Error al obtener message:", err)
	}

	fmt.Printf("PDFRequest: %+v\n", pdfRequest)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Solicitud procesada con éxito " + messageId,
	}, nil
}

func (h *RequestHandler) findSQSUrlToProccessRequest(targetKey string) string {
	target, err := configuration.FindTargetByKey(targetKey)
	if err != nil {
		log.Fatal("Error al obtener el target:", err)
	}

	if url, exists := target.TARGET_CONFIG["Url"]; exists {
		return url.(string)
	}

	log.Fatalf("La clave 'Url' no existe en el target con la clave '%s'", targetKey)
	return ""
}

type PDFRequest struct {
	Target        string            `json:"Target"`
	Configuration map[string]string `json:"Configuration"`
	Payload       map[string]string `json:"Payload"`
	Metadata      map[string]string `json:"Metadata"`
}
