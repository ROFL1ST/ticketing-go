package routes

import (
	"github.com/gofiber/fiber/v2"
	"ticketing-backend/controllers"
	"ticketing-backend/middlewares"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Public (register, login)
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	// User routes
	user := api.Group("/user", middleware.Protected, middleware.OnlyUser)
	user.Get("/tickets", controllers.GetTicketsByUser)
	user.Post("/tickets", controllers.CreateTicket)
	user.Get("/tickets/:id", controllers.GetTicketById)
	user.Post("/tickets/:id/comment", controllers.AddComment)
	user.Delete("/tickets/comment/:id", controllers.DeleteComment)

	// Admin routes
	admin := api.Group("/admin", middleware.Protected, middleware.OnlyAdmin)
	admin.Get("/tickets", controllers.AdminGetTickets)
	admin.Get("/tickets/:id", controllers.AdminGetTicketById)
	admin.Put("/tickets/:id/status", controllers.UpdateTicketStatus)
}
