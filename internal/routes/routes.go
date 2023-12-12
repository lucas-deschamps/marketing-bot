package routes

import (
	"github.com/gofiber/fiber/v2"

	. "github.com/lucas-deschamps/marketing-bot/internal/routes/coupon"
	. "github.com/lucas-deschamps/marketing-bot/internal/routes/webhook"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	SetupWebhookRoutes(v1)
	SetupCouponRoutes(v1)
}
