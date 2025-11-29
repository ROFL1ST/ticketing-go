package controllers

import (
	"github.com/gofiber/fiber/v2"
	"ticketing-backend/config"
	"ticketing-backend/models"
)

func GetTicketsByUser(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	var tickets []models.Ticket
	config.DB.Where("user_id = ?", userID).Preload("Comments").Find(&tickets)
	return c.JSON(tickets)
}

func CreateTicket(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	var body models.Ticket

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	body.UserID = userID.(uint)
	body.Status = "pending"

	config.DB.Create(&body)
	return c.JSON(body)
}

func GetTicketById(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("userID")

	var ticket models.Ticket
	config.DB.Where("id = ? AND user_id = ?", id, userID).Preload("Comments").First(&ticket)

	return c.JSON(ticket)
}

func AddComment(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	var body models.Comment

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	body.UserID = userID.(uint)

	config.DB.Create(&body)
	return c.JSON(body)
}
