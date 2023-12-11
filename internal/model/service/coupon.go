package service

import "github.com/gofiber/fiber/v2"

type CouponService interface {
	Track(c *fiber.Ctx) error
}
