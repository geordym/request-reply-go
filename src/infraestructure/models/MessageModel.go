package models

type MessageModel struct {
	JobId      string                 `json:"jobId"`
	Payload    string                 `json:"payload"`
	Attributes map[string]interface{} `json:"attributes"`
}
