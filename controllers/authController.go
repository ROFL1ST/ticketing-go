package controllers

import (
	"ticketing-backend/config"
	"ticketing-backend/middlewares"
	"ticketing-backend/models"
	"ticketing-backend/utils"

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
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid email or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid email or password")
	}

	token, _ := middleware.GenerateToken(user.ID, user.Role)

	data := fiber.Map{
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	}

	return utils.Success(c, "Login successful", fiber.Map{
		"user":  data,
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
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to create user")
	}

	data := fiber.Map{
		"name":  user.Name,
		"email": user.Email,
		"role":  "user",
	}

	return utils.Created(c, "User created successfully", data)
}
