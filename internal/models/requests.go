package models

type (
	CreateEventRequest struct {
		Title          string  `json:"title" validate:"required"`
		Description    string  `json:"description" validate:"required"`
		IsPaid         bool    `json:"is_paid"`
		CoverUrl       string  `json:"cover_url" validate:"required"`
		Capacity       int     `json:"capacity" validate:"required"`
		IsDoubleVerify bool    `json:"is_double_verify"`
		OrganizatorId  int     `json:"organizator_id" validate:"required"`
		FormData       []Field `json:"form_data" validate:"required"`
	}

	TakeTicketRequest struct {
		UserId   int         `json:"user_id"`
		EventId  int         `json:"event_id"`
		FormData interface{} `json:"form_data"`
	}

	QrRequestData struct {
		URL        string `json:"url"`
		ObjectName string `json:"object_name"`
	}
)
