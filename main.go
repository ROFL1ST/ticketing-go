package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"ticketing-backend/config"
	"ticketing-backend/models"
	"ticketing-backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system env variables")
	}
	app := fiber.New()
	config.Connect()

	config.DB.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Comment{})

	routes.RegisterRoutes(app)

	app.Listen(":" + os.Getenv("APP_PORT"))
}
