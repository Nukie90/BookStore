package routes

import (
	"fmt"
	"main/api/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CustomerRoutes(app *fiber.App, db *gorm.DB) {
	customer := fiber.New()

	app.Mount("/customer", customer)

	customer.Post("/addtocart", func(c *fiber.Ctx) error {
		fmt.Println("addtocart")
		return handler.AddToCart(c, db)
	})

	customer.Get("/getcart", func(c *fiber.Ctx) error {
		fmt.Println("getcart")
		return handler.GetCart(c, db)
	})

	customer.Get("/checkout", func(c *fiber.Ctx) error {
		fmt.Println("checkout")
		return handler.Checkout(c, db)
	})

	customer.Delete("/remove/:title", func(c *fiber.Ctx) error {
		fmt.Println("removeItem")
		return handler.RemoveItem(c, db)
	})
}