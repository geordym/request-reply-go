package repository

import (
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/entity"

	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type JobResultDynamoAwsRepositoryImpl struct {
	DynamoDBClient *dynamodb.Client
	TableName      string
}

// NewJobResultDynamoAwsRepository crea una nueva instancia del repositorio con inyección de dependencias
func NewJobResultDynamoAwsRepository(client *dynamodb.Client, tableName string) *JobResultDynamoAwsRepositoryImpl {
	return &JobResultDynamoAwsRepositoryImpl{
		DynamoDBClient: client,
		TableName:      tableName,
	}
}

func (j *JobResultDynamoAwsRepositoryImpl) FindJobResultByJobId(jobId string) (entity.JobResultEntity, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(j.TableName),
		Key: map[string]types.AttributeValue{
			"jobId": &types.AttributeValueMemberS{Value: jobId},
		},
	}

	result, err := j.DynamoDBClient.GetItem(context.TODO(), input)
	if err != nil {
		log.Printf("Error al consultar DynamoDB: %v", err)
		return entity.JobResultEntity{}, err
	}

	if result.Item == nil {
		return entity.JobResultEntity{}, errors.New("job result not found")
	}

	var jobResult entity.JobResultEntity
	err = attributevalue.UnmarshalMap(result.Item, &jobResult)
	if err != nil {
		log.Fatal("Error en unmarshal:", err)
	}

	return jobResult, nil
}
