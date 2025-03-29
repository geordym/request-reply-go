package util

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/geordy/request-reply-lambda-go/src/infraestructure/models"
)

type ISQSMessageSerializer interface {
	SerializeMessage(messageModel models.MessageModel) string
}

type LaravelSQSMessageSerializer struct{}

func (l LaravelSQSMessageSerializer) SerializeMessage(messageModel models.MessageModel) string {

	messageModelJson, err1 := json.Marshal(messageModel)
	if err1 != nil {
		log.Fatalf("Error serializing message: %v", err1)
	}

	messageModelJsonLength := len(messageModelJson) // Obtiene la longitud correcta

	serializedCommand := fmt.Sprintf(
		"O:23:\"App\\Jobs\\ProcessPodcast\":1:{s:34:\"\u0000App\\Jobs\\ProcessPodcast\u0000podcastId\";s:%d:\"%s\";}",
		messageModelJsonLength, messageModelJson,
	)

	message := map[string]interface{}{
		"uuid":        messageModel.JobId,
		"displayName": "work",
		"job":         "Illuminate\\Queue\\CallQueuedHandler@call",
		"maxTries":    nil,
		"timeout":     nil,
		"data": map[string]interface{}{
			"commandName": "App\\Jobs\\ProcessPodcast",
			"command":     serializedCommand,
		},
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Error serializing message: %v", err)
	}
	return string(jsonData)
}
