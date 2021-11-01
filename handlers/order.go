package handlers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/stunning-memory/database"
	"github.com/jalexanderII/stunning-memory/middleware"
	"github.com/jalexanderII/stunning-memory/models"
)

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user" validate:"dive"`
	Product   Product   `json:"product" validate:"dive"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product, CreatedAt: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	var user models.User
	var product models.Product

	if err := CheckToken(c); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := findUser(order.UserRef, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := findProduct(order.ProductRef, &product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	responseOrder := CreateResponseOrder(
		order,
		CreateResponseUser(user),
		CreateResponseProduct(product),
	)
	errs := middleware.ValidateStruct(&responseOrder)
	if errs != nil {
		return c.JSON(errs)
	}
	database.Database.Db.Create(&order)
	responseOrder.ID = order.ID

	return c.Status(fiber.StatusOK).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	var orders []models.Order
	responseOrders := make([]Order, len(orders))
	database.Database.Db.Find(&orders)

	for _, order := range orders {
		var user models.User
		var product models.Product
		database.Database.Db.Find(&user, "id = ?", order.UserRef)
		database.Database.Db.Find(&product, "id = ?", order.ProductRef)
		responseOrders = append(responseOrders, CreateResponseOrder(
			order,
			CreateResponseUser(user),
			CreateResponseProduct(product),
		))
	}

	return c.Status(fiber.StatusOK).JSON(responseOrders)
}

func findOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	var order models.Order
	var user models.User
	var product models.Product

	if err := findOrder(id, &order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	database.Database.Db.First(&user, order.UserRef)
	database.Database.Db.First(&product, order.ProductRef)

	responseOrder := CreateResponseOrder(
		order,
		CreateResponseUser(user),
		CreateResponseProduct(product),
	)

	return c.Status(fiber.StatusOK).JSON(responseOrder)
}
