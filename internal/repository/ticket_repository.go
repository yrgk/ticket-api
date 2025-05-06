package repository

import (
	"encoding/json"
	"ticket-api/internal/models"
	"ticket-api/internal/utils"
	"ticket-api/pkg/postgres"
	"time"
)

func TakeTicket(body models.Ticket) error {
	if err := postgres.DB.Create(&body).Error; err != nil {
		return err
	}

	if err := postgres.DB.Exec("UPDATE forms SET participants_count = participants_count + 1 WHERE id = ?", body.FormId).Error; err != nil {
		return err
	}

	return nil
}

func GetTicket(id string, userId int) models.TicketResponse {
	var ticket models.TicketResponse
	postgres.DB.Raw("SELECT f.title, t.qr_code_url, t.form_id, t.variety, t.is_activated FROM forms f, tickets t WHERE t.ticket_id = ? AND t.user_id = ?", id, userId).Scan(&ticket)

	return ticket
}

func GetTicketForChecking(id string, userId int) models.TicketCheckResponse {
	var ticket models.TicketCheckResponse
	postgres.DB.Raw("SELECT variety, is_activated, user_id FROM tickets WHERE ticket_id = ?", id).Scan(&ticket)

	return ticket
}

func CheckTicket(ticketId string, validatorId int) models.TicketCheckResponse {
	var ticket models.Ticket
	postgres.DB.Raw("SELECT * FROM tickets WHERE ticket_id = ?", ticketId).Scan(&ticket)

	var title string
	postgres.DB.Raw("SELECR title FROM forms WHERE form_id = ?", ticket.FormId).Scan(&title)

	response := models.TicketCheckResponse{
		Title:       title,
		IsActivated: ticket.IsActivated,
	}

	return response
	// var validatorIDs []int
	// postgres.DB.Raw("SELECT validator_id FROM validators WHERE event_id = ?", ticket.FormId).Scan(&validatorIDs)

	// if validatorId == event.OrganizatorId {

	// 	ticketData := GetTicketForChecking(ticketId, validatorId)
	// 	return ticketData
	// }

	// for _, id := range validatorIDs {
	// 	if validatorId == id {

	// 		ticketData := GetTicketForChecking(ticketId, validatorId)
	// 		return ticketData
	// 	}
	// }

	// return models.TicketCheckResponse{}
}

func ValidateTicket(ticketId, verifierId string) error {
	// Changing status of Ticket
	postgres.DB.Raw("UPDATE tickets SET is_activated = TRUE WHERE ticket_id = ?", ticketId)

	// Updating data in clickhouse

	// Deleting qr code from s3
	if err := utils.DeleteFromS3(ticketId); err != nil {
		return err
	}

	return nil
}

func GetMyTickets(id int) []models.MyTicketResponse {
	var tickets []models.MyTicketResponse
	postgres.DB.Raw("SELECT t.variety, t.is_activated, e.title, e.base_price, e.cover_url FROM tickets t JOIN events e ON t.event_id = e.id WHERE t.user_id = ?", id).Scan(&tickets)

	return tickets
}

func CheckValidator(eventId, validatorId int) bool {
	var validator models.Validator
	postgres.DB.Raw("SELECT * FROM validators WHERE event_id = ? AND validator_id = ?", eventId, validatorId).Scan(&validator)

	return validator.EventId != 0
}

func UploadUserData(body models.Ticket, ticketBody models.TakeTicketRequest) error {
	formData, err := json.Marshal(ticketBody.FormData)
	if err != nil {
		return err
	}

	data := models.TicketMeta{
		UserId:     body.UserId,
		FormId:     body.FormId,
		TicketId:   body.TicketId,
		Variety:    body.Variety,
		TimeBought: time.Now(),
		UserData:   formData,
	}

	if err := postgres.DB.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
