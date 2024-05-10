package internal

import (
	UH "main/internal/users/handler"
	UU "main/internal/users/usecase"
	UR "main/internal/users/repository"

	BH "main/internal/books/handler"
	BR "main/internal/books/repository"
	BU "main/internal/books/usecase"
	"main/internal/validating"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes function is used to set up the routes for the application
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	userHandler := UH.NewUserHandler(UU.NewUserUsecase(UR.NewUserRepo(db)))
	bookHandler := BH.NewBookHandler(BU.NewBookUsecase(BR.NewBookRepo(db)))

	app.Post("/login", userHandler.Login)
	app.Post("/register", userHandler.AddUser)

	customer := app.Group("/customer")
	{
		customer.Use(validating.IsCustomer)
	}

	owner := app.Group("/owner")
	{
		owner.Use(validating.IsOwner)
		owner.Post("/addbook", bookHandler.AddBook)
	}

}