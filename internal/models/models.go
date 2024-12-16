package models

import "gorm.io/gorm"

type (
	Event struct {
		gorm.Model
		Title       string
		Description string
		IsPaid      bool

	}

	Ticket struct {
		gorm.Model
		CodeUrl string
		UserId  int
		EventId int
		// UserData
	}
)
