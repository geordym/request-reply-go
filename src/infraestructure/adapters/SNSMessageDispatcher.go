package adapters

import (
	"context"
	"encoding/json"

	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/geordy/request-reply-lambda-go/src/domain/models"
)

type SNSMessageDispatcher struct{}

func (s *SNSMessageDispatcher) EnviarMensaje(message models.MessageModel) (*string, error) {

	topicArn := ""
	region := "us-east-1"

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("error al cargar configuración de AWS: %v", err)
	}

	client := sns.NewFromConfig(cfg)

	// Serializar Payload a JSON
	messageBody, err := json.Marshal(message.Payload)
	if err != nil {
		return nil, fmt.Errorf("error al serializar payload: %v", err)
	}

	// Construir atributos dinámicos
	/*messageAttributes := map[string]sns.MessageAttributeValue{}
	for key, value := range message.Attributes {
		// Realizar la assertión de tipo a string
		strValue, ok := value.(string)
		if !ok {
			// Si no es una cadena, manejar el error o continuar
			return nil, nil
		}

		// Asignar el valor convertido a aws.String
		messageAttributes[key] = sns.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(strValue),
		}
	} */

	resp, err := client.Publish(context.TODO(), &sns.PublishInput{
		Message:           aws.String(string(messageBody)),
		TopicArn:          aws.String(topicArn),
		MessageAttributes: nil,
	})

	if err != nil {
		return nil, fmt.Errorf("error al publicar mensaje en SNS: %v", err)
	}

	fmt.Println("Mensaje publicado con ID:", *resp.MessageId)
	return resp.MessageId, nil
}
