package chatbot

import . "github.com/lucas-deschamps/marketing-bot/internal/model"

type chatbot struct{}

func New() *chatbot {
	return &chatbot{}
}

func (c *chatbot) Message(text string, payload JSON) ChatbotResponse {
	response := ChatbotResponse{
		FulfillmentText: text,
		Payload:         payload,
	}

	return response
}
