package handler

import (
	"main/data/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddBook(c *fiber.Ctx, db *gorm.DB) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims["user_type"] != "Owner" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	book := new(model.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	existingBook := &model.Book{}
    if err := db.Where("title = ?", book.Title).First(existingBook).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            // Book does not exist, create a new one
			book.ID = uuid.New()

            db.Create(&book)
        } else {
            // An error occurred while checking for the book
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "Error checking for existing book",
            })
        }
    } else {
        // Book exists, update it
		existingBook.Stock += book.Stock
        db.Model(existingBook).Updates(existingBook)
    }

    return c.JSON(book)

}