package bootstrap

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/geordy/request-reply-lambda-go/src/domain/ports"
	"github.com/geordy/request-reply-lambda-go/src/domain/usecase"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/adapters"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/repository"
)

func HandleRequestUseCase() *usecase.HandleRequestUseCase {
	return usecase.NewHandleRequestUseCase(JobResultPersistencePort())
}

func JobResultPersistencePort() ports.IJobResultPersistencePort {
	return adapters.NewJobResultPersistenceAdapterDynamoAws(JobResultRepository())
}

func JobResultRepository() repository.IJobResultRepository {
	client, err := NewDynamoDBClient()
	if err != nil {
		log.Fatalf("Error creando el cliente de DynamoDB: %v", err)
	}

	return repository.NewJobResultDynamoAwsRepository(client, "JobsTable")
}

func NewDynamoDBClient() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("error cargando configuración de AWS: %w", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}
