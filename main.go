package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/stunning-memory/config"
	"github.com/jalexanderII/stunning-memory/database"
	"github.com/jalexanderII/stunning-memory/routes"
)

func main() {
	app := fiber.New()
	//app.Use(cors.New(), cache.New())

	database.ConnectDb()
	routes.SetupRoutes(app)

	config.Logger.Info("Connecting to server")
	log.Fatal(app.Listen(":9092"))
}
