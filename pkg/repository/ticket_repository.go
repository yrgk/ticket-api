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
