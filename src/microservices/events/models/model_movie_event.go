package models

type MovieEvent struct {

	// Идентификатор фильма
	MovieId int32 `json:"movie_id"`

	// Название фильма
	Title string `json:"title"`

	// Действие с фильмом
	Action string `json:"action"`

	// Идентификатор пользователя (опционально)
	UserId int32 `json:"user_id,omitempty"`

	// Рейтинг (опционально)
	Rating float32 `json:"rating,omitempty"`

	// Жанры фильма (опционально)
	Genres []string `json:"genres,omitempty"`

	// Описание фильма (опционально)
	Description string `json:"description,omitempty"`
}
