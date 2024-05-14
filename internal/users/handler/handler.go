package handler

import (
	"main/data/entity"
	"main/internal/users/usecase"
	_ "net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	//destroy the token

	return c.JSON(fiber.Map{
		"message": "You have been logged out",
	})
}

func (uh *UserHandler) AddToCart(c *fiber.Ctx) (err error) {
	//add book to cart
	shoppingCart := new(entity.ShoppingCart)
	if err := c.BodyParser(shoppingCart); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invalid input",
		})
	}

	userID := c.Locals("user_id")
	shoppingCart.UserID = uuid.MustParse(userID.(string))

	if err := uh.userUsecase.AddToCart(shoppingCart); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invalid input",
		})
	}


	return c.JSON(fiber.Map{
		"message": "Book added to cart",
	})
}

func (uh *UserHandler) GetCart(c *fiber.Ctx) (err error) {
	//get all books in cart
	userID := c.Locals("user_id")
	shoppingCart, err := uh.userUsecase.GetCart(userID.(string))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "",
		})
	}

	return c.JSON(shoppingCart)
}

func (uh *UserHandler) RemoveFromCart(c *fiber.Ctx) (err error) {
	//remove book from cart
	shoppingCart := new(entity.ShoppingCart)
	if err := c.BodyParser(shoppingCart); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invalid input",
		})
	}

	userID := c.Locals("user_id")
	shoppingCart.UserID = uuid.MustParse(userID.(string))

	if err := uh.userUsecase.RemoveFromCart(shoppingCart); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invalid input",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Book removed from cart",
	})
}

func (uh *UserHandler) CheckDailySellRecord(c *fiber.Ctx) (err error) {
	//check daily sell record
	records, err := uh.userUsecase.CheckDailySellRecord()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
		})
	}

	if len(records) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "No records found",
		})
	}

	return c.JSON(records)
}