package api

import (
    "main/api/handler"
	"main/api/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    }) 
    //user
    routes.UserRoutes(app, db)
    //owner
    routes.OwnerRoutes(app, db)
    //customer
    routes.CustomerRoutes(app, db)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Get("/accessible", func(c *fiber.Ctx) error {
        return handler.AccessibleArea(c)
    })

    app.Get("/restictions", func(c *fiber.Ctx) error {
        return handler.RestrictedArea(c)
    })

    app.Post("/login", func(c *fiber.Ctx) error {
        return handler.Login(c, db)
    })

    app.Get("/logout", func(c *fiber.Ctx) error {
        return handler.Logout(c)
    })

    app.Get("/browse/:how", func(c *fiber.Ctx) error {
        return handler.BrowseBooks(c, db)
    })
}