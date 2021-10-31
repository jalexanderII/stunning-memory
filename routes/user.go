package routes

import (
	"errors"

	"github.com/jalexanderII/stunning-memory/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/stunning-memory/database"
	"github.com/jalexanderII/stunning-memory/models"
	"gorm.io/gorm/clause"
)

// User To be used as a serializer
type User struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name" validate:"required"`
	Email *string `json:"email" validate:"required,email"`
}

// CreateResponseUser Takes in a model and returns a serializer
func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, Name: userModel.Name, Email: userModel.Email}
}

type UpdateUserResponse struct {
	Name  string  `json:"name"`
	Email *string `json:"email"`
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	responseUser := CreateResponseUser(user)
	errs := middleware.ValidateStruct(&responseUser)
	if errs != nil {
		return c.JSON(errs)
	}

	database.Database.Db.Create(&user)
	responseUser.ID = user.ID

	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	responseUsers := make([]User, len(users))

	database.Database.Db.Find(&users)
	for _, user := range users {
		responseUsers = append(responseUsers, CreateResponseUser(user))
	}

	return c.Status(fiber.StatusOK).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {

		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	var user models.User

	if err := findUser(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)

	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	var updateUserResponse UpdateUserResponse

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	if err = findUser(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err = c.BodyParser(&updateUserResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	user.Name = updateUserResponse.Name
	user.Email = updateUserResponse.Email
	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}
	var user models.User

	if err := findUser(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := database.Database.Db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}
