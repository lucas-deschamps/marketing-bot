package model

// Docs: https://cloud.google.com/dialogflow/es/docs/reference/rest/v2/WebhookResponse
type ChatbotResponse struct {
	FulfillmentText string `json:"fulfillmentText,omitempty"`
	Payload         JSON   `json:"payload,omitempty"`
}

type Chatbot interface {
	Message(text string, payload JSON) ChatbotResponse
}
