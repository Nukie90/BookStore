package main

import (
	"fmt"
	"log"
	"main/api"
	"main/data"
	"main/data/model"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := data.DbConnection()
	if err != nil {
		log.Fatal(err)
	}
	
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.ShoppingCart{})

	db.AutoMigrate(&model.Book{})

	db.AutoMigrate(&model.SoldRecord{})

	app := fiber.New()

	api.SetupRoutes(app, db)
	api.SetupMiddleware(app)

	portNum := os.Getenv("PORT")
	if portNum == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Println("Server is running on port: " + portNum)
	app.Listen(":" + portNum)
}