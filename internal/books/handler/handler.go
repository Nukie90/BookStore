package handler

import (
	"main/data/entity"
	"main/internal/books/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BookHandler struct {
	usecase *usecase.BookUsecase
}

func NewBookHandler(usecase *usecase.BookUsecase) *BookHandler {
	return &BookHandler{usecase: usecase}

}

func (bh *BookHandler) AddBook(c *fiber.Ctx) (err error) {
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

func (bh *BookHandler) BrowseBook(c *fiber.Ctx) (err error) {
	how := c.Params("how")
	books, err := bh.usecase.BrowseBook(how)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(books)
}

func (bh *BookHandler) Checkout(c *fiber.Ctx) (err error) {
	userID := c.Locals("user_id")
	userID = uuid.MustParse(userID.(string))

	if err := bh.usecase.Checkout(userID.(uuid.UUID)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Checkout successful",
	})
}