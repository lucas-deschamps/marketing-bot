package controller

import "github.com/gofiber/fiber/v2"

type CouponController interface {
	Get(c *fiber.Ctx) error
}
