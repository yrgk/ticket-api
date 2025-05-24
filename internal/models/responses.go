package models

import "encoding/json"

type (
	TicketResponse struct {
		ID                uint    `json:"id"`
		Title             string  `json:"title"`
		QrCodeUrl         string  `json:"qr_code_url"`
		CoverUrl          string  `json:"cover_url"`
		TicketNumber      int     `json:"ticket_number"`
		ParticipantsLimit int     `json:"participants_limit"`
		FormId            int     `json:"form_id"`
		Variety           Variety `json:"variety"`
		IsActivated       bool    `json:"is_activated"`
	}

	MyTicketResponse struct {
		Title             string `json:"title"`
		CoverUrl          string `json:"cover_url"`
		TicketId          string `json:"ticket_id"`
		TicketNumber      int    `json:"ticket_number"`
		ParticipantsLimit int    `json:"participants_limit"`
		IsActivated       bool   `json:"is_activated"`
	}

	MyProjectsResponse struct {
		Forms []struct {
			PublicId          string `json:"public_id"`
			Title             string `json:"title"`
			ParticipantsCount int    `json:"participants_count"`
			ParticipantsLimit int    `json:"participants_limit"`
		}
	}

	TicketCheckResponse struct {
		Title       string `json:"title"`
		IsActivated bool   `json:"is_activated"`
		VarietyId   int    `json:"variety_id"`
		// UserId      int  `json:"user_id"`
	}

	LayoutResponse struct {
		Title  string          `json:"title"`
		Type   string          `json:"type"`
		Schema json.RawMessage `json:"schema" gorm:"type:jsonb"`
		Zones  json.RawMessage `json:"zones" gorm:"type:jsonb"`
	}
)
