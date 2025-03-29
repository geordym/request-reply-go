package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/adapters"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/configuration"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/models"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/sqs/util"
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

	var target = pdfRequest.Target

	var messageRequest = models.MessageModel{
		JobId:   uuid.New().String(),
		Payload: string(pdfRequestJSON),
	}

	messageRequestSerializer, err3 := h.findMessageSerializerByTarget(target)
	if err3 != nil {
		log.Fatal("Error al convertir PDFRequest a JSON:", err3)
	}

	messageRequestSerializedJson := messageRequestSerializer.SerializeMessage(messageRequest)

	messageRequestJsonString := string(messageRequestSerializedJson)
	log.Println(messageRequestJsonString)

	sqsMessagePublisher, err := adapters.NewSQSMessagePublisher()
	if err != nil {
		log.Fatal("Error al obtener JobResult:", err)
	}

	var messageId, err1 = sqsMessagePublisher.PublishMessage(messageRequestJsonString, h.findSQSUrlToProccessRequest(pdfRequest.Target))
	if err1 != nil {
		log.Fatal("Error al obtener message:", err1)
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

func (h *RequestHandler) findMessageSerializerByTarget(target string) (util.ISQSMessageSerializer, error) {
	t, err := configuration.FindTargetByKey(target)
	if err != nil {
		return nil, fmt.Errorf("error finding target: %v", err)
	}

	var messageSerializer util.ISQSMessageSerializer

	switch target {
	case "DOMPDF":
		messageSerializer = util.LaravelSQSMessageSerializer{}
	default:
		return nil, fmt.Errorf("no serializer found for MESSAGE_TEMPLATE: %s", t.MESSAGE_TEMPLATE)
	}

	return messageSerializer, nil
}

type PDFRequest struct {
	Target        string            `json:"Target"`
	Configuration map[string]string `json:"Configuration"`
	Payload       map[string]string `json:"Payload"`
	Metadata      map[string]string `json:"Metadata"`
}
