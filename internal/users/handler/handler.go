package handler

import (
	"main/data/entity"
	"main/internal/users/usecase"
	_ "net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: usecase}
}

func (uh *UserHandler) AddUser(c *fiber.Ctx) (err error) {
	user := new(entity.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := uh.userUsecase.AddUser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}

func (uh *UserHandler) Login(c *fiber.Ctx) (err error) {
	user := new(entity.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invalid input",
		})
	}

	token, err := uh.userUsecase.Login(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invalid input",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func (uh *UserHandler) Logout(c *fiber.Ctx) (err error) {
	return c.JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}