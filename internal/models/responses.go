package models

type (
	TicketResponse struct {
		Title     string `json:"title"`
		QrCodeUrl string `json:"qr_code_url"`
		FormId    int    `json:"form_id"`
		Variety     string `json:"variety"`
		IsActivated bool `json:"is_activated"`
	}

	MyTicketResponse struct {
		Title       string `json:"title"`
		CoverUrl    string `json:"cover_url"`
		Variety     string `json:"variety"`
		IsActivated bool   `json:"is_activated"`
	}

	MyProjectsResponse struct {
		Forms []struct {
			PublicId string
			Title    string
		}
	}

	TicketCheckResponse struct {
		Title string `json:"title"`
		IsActivated bool `json:"is_activated"`
		// Variety string `json:"variety"`
		// UserId      int  `json:"user_id"`
	}
)
