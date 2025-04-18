package models

import "time"

type (
	CreateEventRequest struct {
		Title         string                 `json:"title" validate:"required"`
		Description   string                 `json:"description" validate:"required"`
		IsPaid        bool                   `json:"is_paid"`
		CoverUrl      string                 `json:"cover_url" validate:"required"`
		PageData      map[string]interface{} `json:"page_data"`
		BasePrice     int                    `json:"base_price"`
		StartTime     time.Time              `json:"start_time"`
		Duration      string                 `json:"duration"`
		Capacity      int                    `json:"capacity" validate:"required"`
		OrganizatorId int                    `json:"organizator_id" validate:"required"`
		FormData      []Field                `json:"form_data" validate:"required"`
		// IsDoubleVerify bool      `json:"is_double_verify"`
	}

	TakeTicketRequest struct {
		UserId   int         `json:"user_id"`
		FormId   int         `json:"form_id"`
		Variety  string      `json:"variety"`
		FormData interface{} `json:"form_data"`
	}

	QrRequestData struct {
		URL        string `json:"url"`
		ObjectName string `json:"object_name"`
	}
)
