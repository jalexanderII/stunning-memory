package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/stunning-memory/config"
	"github.com/jalexanderII/stunning-memory/database"
	"github.com/jalexanderII/stunning-memory/middleware"
	"github.com/jalexanderII/stunning-memory/routes"
)

func main() {
	database.ConnectDb()

	app := fiber.New()
	middleware.FiberMiddleware(app)
	routes.SetupRoutes(app)

	// Start server (with graceful shutdown).
	config.StartServerWithGracefulShutdown(app)
}
