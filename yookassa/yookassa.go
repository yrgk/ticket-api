package yookassa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"ticket-api/config"
)

// PaymentRequest структура запроса к ЮKassa
type PaymentRequest struct {
	Amount struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"amount"`
	Capture      bool   `json:"capture"`
	Description  string `json:"description"`
	Recipient    string `json:"recipient"`
	Confirmation struct {
		Type      string `json:"type"`
		ReturnURL string `json:"return_url"`
	} `json:"confirmation"`
}

// PaymentResponse структура ответа от ЮKassa
type PaymentResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// CreatePayment функция для создания платежа
func CreatePayment(totalAmount float64, authorID string) (*PaymentResponse, error) {
	// Распределяем средства
	platformShare := totalAmount * 0.1
	authorShare := totalAmount * 0.9

	// Формируем тело запроса
	requestBody := &PaymentRequest{
		Capture:     true,
		Description: "Оплата за выпуск подкаста",
		Recipient:   authorID, // Укажите ID получателя
	}
	requestBody.Amount.Value = fmt.Sprintf("%.2f", platformShare+authorShare)
	requestBody.Amount.Currency = "RUB"
	requestBody.Confirmation.Type = "redirect"
	requestBody.Confirmation.ReturnURL = "https://your-return-url.com"

	// Кодируем запрос в JSON
	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Создаем HTTP-запрос
	req, err := http.NewRequest("POST", config.Config.YookassaAPIURL, bytes.NewBuffer(requestData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Добавляем заголовки
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(config.Config.YooKassaShopID, config.Config.YooKassaSecretKey)

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Обрабатываем ответ
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var paymentResp PaymentResponse
	err = json.NewDecoder(resp.Body).Decode(&paymentResp)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &paymentResp, nil
}

// func main() {
// 	totalAmount := 1000.0
// 	authorID := "author_account_id"

// 	resp, err := CreatePayment(totalAmount, authorID)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Printf("Payment created successfully! ID: %s, Status: %s\n", resp.ID, resp.Status)
// }
