package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ticket-api/config"
	"ticket-api/internal/models"
	"ticket-api/pkg/clickhouse"
	"ticket-api/pkg/postgres"
	"time"
)

func TakeTicket(body models.Ticket) error {
	if err := postgres.DB.Create(&body).Error; err != nil {
		return err
	}

	if err := postgres.DB.Exec("UPDATE events SET participants_count = participants_count + 1 WHERE id = ?", body.FormId).Error; err != nil {
		return err
	}

	return nil
}

func CreateQrCode(body models.TakeTicketRequest, ticketId string, objectName string) ([]byte, error) {
	url := fmt.Sprintf("%s?startapp=%s", config.Config.WebappName, ticketId)

	data := models.QrRequestData{
		URL:        url,
		ObjectName: objectName,
	}

	// Кодирование данных в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Создание HTTP-клиента с таймаутом
	client := &http.Client{Timeout: 10 * time.Second}

	// Создание запроса
	// req, err := http.NewRequest("POST", "http://185.197.75.220:8000/create", bytes.NewBuffer(jsonData))
	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/create", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Отправка запроса
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Чтение ответа
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return response, nil
}

func GetTicket(id int, userId int) models.TicketResponse {
	var ticket models.TicketResponse
	postgres.DB.Raw("SELECT e.title, t.qr_code_url, t.event_id, t.variety, t.is_activated FROM events e, tickets t WHERE e.id = ? AND t.user_id = ?", id, userId).Scan(&ticket)

	return ticket
}

func GetTicketForChecking(id string, userId int) models.TicketCheckResponse {
	var ticket models.TicketCheckResponse
	postgres.DB.Raw("SELECT variety, is_activated, user_id FROM tickets WHERE ticket_id = ?", id).Scan(&ticket)

	return ticket
}

func CheckTicket(ticketId string, validatorId int) models.TicketCheckResponse {
	// var ticket models.Ticket
	// postgres.DB.Raw("SELECT * FROM tickets WHERE ticket_id = ?", ticketId).Scan(&ticket)

	// event := GetEvent(ticket.FormId)

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

	return models.TicketCheckResponse{}
}

func ValidateTicket(ticketId, verifierId string) error {
	// Changing status of Ticket
	postgres.DB.Raw("UPDATE tickets SET is_activated = TRUE WHERE ticket_id = ?", ticketId)

	// Updating data in clickhouse

	// Deleting qr code from s3

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

	if err := clickhouse.DB.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
