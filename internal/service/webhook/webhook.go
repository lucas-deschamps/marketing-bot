package webhook

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	. "github.com/lucas-deschamps/marketing-bot/internal/model"
	. "github.com/lucas-deschamps/marketing-bot/internal/service/chatbot"

	"github.com/lucas-deschamps/marketing-bot/internal/pkg/chatbot"
)

type webhookService struct{}

func NewService() *webhookService {
	return &webhookService{}
}

func (s *webhookService) ProcessMessage(c *fiber.Ctx) error {
	fmt.Println(string(c.Body()))

	jsonReq := new(DialogFlowRequest)

	if err := c.BodyParser(jsonReq); err != nil {
		fmt.Println("Error when parsing incoming body request: ", err)
		return c.SendStatus(500)
	}

	action := jsonReq.QueryResult.Action
	parameters := jsonReq.QueryResult.Parameters

	fmt.Println(action, parameters)

	var bot Chatbot = chatbot.New()

	if response := s.VerifyUserInteraction(action, parameters); response != nil {
		fmt.Println("responseeee: ", response)

		botResponse := bot.Payload(response)

		return c.JSON(botResponse)
	} else {
		msg := bot.Message("Sorry, I didn't understand what you said.")
		return c.JSON(msg)
	}
}

func (s *webhookService) VerifyUserInteraction(action string, parameter Parameters) any {
	switch action {
	case "welcome_name":
		return CouponService(parameter)
	}

	return nil
}
