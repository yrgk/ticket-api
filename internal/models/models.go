package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type (
	Ticket struct {
		gorm.Model
		QrCodeUrl    string          `json:"qr_code_url"`
		CoverUrl     string          `json:"cover_url"`
		TicketNumber int             `json:"ticket_number"`
		UserId       int             `json:"user_id"`
		FormId       int             `json:"form_id"`
		TicketId     string          `json:"ticket_id"`
		VarietyId    int             `json:"variety_id"`
		IsActivated  bool            `json:"is_activated"`
		UserData     json.RawMessage `json:"user_data" gorm:"type:jsonb"`
	}

	// Clickhouse model (not at the moment)
	TicketMeta struct {
		gorm.Model
		// UserId        int             `json:"user_id"`
		// FormId        int             `json:"form_id"`
		// TicketId      string          `json:"ticket_id"`
		// IsActivated   bool            `json:"is_activated"`
		// TimeBought    time.Time       `json:"time_bought"`
		// TimeActivated time.Time       `json:"time_activated"`
		UserData json.RawMessage `json:"user_data" gorm:"type:jsonb"`
	}

	Validator struct {
		ValidatorId int  `json:"validator_id" validate:"required"`
		EventId     uint `json:"event_id" validate:"required"`
	}

	Variety struct {
		ID       uint   `json:"id" gorm:"primarykey"`
		FormId   int    `json:"form_id"`
		Title    string `json:"title"`
		CoverUrl string `json:"cover_url"`
		Price    int    `json:"price"`
	}

	Layout struct {
		gorm.Model
		Title  string
		Type   string
		FormId int
		Schema json.RawMessage `gorm:"type:jsonb"`
		Zones  json.RawMessage `gorm:"type:jsonb"`
	}
)
