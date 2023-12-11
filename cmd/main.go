package main

import (
	"fmt"
	"log"
	"os"

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

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
