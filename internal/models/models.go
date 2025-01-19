package models

import (
	"gorm.io/gorm"
)

type (
	Event struct {
		gorm.Model
		Title             string `json:"title"`
		Description       string `json:"description"`
		PageData          string `json:"page_data"`
		IsPaid            bool   `json:"is_paid"`
		CoverUrl          string `json:"cover_url"`
		BasePrice         int    `json:"base_price"`
		ParticipantsCount int    `gorm:"default:0" json:"participants_count"`
		Capacity          int    `gorm:"check: participants_count <= capacity" json:"capacity"`
		IsDoubleVerify    bool   `json:"is_double_verify"`
		OrganizatorId     int    `json:"organizator_id"`
	}

	Ticket struct {
		// gorm.Model
		QrCodeUrl   string `json:"qr_code_url"`
		UserId      int    `json:"user_id"`
		EventId     int    `json:"event_id"`
		TicketId    string `json:"ticket_id"`
		IsActivated bool   `json:"is_activated"`
	}

	Field struct {
		Name    string `bson:"name" json:"name" validate:"required"`
		Type    string `bson:"type" json:"type" validate:"required"`
		EventId uint   `bson:"type" json:"event_id" validate:"required"`
	}

	Checker struct {
		CheckerId int  `json:"checker_id" validate:"required"`
		EventId   uint `json:"event_id" validate:"required"`
	}
)
