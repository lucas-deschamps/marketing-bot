package webhook

import (
	"github.com/gofiber/fiber/v2"

	"github.com/lucas-deschamps/marketing-bot/internal/controller/webhook"

	. "github.com/lucas-deschamps/marketing-bot/internal/model/controller"
)

func SetupWebhookRoutes(r fiber.Router) {
	router := r.Group("/webhook")

	var controller WebhookController = webhook.NewController()

	router.Post("/", controller.Send)
}
