package models

type (
	TakeTicketRequest struct {
		UserId   int         `json:"user_id"`
		FormId   string      `json:"form_id"`
		Variety  string      `json:"variety"`
		FormData interface{} `json:"form_data"`
	}

	QrRequestData struct {
		URL        string `json:"url"`
		ObjectName string `json:"object_name"`
	}
)
