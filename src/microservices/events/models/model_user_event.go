package models

import (
	"time"
)

type UserEvent struct {

	// Идентификатор пользователя
	UserId int32 `json:"user_id"`

	// Имя пользователя (опционально)
	Username string `json:"username,omitempty"`

	// Email пользователя (опционально)
	Email string `json:"email,omitempty"`

	// Действие пользователя
	Action string `json:"action"`

	// Время события
	Timestamp time.Time `json:"timestamp"`
}
