package clickhouse

import (
	"ticket-api/internal/models"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitClickhouse() {
	dsn := "tcp://localhost:9000"
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.TicketMeta{},
	)

	DB = db
}
