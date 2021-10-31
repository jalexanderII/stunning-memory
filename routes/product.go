package routes

import (
	"errors"

	"github.com/jalexanderII/stunning-memory/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/stunning-memory/database"
	"github.com/jalexanderII/stunning-memory/models"
)

// Product To be used as a serializer
type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" validate:"required"`
	SKU   string `json:"sku" validate:"required,sku"`
	Price uint   `json:"price" validate:"gte=0"`
}

// CreateResponseProduct Takes in a model and returns a serializer
func CreateResponseProduct(productModel models.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, SKU: productModel.SKU, Price: productModel.Price}
}

type UpdateProductResponse struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price uint   `json:"price"`
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	responseProduct := CreateResponseProduct(product)
	errs := middleware.ValidateStruct(&responseProduct)
	if errs != nil {
		return c.JSON(errs)
	}
	database.Database.Db.Create(&product)
	responseProduct.ID = product.ID

	return c.Status(fiber.StatusOK).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	responseProducts := make([]Product, len(products))

	database.Database.Db.Find(&products)
	for _, product := range products {
		responseProducts = append(responseProducts, CreateResponseProduct(product))
	}

	return c.Status(fiber.StatusOK).JSON(responseProducts)
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {

		return errors.New("product does not exist")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	var product models.Product

	if err := findProduct(id, &product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	responseProduct := CreateResponseProduct(product)

	return c.Status(fiber.StatusOK).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product
	var updateProductResponse UpdateProductResponse

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	if err = findProduct(id, &product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err = c.BodyParser(&updateProductResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	product.Name = updateProductResponse.Name
	product.SKU = updateProductResponse.SKU
	product.Price = updateProductResponse.Price
	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)
	return c.Status(fiber.StatusOK).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	var product models.Product

	if err := findProduct(id, &product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := database.Database.Db.Where("id = ?", id).Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	responseProduct := CreateResponseProduct(product)
	return c.Status(fiber.StatusOK).JSON(responseProduct)
}
