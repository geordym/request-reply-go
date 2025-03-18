package models

type MessageModel struct {
	ID         string                 `json:"id"`
	Payload    map[string]interface{} `json:"payload"`
	Attributes map[string]interface{} `json:"payload"`
}
