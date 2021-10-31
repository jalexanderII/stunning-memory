package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"time"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func SetupRoutes(app *fiber.App){
	// monitoring api stats
	app.Get("/dashboard", monitor.New())
	// entrypoint
	app.Get("/", welcome)
	// User endpoints
	api := app.Group("/api")
	api.Post("/users", CreateUser)
	api.Get("/users", timeout.New(GetUsers, 5 * time.Second))
	api.Get("/users/:id", GetUser)
	api.Put("/users/:id", UpdateUser)
	api.Delete("/users/:id", DeleteUser)
	// Product endpoints
	api.Post("/products", CreateProduct)
	api.Get("/products", timeout.New(GetProducts, 5 * time.Second))
	api.Get("/products/:id", GetProduct)
	api.Put("/products/:id", UpdateProduct)
	api.Delete("/products/:id", DeleteProduct)
	// Order endpoints
	api.Post("/orders", CreateOrder)
	api.Get("/orders", timeout.New(GetOrders, 5 * time.Second))
	api.Get("/orders/:id", GetOrder)
}

