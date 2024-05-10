package handler

import (
	"main/data/entity"
	"main/internal/books/usecase"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	usecase *usecase.BookUsecase
}

func NewBookHandler(usecase *usecase.BookUsecase) *BookHandler {
	return &BookHandler{usecase: usecase}

}

func (bh *BookHandler) AddBook(c *fiber.Ctx) (err error) {
	if c.Locals("user_type") != "Owner" {
		return c.SendStatus(403)
	}

	book := new(entity.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := bh.usecase.AddBook(book); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(book)
}
