package models

type (
	// EventResponse struct {
	// 	Event Event
	// 	Form  []FieldResponse
	// }

	TicketResponse struct {
		Title       string `json:"title"`
		QrCodeUrl   string `json:"qr_code_url"`
		EventId     int    `json:"event_id"`
		Variety     string `json:"variety"`
		IsActivated bool   `json:"is_activated"`
	}

	MyTicketResponse struct {
		Title       string `json:"title"`
		CoverUrl    string `json:"cover_url"`
		Variety     string `json:"variety"`
		IsActivated bool   `json:"is_activated"`
	}

	TicketCheckResponse struct {
		Title       string `json:"title"`
		Variety     string `json:"variety"`
		IsActivated bool   `json:"is_activated"`
		UserId      int    `json:"user_id"`
	}
)
