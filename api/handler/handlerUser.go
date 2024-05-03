package handler

import (
	"main/data/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(c *fiber.Ctx, db *gorm.DB) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	user.ID= uuid.New()

	switch user.UserType {
	case strconv.Itoa(1):
		user.UserType = "Customer"
	case strconv.Itoa(2):
		user.UserType = "Owner"
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user type",
		})
	}

	db.Create(&user)

	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx, db *gorm.DB) error {
	var users []model.User

	db.Find(&users)

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx, db *gorm.DB) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "Could not parse claims",
		})
	}

	username, ok := claims["username"].(string)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "Username not found in claims",
		})
	}

	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(user)
}