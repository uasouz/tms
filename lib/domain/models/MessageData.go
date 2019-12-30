package models


type MessageData struct {
	Type    string      `json:"type"`
	Subject string      `json:"subject"`
	To      []string    `json:"destinations"`
	Message interface{} `json:"message"`
}
