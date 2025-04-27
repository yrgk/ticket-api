package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type (
	Ticket struct {
		gorm.Model
		QrCodeUrl   string `json:"qr_code_url"`
		UserId      int    `json:"user_id"`
		FormId      int    `json:"form_id"`
		TicketId    string `json:"ticket_id"`
		Variety     string `json:"variety"`
		IsActivated bool   `json:"is_activated"`
	}

	// Clickhouse model (not at the moment)
	TicketMeta struct {
		gorm.Model
		UserId        int             `json:"user_id"`
		FormId        int             `json:"form_id"`
		TicketId      string          `json:"ticket_id"`
		Variety       string          `json:"variety"`
		IsActivated   bool            `json:"is_activated"`
		TimeBought    time.Time       `json:"time_bought"`
		TimeActivated time.Time       `json:"time_activated"`
		UserData      json.RawMessage `json:"user_data"`
	}

	Validator struct {
		ValidatorId int  `json:"validator_id" validate:"required"`
		EventId     uint `json:"event_id" validate:"required"`
	}
)
