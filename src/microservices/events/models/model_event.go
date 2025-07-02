package models

import (
	"time"
)

type Event struct {

	// Уникальный идентификатор события
	Id string `json:"id"`

	// Тип события
	Type string `json:"type"`

	// Время события
	Timestamp time.Time `json:"timestamp"`

	// Полезная нагрузка события (зависит от типа события)
	Payload map[string]interface{} `json:"payload"`
}
