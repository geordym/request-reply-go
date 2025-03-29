package adapters

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/configuration"
)

type SQSMessagePublisher struct {
	client *sqs.Client
}

func NewSQSMessagePublisher() (*SQSMessagePublisher, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(configuration.LAMBDA_AWS_REGION))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	client := sqs.NewFromConfig(cfg)

	return &SQSMessagePublisher{
		client: client,
	}, nil
}

func (j *SQSMessagePublisher) PublishMessage(message string, queueUrl string) (string, error) {

	log.Println("La url del SQS ES: " + queueUrl)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := j.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    &queueUrl,
		MessageBody: &message,
	})

	if err != nil {
		return "", fmt.Errorf("unable to send message to SQS, %v", err)
	}

	fmt.Println("Estoy publicando el mensaje")
	return *resp.MessageId, nil
}
