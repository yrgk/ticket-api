package models

type (
	TakeTicketRequest struct {
		UserId    int         `json:"user_id"`
		FormId    int         `json:"form_id"`
		VarietyId int         `json:"variety_id"`
		FormData  interface{} `json:"form_data"`
	}

	QrRequestData struct {
		URL        string `json:"url"`
		ObjectName string `json:"object_name"`
	}
)
