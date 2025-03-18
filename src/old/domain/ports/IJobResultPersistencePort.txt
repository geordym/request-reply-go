package ports

import "github.com/geordy/request-reply-lambda-go/src/domain/models"

type IJobResultPersistencePort interface {
	FindJobResultByJobId(jobId string) (response models.JobResultModel, err error)
}
