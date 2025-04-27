package main

import (
	"ticket-api/config"
	"ticket-api/internal/handlers"
	"ticket-api/internal/models"
	// "ticket-api/pkg/clickhouse"
	"ticket-api/pkg/postgres"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main()  {
	config.GetConfig()

	postgres.ConnectDb()

	postgres.DB.AutoMigrate(
		models.Form{},
		models.Field{},
		models.Ticket{},
		models.Validator{},
	)

	// clickhouse.InitClickhouse()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	// Setting up CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// Setting up logger
	app.Use(logger.New())

	handlers.SetupRoutes(app)

	// app.Listen(":8080")
	app.ListenTLS(":8080", "/etc/letsencrypt/live/catalogio.space/fullchain.pem", "/etc/letsencrypt/live/catalogio.space/privkey.pem")
}