package repository

import (
	"encoding/json"
	"errors"
	"log"
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
	postgres.DB.Raw("SELECT f.title, t.qr_code_url, t.form_id, t.variety, t.is_activated FROM forms f, tickets t WHERE t.ticket_id = ? AND t.user_id = ? AND f.id = t.form_id;", id, userId).Scan(&ticket)

	return ticket
}

func GetTicketForChecking(id string, userId int) models.TicketCheckResponse {
	var ticket models.TicketCheckResponse
	postgres.DB.Raw("SELECT variety, is_activated, user_id FROM tickets WHERE ticket_id = ?", id).Scan(&ticket)

	return ticket
}

func CheckTicket(ticketId string, validatorId int) (models.TicketCheckResponse, error) {
	var ticket models.Ticket
	postgres.DB.Raw("SELECT form_id, is_activated, variety FROM tickets WHERE ticket_id = ?", ticketId).Scan(&ticket)

	var form models.Form
	postgres.DB.Raw("SELECT title, user_id FROM forms WHERE id = ?", ticket.FormId).Scan(&form)

	// Checking if validator is real
	if form.UserId != validatorId {
		return models.TicketCheckResponse{}, errors.New("forbidden")
	}

	response := models.TicketCheckResponse{
		Title:       form.Title,
		IsActivated: ticket.IsActivated,
		Variety:     ticket.Variety,
	}

	return response, nil
}

// ticketId is Public
func ValidateTicket(ticketId string, userId int) error {
	var ticket models.Ticket
	postgres.DB.Raw("SELECT form_id FROM tickets WHERE ticket_id = ?", ticketId).Scan(&ticket)

	var form models.Form
	postgres.DB.Raw("SELECT user_id FROM forms WHERE id = ?", ticket.FormId).Scan(&form)

	// Checking if validator id is valid
	if userId != form.UserId {
		return errors.New("incorrect validator id")
	}

	// Changing status of Ticket and deleting QR-code link
	postgres.DB.Exec("UPDATE tickets SET qr_code_url = NULL, is_activated = TRUE WHERE ticket_id = ?", ticketId)

	// Updating data in clickhouse

	// Deleting qr code from s3
	if err := utils.DeleteFromS3(ticketId); err != nil {
		log.Println("ERROR DELETING QR CODE FROM S3", ticketId, err)
	}

	return nil
}

func GetMyTickets(id int) []models.MyTicketResponse {
	var tickets []models.MyTicketResponse
	postgres.DB.Raw("SELECT t.variety, t.ticket_id, t.is_activated, f.title FROM tickets t JOIN forms f ON t.form_id = f.id WHERE t.user_id = ?", id).Scan(&tickets)

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
