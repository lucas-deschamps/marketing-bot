package coupon

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/lucas-deschamps/marketing-bot/internal/model/service"
	"github.com/lucas-deschamps/marketing-bot/internal/service/coupon"
)

type couponController struct{}

func NewController() *couponController {
	return &couponController{}
}

func (ctrl *couponController) Get(c *fiber.Ctx) error {
	var service CouponService = coupon.NewService()

	return service.Track(c)
}
