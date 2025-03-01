package models

type MessageModel struct {
	ID         string                 `json:"id"`      // Identificador único del mensaje
	Payload    map[string]interface{} `json:"payload"` // Datos específicos del evento
	Attributes map[string]interface{} `json:"payload"` // Datos específicos del evento
}
