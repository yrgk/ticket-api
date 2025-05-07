package bot

import (
	"fmt"
	"net/http"
	"ticket-api/config"
	"ticket-api/internal/models"
	"ticket-api/pkg/postgres"
)

func SendTicketInChat(userId, formId int, ticketId string) error {
	var form models.Form
	postgres.DB.Raw("SELECT title FROM forms WHERE id = ?", formId).Scan(&form)

	text := fmt.Sprintf("✅ Регистрация прошла успешно! %s", form.Title)

	url := fmt.Sprintf("%s/ticket/%s?user_id=%d", config.Config.ClientUrl, ticketId, userId)

	req := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s&reply_markup={\"inline_keyboard\":[[{\"text\":\"Открыть\",\"web_app\":{\"url\":\"%s\"}}]]}", config.Config.BotToken, userId, text, url)

	_, err := http.Get(req)
	if err != nil {
		return err
	}

	return nil
}
