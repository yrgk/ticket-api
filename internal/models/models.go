package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type JSONB map[string]interface{}

type (
	Event struct {
		gorm.Model
		Title             string          `json:"title"`
		Description       string          `json:"description"`
		PageData          json.RawMessage `json:"page_data" gorm:"type:jsonb"`
		IsPaid            bool            `json:"is_paid"`
		CoverUrl          string          `json:"cover_url"`
		BasePrice         int             `json:"base_price"`
		StartTime         time.Time       `json:"start_time"`
		Duration          time.Duration   `json:"duration"`
		ParticipantsCount int             `gorm:"default:0" json:"participants_count"`
		Capacity          int             `gorm:"check: participants_count <= capacity" json:"capacity"`
		OrganizatorId     int             `json:"organizator_id"`
	}

	Ticket struct {
		gorm.Model
		QrCodeUrl   string `json:"qr_code_url"`
		UserId      int    `json:"user_id"`
		EventId     int    `json:"event_id"`
		TicketId    string `json:"ticket_id"`
		Variety     string `json:"variety"`
		IsActivated bool   `json:"is_activated"`
	}

	TicketMeta struct {
		UserId        int             `json:"user_id"`
		EventId       int             `json:"event_id"`
		TicketId      string          `json:"ticket_id"`
		Variety       string          `json:"variety"`
		IsActivated   bool            `json:"is_activated"`
		TimeBought    time.Time       `json:"time_bought"`
		TimeActivated time.Time       `json:"time_activated"`
		UserData      json.RawMessage `json:"user_data"`
	}

	Field struct {
		Name    string `bson:"name" json:"name" validate:"required"`
		Type    string `bson:"type" json:"type" validate:"required"`
		EventId uint   `bson:"type" json:"event_id" validate:"required"`
	}

	Validator struct {
		ValidatorId int  `json:"validator_id" validate:"required"`
		EventId     uint `json:"event_id" validate:"required"`
	}
)
