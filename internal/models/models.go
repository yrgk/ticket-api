package models

import "gorm.io/gorm"

type (
	Event struct {
		gorm.Model
		Title             string `json:"title"`
		Description       string `json:"description"`
		IsPaid            bool   `json:"is_paid"`
		CoverUrl          string `json:"cover_url"`
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
		IsActivated bool   `json:"is_activated"`
	}

	Field struct {
		Name string `bson:"name" json:"name" validate:"required"`
		Type string `bson:"type" json:"type" validate:"required"`
	}
)
