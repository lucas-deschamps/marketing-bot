package model

// Docs: https://cloud.google.com/dialogflow/es/docs/reference/rest/v2/WebhookResponse
type ChatbotResponse struct {
	FulfillmentText string `json:"fulfillmentText,omitempty"`
}

// type ChatbotPayload struct {
// 	Payload JSON `json:"payload,omitempty"`
// }

type Chatbot interface {
	Message(text string) ChatbotResponse
	Payload(payload any) JSON
}
