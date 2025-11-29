package controllers

import (
	"github.com/gofiber/fiber/v2"
	"ticketing-backend/config"
	"ticketing-backend/models"
	"ticketing-backend/utils"
	"strconv"
)

func GetTicketsByUser(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	var tickets []models.Ticket

	if err := config.DB.Where("user_id = ?", userID).
		Preload("Comments").
		Find(&tickets).Error; err != nil {
		return utils.Error(c, 500, "Failed to fetch tickets")
	}

	return utils.Success(c, "Tickets fetched", tickets)
}

func CreateTicket(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	var body models.Ticket

	if err := c.BodyParser(&body); err != nil {
		return utils.Error(c, 400, "Invalid request body")
	}

	body.UserID = userID.(uint)
	body.Status = "pending"

	if err := config.DB.Create(&body).Error; err != nil {
		return utils.Error(c, 500, "Failed to create ticket")
	}

	return utils.Created(c, "Ticket created", body)
}

func GetTicketById(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("userID")

	var ticket models.Ticket

	if err := config.DB.
		Where("id = ? AND user_id = ?", id, userID).
		Preload("Comments").
		First(&ticket).Error; err != nil {
		return utils.Error(c, 404, "Ticket not found")
	}

	return utils.Success(c, "Ticket fetched", ticket)
}

func AddComment(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	ticketID := c.Params("id")

	var body struct {
		Message string `json:"message"`
	}

	if err := c.BodyParser(&body); err != nil {
		return utils.Error(c, 400, "Invalid request body")
	}

	comment := models.Comment{
		TicketID: parseID(ticketID),
		UserID:   userID.(uint),
		Message:  body.Message,
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		return utils.Error(c, 500, "Failed to add comment")
	}

	return utils.Created(c, "Comment added", comment)
}

func parseID(id string) uint {
	val, _ := strconv.Atoi(id)
	return uint(val)
}
