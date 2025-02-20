package clickhouse

import (
	"ticket-api/internal/models"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitClickhouse() {
	// var (
	// 	ctx       = context.Background()
	// 	conn, err = clickhouse.Open(&clickhouse.Options{
	// 		Addr: []string{"localhost:9000"},
	// 		Auth: clickhouse.Auth{
	// 			Database: "default",
	// 			Username: "default",
	// 			Password: "",
	// 		},
	// 		// TLS: &tls.Config{
	// 		// 	InsecureSkipVerify: true,
	// 		// },
	// 	})
	// )

	// if err != nil {
	// 	panic(err)
	// }

	// if err := conn.Ping(ctx); err != nil {
	// 	if exception, ok := err.(*clickhouse.Exception); ok {
	// 		fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
	// 	}
	// 	panic(err)
	// }

	// err = conn.Exec(context.Background(), `
		// CREATE TABLE IF NOT EXISTS ticket_meta (
		// 	user_id        Int32,
		// 	event_id       Int32,
		// 	ticket_id      String,
		// 	variety        String,
		// 	is_activated   Bool,
		// 	time_bought    DateTime,
		// 	time_activated DateTime,
		// 	user_data      String
		// ) ENGINE = MergeTree()
		// ORDER BY (user_id, event_id, ticket_id);
	// `)
	// if err != nil {
	// 	panic(err)
	// }

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
