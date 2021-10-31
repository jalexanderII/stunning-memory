package routes

import "github.com/gofiber/fiber/v2"

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func SetupRoutes(app *fiber.App){
	app.Get("/", welcome)
	// User endpoints
	api := app.Group("/api")
	api.Post("/users", CreateUser)
	api.Get("/users", GetUsers)
	api.Get("/users/:id", GetUser)
	api.Put("/users/:id", UpdateUser)
	api.Delete("/users/:id", DeleteUser)
	// Product endpoints
	api.Post("/products", CreateProduct)
	api.Get("/products", GetProducts)
	api.Get("/products/:id", GetProduct)
	api.Put("/products/:id", UpdateProduct)
	api.Delete("/products/:id", DeleteProduct)
	// Order endpoints
	api.Post("/orders", CreateOrder)
	api.Get("/orders", GetOrders)
	api.Get("/orders/:id", GetOrder)
}

