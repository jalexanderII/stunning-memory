package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/stunning-memory/config"
	"github.com/jalexanderII/stunning-memory/database"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", welcome)

	config.Logger.Info("Connecting to server")
	app.Listen(":9092")
}
