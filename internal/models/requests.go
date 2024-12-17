package models

type (
	CreateEventRequest struct {
	}

	TakeTicketRequest struct {
		UserId  int `json:"user_id"`
		EventId int `json:"event_id"`
	}

	QrRequestData struct {
		URL        string `json:"url"`
		ObjectName string `json:"object_name"`
	}
)
