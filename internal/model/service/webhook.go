package service

import "github.com/gofiber/fiber/v2"

type WebhookService interface {
	ProcessMessage(c *fiber.Ctx) error
}
