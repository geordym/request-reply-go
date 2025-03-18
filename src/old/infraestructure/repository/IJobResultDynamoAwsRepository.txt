package repository

import (
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/entity"
)

type IJobResultRepository interface {
	FindJobResultByJobId(jobId string) (entity.JobResultEntity, error)
}
