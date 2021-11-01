package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/stunning-memory/database"
	"github.com/jalexanderII/stunning-memory/middleware"
	"github.com/jalexanderII/stunning-memory/models"
	"gorm.io/gorm/clause"
)

// User To be used as a serializer
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username"`
	Email    string `json:"email" validate:"required,email"`
}

// CreateResponseUser Takes in a model and returns a serializer
func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, Name: userModel.Name, Username: userModel.Username, Email: userModel.Email}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := CheckToken(c); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	hash, err := HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}
	user.Password = hash

	responseUser := CreateResponseUser(user)
	errs := middleware.ValidateStruct(&responseUser)
	if errs != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errs,
		})
	}

	if err := database.Database.Db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err.Error()})
	}
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

type UpdateUserResponse struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	var updateUserResponse UpdateUserResponse

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	if err := CheckToken(c); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	if err = findUser(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err = c.BodyParser(&updateUserResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	user.Name = updateUserResponse.Name
	user.Username = updateUserResponse.Username
	user.Email = updateUserResponse.Email
	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	type DeleteUser struct {
		Password string `json:"password"`
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var deleteUser DeleteUser
	if err := CheckToken(c); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	if err := c.BodyParser(&deleteUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if !ValidUser(id, deleteUser.Password) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})
	}
	
	var user models.User
	if err := database.Database.Db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}
