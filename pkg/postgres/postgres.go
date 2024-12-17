package postgres

import (
	"ticket-api/config"
	// "ticket-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := config.Config.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		// models.Event{},
		// models.Ticket{},
	)

	DB = db
}
