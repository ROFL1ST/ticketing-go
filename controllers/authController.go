package controllers

import (
	"ticketing-backend/config"
	"ticketing-backend/middlewares"
	"ticketing-backend/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	c.BodyParser(&body)

	var user models.User
	err := config.DB.Where("email = ?", body.Email).First(&user).Error
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid login"})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Wrong password"})
	}

	token, _ := middleware.GenerateToken(user.ID, user.Role)

	return c.JSON(fiber.Map{
		"user":  user,
		"token": token,
	})
}

func Register(c *fiber.Ctx) error {
	body := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}{}
	c.BodyParser(&body)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create user"})
	}
	return c.JSON(user)
}
