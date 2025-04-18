package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")


	// Create form
	v1.Get("/form", CreateFormHandler)

	// Get form
	v1.Get("/form/:id", GetFormHandler)

	// Take a ticket
	v1.Post("/ticket", TakeTicketHandler)

	// Get my tickets
	v1.Get("/ticket/my", GetMyTicketsHandler)

	// Get ticket
	v1.Get("/ticket/:id", GetTicketHandler)

	// Check ticket from QR-code link
	v1.Get("/ticket/check/:ticket_id", CheckTicketHandler)

	// Verify ticket
	v1.Post("/ticket/verify", ValidateTicketHandler)

	// Pay ticket
	v1.Post("/ticket/pay", PayTicketHandler)
}
