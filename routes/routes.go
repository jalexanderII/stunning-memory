package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/jalexanderII/stunning-memory/handlers"
	"github.com/jalexanderII/stunning-memory/middleware"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func SetupRoutes(app *fiber.App){
	app.Get("/", welcome)
	api := app.Group("/api")

	// monitoring api stats
	api.Get("/dashboard", monitor.New())

	// Auth
	auth := api.Group("/auth")
	auth.Get("/token/new", handlers.GetNewAccessToken)
	auth.Post("/login", handlers.Login)

	// User endpoints
	users := api.Group("/users")
	users.Post("/", handlers.CreateUser)
	users.Get("/", handlers.GetUsers)
	users.Get("/:id", handlers.GetUser)
	users.Patch("/:id",  middleware.Protected(), handlers.UpdateUser)
	users.Delete("/:id", middleware.Protected(), handlers.DeleteUser)
	// Product endpoints
	products := api.Group("/products")
	products.Get("/", handlers.GetProducts)
	products.Get("/:id", handlers.GetProduct)
	products.Post("/",  middleware.Protected(), handlers.CreateProduct)
	products.Patch("/:id",  middleware.Protected(), handlers.UpdateProduct)
	products.Delete("/:id",  middleware.Protected(), handlers.DeleteProduct)
	// Order endpoints
	orders := api.Group("/orders")
	orders.Get("/", handlers.GetOrders)
	orders.Get("/:id", handlers.GetOrder)
	orders.Post("/",  middleware.Protected(), handlers.CreateOrder)
}

