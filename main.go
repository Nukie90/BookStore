package main

import (
	"flag"
	"fmt"
	"log"
	"main/api"
	"main/infrastructure"
	"os"

	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func main() {
	envFlag := flag.String("env", "common", "a string")

	flag.Parse()

	configDetail, err := infrastructure.LoadConfig(*envFlag)
	if err != nil {
		log.Fatal(err)
	}

	dbConfig := infrastructure.NewGormConfig(configDetail)
	db, err := dbConfig.Connection()
	if err != nil {
		log.Fatal(err)
	}
	dbConfig.AutoMigrate(db)

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