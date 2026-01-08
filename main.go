package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"ticketing-backend/config"
	"ticketing-backend/models"
	"ticketing-backend/routes"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system env variables")
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",

		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	config.Connect()	

	config.DB.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Comment{})
	config.SeedAdmin()

	routes.RegisterRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
