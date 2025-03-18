package ports

import "github.com/geordy/request-reply-lambda-go/src/domain/models"

type IRequestHandler interface {
	HandleRequest(request models.ClientRequest) (response models.ServerResponse, err error)
}
