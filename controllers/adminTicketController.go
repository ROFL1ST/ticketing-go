package controllers

import (
	"github.com/gofiber/fiber/v2"
	"ticketing-backend/config"
	"ticketing-backend/models"
	"ticketing-backend/utils"
)

func AdminGetTickets(c *fiber.Ctx) error {
	var tickets []models.Ticket

	if err := config.DB.Preload("Comments").Find(&tickets).Error; err != nil {
		return utils.Error(c, 500, "Failed to fetch tickets")
	}

	return utils.Success(c, "Tickets fetched", tickets)
}

func AdminGetTicketById(c *fiber.Ctx) error {
	id := c.Params("id")
	var ticket models.Ticket

	if err := config.DB.Preload("Comments").First(&ticket, id).Error; err != nil {
		return utils.Error(c, 404, "Ticket not found")
	}

	return utils.Success(c, "Ticket fetched", ticket)
}

func UpdateTicketStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	var body struct {
		Status string `json:"status"`
	}

	if err := c.BodyParser(&body); err != nil {
		return utils.Error(c, 400, "Invalid request body")
	}

	if err := config.DB.Model(&models.Ticket{}).Where("id = ?", id).Update("status", body.Status).Error; err != nil {
		return utils.Error(c, 500, "Failed to update status")
	}

	return utils.Success(c, "Status updated", fiber.Map{
		"id":     id,
		"status": body.Status,
	})
}
