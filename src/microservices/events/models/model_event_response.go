package models

type EventResponse struct {

	// Статус операции
	Status string `json:"status"`

	// Партиция Kafka
	Partition int32 `json:"partition"`

	// Смещение в партиции Kafka
	Offset int32 `json:"offset"`

	Event Event `json:"event"`
}
