package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Create event
	// NOT MADE UP
	v1.Post("/event", CreateEventHandler)

	// Get event
	// NOT MADE UP
	v1.Get("/event/:id", GetEventHandler)



	// Sign up in event (take a ticket)
	v1.Post("/ticket", TakeTicketHandler)

	// Get ticket
	v1.Get("/ticket/:id", GetTicketHandler)

	// Get my tickets
	v1.Get("/ticket/my")

	// Check ticket
	// v1.Get("/")

	// Verify ticket
	v1.Post("/ticket/verify", VerifyTicketHandler)
}