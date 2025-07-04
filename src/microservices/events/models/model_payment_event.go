package models

import (
	"time"
)

type PaymentEvent struct {

	// Идентификатор платежа
	PaymentId int32 `json:"payment_id"`

	// Идентификатор пользователя
	UserId int32 `json:"user_id"`

	// Сумма платежа
	Amount float32 `json:"amount"`

	// Статус платежа
	Status string `json:"status"`

	// Время платежа
	Timestamp time.Time `json:"timestamp"`

	// Тип метода оплаты (опционально)
	MethodType string `json:"method_type,omitempty"`
}
