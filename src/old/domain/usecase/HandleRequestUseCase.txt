package usecase

import (
	"fmt"

	"github.com/geordy/request-reply-lambda-go/src/domain/models"
	"github.com/geordy/request-reply-lambda-go/src/domain/ports"
)

type HandleRequestUseCase struct {
	jobResultPersistencePort ports.IJobResultPersistencePort
}

func NewHandleRequestUseCase(jobResultPersistencePort ports.IJobResultPersistencePort) *HandleRequestUseCase {
	return &HandleRequestUseCase{jobResultPersistencePort: jobResultPersistencePort}
}

func (h *HandleRequestUseCase) HandleRequest(request models.ClientRequest) (models.ServerResponse, error) {

	fmt.Print("Entrando a usecase")

	jobResult, err := h.jobResultPersistencePort.FindJobResultByJobId("job1")
	if err != nil {
		fmt.Println("Error al obtener JobResult:", err)
		return models.ServerResponse{}, err // Retorna el error
	}

	print("Status: " + jobResult.Status)

	// Construcción de la respuesta
	response := models.ServerResponse{
		Body:       "Request processed successfully",
		StatusCode: 200,
	}

	return response, nil
}
