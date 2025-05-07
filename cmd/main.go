package main

import (
	"log"

	"ticket-api/config"
	"ticket-api/internal/handlers"
	// "ticket-api/internal/models"
	// "ticket-api/pkg/clickhouse"
	"ticket-api/pkg/postgres"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main()  {
	log.Println("Config gotten")
	config.GetConfig()

	log.Println("DB connected")
	postgres.ConnectDb()

	log.Println("Tables migrated")
	// postgres.DB.AutoMigrate(
	// 	models.Form{},
	// 	models.Field{},
	// 	models.Ticket{},
	// 	models.TicketMeta{},
	// 	// models.Validator{},
	// )

	// clickhouse.InitClickhouse()

	log.Println("App inited")
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	// Setting up CORS
	log.Println("CORS setted up")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// Setting up logger
	app.Use(logger.New())

	// Setting up routes
	log.Println("Routes setted up")
	handlers.SetupRoutes(app)

	app.Listen(":3001")
}