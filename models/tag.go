package models

// Tag структура тега
type Tag struct {
	// ID тега
	ID string `json:"id"`

	// Название тега
	Name string `json:"name"`

	// Количество писем с этим тегом
	Count int `json:"count"`

	// Метаданные
	Metadata Metadata `json:"metadata"`
}

// TagsListResponse ответ со списком тегов
type TagsListResponse struct {
	Result []Tag `json:"result"`
}
