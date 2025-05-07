package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Create form
	v1.Post("/form", CreateFormHandler)

	// Get my forms
	v1.Get("/form/my", GetMyProjectsHandler)

	// Get form
	v1.Get("/form/:id", GetFormHandler)



	// Take a ticket
	v1.Post("/ticket", TakeTicketHandler)

	// Get my tickets
	v1.Get("/ticket/my", GetMyTicketsHandler)

	// Get ticket
	v1.Get("/ticket/:id", GetTicketHandler)

	// Check ticket from QR-code link (ticket_id = public_id)
	v1.Get("/ticket/check/:ticket_id", CheckTicketHandler)

	// Verify ticket
	v1.Post("/ticket/validate/:ticketId", ValidateTicketHandler)

	// Pay ticket
	v1.Post("/ticket/pay", PayTicketHandler)

	// Export registrations data to excel file
	v1.Get("/export/:id", ExportDataToExcelHandler)
}
