package ports

import "github.com/geordy/request-reply-lambda-go/src/domain/models"

type IMessageDispatcher interface {
	SendMessage(message models.MessageModel) (messageId string, err error)
}
