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
)
