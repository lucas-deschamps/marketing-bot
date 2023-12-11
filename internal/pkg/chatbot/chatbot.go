package chatbot

import . "github.com/lucas-deschamps/marketing-bot/internal/model"

type chatbot struct{}

func New() *chatbot {
	return &chatbot{}
}

func (c *chatbot) Message(text string) ChatbotResponse {
	response := ChatbotResponse{FulfillmentText: text}

	return response
}

func (c *chatbot) Payload(payload any) JSON {
	payloadObj := JSON{"payload": payload}

	return payloadObj
}
