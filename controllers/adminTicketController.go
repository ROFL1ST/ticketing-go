package controllers

import (
	"github.com/gofiber/fiber/v2"
	"ticketing-backend/config"
	"ticketing-backend/models"
)

func AdminGetTickets(c *fiber.Ctx) error {
	var tickets []models.Ticket
	config.DB.Preload("Comments").Find(&tickets)
	return c.JSON(tickets)
}

func AdminGetTicketById(c *fiber.Ctx) error {
	id := c.Params("id")
	var ticket models.Ticket
	config.DB.Preload("Comments").First(&ticket, id)
	return c.JSON(ticket)
}

func UpdateTicketStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var body struct {
		Status string `json:"status"`
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	config.DB.Model(&models.Ticket{}).Where("id = ?", id).Update("status", body.Status)

	return c.JSON(fiber.Map{
		"message": "status updated",
	})
}
