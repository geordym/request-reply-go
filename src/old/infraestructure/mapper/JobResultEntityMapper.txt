package mapper

import (
	"time"

	"github.com/geordy/request-reply-lambda-go/src/domain/models"
	"github.com/geordy/request-reply-lambda-go/src/infraestructure/entity"
)

func ToJobResultModel(jobResultEntity entity.JobResultEntity) models.JobResultModel {
	// Convertir strings a time.Time
	createdAt, _ := time.Parse(time.RFC3339, jobResultEntity.CreatedAt)
	var startedAt, completedAt *time.Time

	if jobResultEntity.StartedAt != nil {
		parsedStartedAt, _ := time.Parse(time.RFC3339, *jobResultEntity.StartedAt)
		startedAt = &parsedStartedAt
	}

	if jobResultEntity.CompletedAt != nil {
		parsedCompletedAt, _ := time.Parse(time.RFC3339, *jobResultEntity.CompletedAt)
		completedAt = &parsedCompletedAt
	}

	return models.JobResultModel{
		JobID:       jobResultEntity.JobID,
		Status:      jobResultEntity.Status,
		CreatedAt:   createdAt,
		StartedAt:   startedAt,
		CompletedAt: completedAt,
		Duration:    jobResultEntity.Duration,
		WorkerID:    jobResultEntity.WorkerID,
		InputData:   jobResultEntity.InputData,
		ResultType:  jobResultEntity.ResultType,
		Result:      jobResultEntity.Result,
		Error:       jobResultEntity.Error,
	}
}

func ToJobResultEntity(jobResultModel models.JobResultModel) entity.JobResultEntity {
	createdAt := jobResultModel.CreatedAt.Format(time.RFC3339)
	var startedAt, completedAt *string

	if jobResultModel.StartedAt != nil {
		startedStr := jobResultModel.StartedAt.Format(time.RFC3339)
		startedAt = &startedStr
	}

	if jobResultModel.CompletedAt != nil {
		completedStr := jobResultModel.CompletedAt.Format(time.RFC3339)
		completedAt = &completedStr
	}

	return entity.JobResultEntity{
		JobID:       jobResultModel.JobID,
		Status:      jobResultModel.Status,
		CreatedAt:   createdAt,
		StartedAt:   startedAt,
		CompletedAt: completedAt,
		Duration:    jobResultModel.Duration,
		WorkerID:    jobResultModel.WorkerID,
		InputData:   jobResultModel.InputData,
		ResultType:  jobResultModel.ResultType,
		Result:      jobResultModel.Result,
		Error:       jobResultModel.Error,
	}
}
