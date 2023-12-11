package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/lucas-deschamps/marketing-bot/internal/routes"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
