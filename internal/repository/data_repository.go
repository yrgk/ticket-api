package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ticket-api/pkg/postgres"

	"github.com/xuri/excelize/v2"
)

func GetUserdata(userId int, formId string) []json.RawMessage {
	var userData []json.RawMessage
	postgres.DB.Raw("SELECT user_data FROM ticket_meta WHERE form_id = ? AND user_id = ?", formId, userId).Scan(&userData)

	return userData
}

func ParseToExcel(userData []json.RawMessage) (*bytes.Buffer, error) {
	// Creating excel file
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	// Parsing first resord for headers
	var first map[string]interface{}
	if err := json.Unmarshal(userData[0], &first); err != nil {
		return nil, fmt.Errorf("invalid JSON format: %w", err)
	}

	// Writing headers
	headers := make([]string, 0, len(first))
	col := 1
	for k := range first {
		cell, _ := excelize.CoordinatesToCellName(col, 1)
		f.SetCellValue(sheet, cell, k)
		headers = append(headers, k)
		col++
	}

	// Writing strings
	for i, raw := range userData {
		var item map[string]interface{}
		if err := json.Unmarshal(raw, &item); err != nil {
			continue
		}
		for j, key := range headers {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+2)
			f.SetCellValue(sheet, cell, item[key])
		}
	}

	// Setting up auto width columns
	for i := range headers {
		col, _ := excelize.ColumnNumberToName(i + 1)
		_ = f.SetColWidth(sheet, col, col, 20)
	}

	// Making buffer
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, fmt.Errorf("failed to write to buffer: %w", err)
	}

	return &buf, nil
}
