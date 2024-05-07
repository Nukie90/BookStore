package routes

import (
	"fmt"
	"main/api/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func OwnerRoutes(app *fiber.App, db *gorm.DB) {
    owner := fiber.New()
	fmt.Println("owner")

    app.Mount("/owner", owner)

	owner.Post("/addbook", func(c *fiber.Ctx) error {
		fmt.Println("addbook")
		return handler.AddBook(c, db)
	})

	owner.Get("/soldrecord/:how", func(c *fiber.Ctx) error {
		fmt.Println("getbooks")
		return handler.CheckSoldRecord(c, db)
	})
}