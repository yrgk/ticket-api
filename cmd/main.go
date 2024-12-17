package main

import (
	"ticket-api/config"
	"ticket-api/internal/handlers"
	"ticket-api/pkg/postgres"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main()  {
	config.GetConfig()

	postgres.ConnectDb()

	app := fiber.New()

	// Setting up CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// Setting up logger
	app.Use(logger.New())

	handlers.SetupRoutes(app)

	app.Listen(":8080")
}