package models

type (
	EventResponse struct {
		Title             string
		Description       string
		IsPaid            bool
		Price             int
		ParticipantsCount int
		Form              interface{}
		IsDoubleVerify    bool
	}

	TicketResponse struct {
		Title     string `json:"title"`
		QrCodeUrl string `json:"qr_code_url"`
	}
)
