package webhook

import (
	"github.com/gofiber/fiber/v2"

	"github.com/lucas-deschamps/marketing-bot/internal/service/webhook"

	. "github.com/lucas-deschamps/marketing-bot/internal/model/service"
)

type webhookController struct{}

func NewController() *webhookController {
	return &webhookController{}
}

func (ctrl *webhookController) Send(c *fiber.Ctx) error {
	var service WebhookService = webhook.NewService()

	return service.ProcessMessage(c)
}
