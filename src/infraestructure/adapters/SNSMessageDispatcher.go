package adapters

import (
	"context"
	"encoding/json"

	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNSMessageDispatcher struct{}

func NuevoSNSMessageDispatcher(region, topicArn string) (*SNSMessageDispatcher, error) {

	return &SNSMessageDispatcher{Client: client, TopicArn: topicArn}, nil
}

// EnviarMensaje publica un mensaje en SNS con atributos
func (s *SNSMessageDispatcher) EnviarMensaje(message MessageModel) (string, error) {

	topicArn := ""

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("error al cargar configuración de AWS: %v", err)
	}

	client := sns.NewFromConfig(cfg)

	// Serializar Payload a JSON
	messageBody, err := json.Marshal(message.Payload)
	if err != nil {
		return "", fmt.Errorf("error al serializar payload: %v", err)
	}

	// Construir atributos dinámicos
	messageAttributes := map[string]sns.MessageAttributeValue{}
	for key, value := range message.Attributes {
		messageAttributes[key] = sns.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(value),
		}
	}

	// Publicar mensaje en SNS
	resp, err := client.Publish(context.TODO(), &sns.PublishInput{
		Message:           aws.String(string(messageBody)),
		TopicArn:          aws.String(s.topicArn),
		MessageAttributes: messageAttributes,
	})

	if err != nil {
		return "", fmt.Errorf("error al publicar mensaje en SNS: %v", err)
	}

	fmt.Println("Mensaje publicado con ID:", *resp.MessageId)
	return *resp.MessageId, nil
}
