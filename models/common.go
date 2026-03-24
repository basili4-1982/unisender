package models

import (
	"encoding/json"
	"time"
)

// ListOptions опции для пагинации
type ListOptions struct {
	Limit  int `json:"limit,omitempty" url:"limit,omitempty"`
	Offset int `json:"offset,omitempty" url:"offset,omitempty"`
}

// PaginatedResponse ответ с пагинацией
type PaginatedResponse struct {
	Items      json.RawMessage `json:"items"`
	TotalCount int             `json:"total_count"`
	Limit      int             `json:"limit"`
	Offset     int             `json:"offset"`
}

// ErrorResponse структура ошибки API
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// Metadata метаданные
type Metadata struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// IDResponse ответ с ID
type IDResponse struct {
	ID string `json:"id"`
}
