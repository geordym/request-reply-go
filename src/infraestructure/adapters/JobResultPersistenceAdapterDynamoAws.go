package adapters

import (
	"github.com/geordy/request-reply-lambda-go/src/domain/models"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/mapper"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/repository"
)

type JobResultPersistenceAdapterDynamoAws struct {
	jobResultRepository repository.IJobResultRepository
}

func NewJobResultPersistenceAdapterDynamoAws(jobResultRepository repository.IJobResultRepository) *JobResultPersistenceAdapterDynamoAws {
	return &JobResultPersistenceAdapterDynamoAws{
		jobResultRepository: jobResultRepository,
	}
}

func (j *JobResultPersistenceAdapterDynamoAws) FindJobResultByJobId(jobId string) (models.JobResultModel, error) {
	result, err := j.jobResultRepository.FindJobResultByJobId(jobId)
	if err != nil {
		return models.JobResultModel{}, err
	}

	return mapper.ToJobResultModel(result), nil
}
