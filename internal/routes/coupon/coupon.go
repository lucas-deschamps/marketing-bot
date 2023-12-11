package coupon

import (
	"github.com/gofiber/fiber/v2"

	"github.com/lucas-deschamps/marketing-bot/internal/controller/webhook"

	. "github.com/lucas-deschamps/marketing-bot/internal/model/controller"
)

func SetupCouponRoutes(r fiber.Router) {
	router := r.Group("/coupon")

	var controller WebhookController = webhook.NewController()

	router.Post("/", controller.Send)
}
