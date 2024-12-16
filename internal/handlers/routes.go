package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Setting up CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// Setting up logger
	app.Use(logger.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")


	// Create event

	// Get event
	v1.Get("/event/:id", GetEventHandler)

	// Sign up in event (take a ticket)

	// Get ticket

	// Verify ticket
}