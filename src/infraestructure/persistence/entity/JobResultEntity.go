package models

import "time"

type JobResultEntity struct {
	JobID       string                 `dynamodbav:"jobId"`
	Status      string                 `dynamodbav:"status"`
	CreatedAt   time.Time              `dynamodbav:"createdAt"`
	StartedAt   *time.Time             `dynamodbav:"startedAt,omitempty"`
	CompletedAt *time.Time             `dynamodbav:"completedAt,omitempty"`
	Duration    int64                  `dynamodbav:"duration"`
	WorkerID    string                 `dynamodbav:"workerId"`
	InputData   map[string]interface{} `dynamodbav:"inputData"`
	ResultType  string                 `dynamodbav:"resultType"`
	Result      map[string]interface{} `dynamodbav:"result"`
	Error       *string                `dynamodbav:"error,omitempty"`
}
