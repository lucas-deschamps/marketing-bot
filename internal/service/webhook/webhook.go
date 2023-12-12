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
	jsonReq := new(DialogFlowRequest)

	if err := c.BodyParser(jsonReq); err != nil {
		fmt.Println("Error when parsing incoming body request: ", err)
		return c.SendStatus(500)
	}

	action := jsonReq.QueryResult.Action
	parameters := jsonReq.QueryResult.Parameters

	var bot Chatbot = chatbot.New()

	if response := s.VerifyUserInteraction(action, parameters); response != nil {
		return c.JSON(response)
	} else {
		msg := bot.Message("Sorry, I didn't understand what you said.")
		return c.JSON(msg)
	}
}

func (s *webhookService) VerifyUserInteraction(action string, parameters Parameters) any {
	switch action {
	case "welcome_name":
		return CouponOffer(parameters)
	case "coupon_yes":
		return CouponYes(parameters)
	case "coupon_no":
		return CouponNo(parameters)
	}

	return nil
}
