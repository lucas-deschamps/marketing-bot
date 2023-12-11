package coupon

import (
	"github.com/gofiber/fiber/v2"

	"github.com/lucas-deschamps/marketing-bot/internal/controller/coupon"

	. "github.com/lucas-deschamps/marketing-bot/internal/model/controller"
)

func SetupCouponRoutes(r fiber.Router) {
	router := r.Group("/coupon")

	var controller CouponController = coupon.NewController()

	router.Get("/", controller.Get)
}
