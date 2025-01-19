package repository

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ticket-api/config"
	"ticket-api/internal/models"
	"ticket-api/pkg/postgres"
	"time"
)

func TakeTicket(body models.Ticket) error {
	if err := postgres.DB.Create(&body).Error; err != nil {
		return err
	}

	if err := postgres.DB.Exec("UPDATE events SET participants_count = participants_count + 1 WHERE id = ?", body.EventId).Error; err != nil {
		return err
	}

	return nil
}

func CreateQrCode(body models.TakeTicketRequest) ([]byte, error) {
	original := fmt.Sprintf("%d%d%s", body.UserId, body.EventId, time.Now())
	hash := md5.New()
	hash.Write([]byte(original))

	md5string := hex.EncodeToString(hash.Sum(nil))
	url := fmt.Sprintf("%s?startapp=check_%s", config.Config.WebappName, md5string)

	data := models.QrRequestData{
		URL:        url,
		ObjectName: md5string,
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
	postgres.DB.Raw("SELECT e.title, t.qr_code_url FROM events e, tickets t WHERE e.id = ? AND t.user_id = ?", id, userId).Scan(&ticket)

	return ticket
}

func VerifyTicket(ticketId, verifierId string) error {
	// Checking if data is valid
	// Changing status of Ticket
	postgres.DB.Raw("UPDATE is_activated FROM tickets WHERE ticket_id = ? AND ")
	// Deleting qr code from s3
	return nil
}

func GetMyTickets(id int) []models.Ticket {
	var tickets []models.Ticket
	postgres.DB.Raw("SELECT * FROM tickets WHERE user_id = ?", id).Scan(&tickets)

	return tickets
}