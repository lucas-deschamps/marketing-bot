package controller

import "github.com/gofiber/fiber/v2"

type WebhookController interface {
	Send(c *fiber.Ctx) error
}
