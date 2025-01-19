package models

type (
	EventResponse struct {
		Event Event
		Form  []Field
	}

	TicketResponse struct {
		Title     string `json:"title"`
		QrCodeUrl string `json:"qr_code_url"`
	}
)
