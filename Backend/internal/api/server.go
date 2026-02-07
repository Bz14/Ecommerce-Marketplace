package api

import (
	"ecommerce/config"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config *config.AppConfig) {
	app := fiber.New()

	app.Get("/health", HealthCheck)

	fmt.Printf("Starting server on port %s\n", config.Port)
	app.Listen(":" + config.Port)
}

func HealthCheck(c *fiber.Ctx) error{
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "OK",
		"message": "Server is running",
	})
}