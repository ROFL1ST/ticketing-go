package middleware

import (
	"os"
	"strings"
	"ticketing-backend/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":   userID,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}

func Protected(c *fiber.Ctx) error {
	secret := os.Getenv("JWT_SECRET")
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return utils.Error(c, fiber.StatusUnauthorized, "Missing or malformed JWT")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return utils.Error(c, fiber.StatusUnauthorized, "Missing or malformed JWT")
	}

	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid or expired JWT")
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("userID", uint(claims["id"].(float64)))
	c.Locals("role", claims["role"].(string))

	return c.Next()
}


func OnlyAdmin(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role != "admin" {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Next()
}

func OnlyUser(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role != "user" {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Next()
}
