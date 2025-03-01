package entity

type JobResultEntity struct {
	JobID       string                 `dynamodbav:"jobId"`                 // Partition Key
	Status      string                 `dynamodbav:"status"`                // Estado del job
	CreatedAt   string                 `dynamodbav:"createdAt"`             // Fecha de creación (ISO 8601)
	StartedAt   *string                `dynamodbav:"startedAt,omitempty"`   // Fecha de inicio (ISO 8601)
	CompletedAt *string                `dynamodbav:"completedAt,omitempty"` // Fecha de finalización (ISO 8601)
	Duration    int64                  `dynamodbav:"duration"`              // Duración en segundos
	WorkerID    string                 `dynamodbav:"workerId"`              // ID del worker
	InputData   map[string]interface{} `dynamodbav:"inputData"`             // Datos de entrada
	ResultType  string                 `dynamodbav:"resultType"`            // Tipo de resultado
	Result      map[string]interface{} `dynamodbav:"result"`                // Resultado del job
	Error       *string                `dynamodbav:"error,omitempty"`       // Mensaje de error si falló
}
