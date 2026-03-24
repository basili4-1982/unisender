package models

import "time"

// Suppression запись стоп-листа
type Suppression struct {
	// Email
	Email string `json:"email"`

	// Причина
	Reason string `json:"reason"`

	// Тип: 'email', 'domain'
	Type string `json:"type"`

	// Дата добавления
	CreatedAt time.Time `json:"created_at"`
}

// SetSuppressionRequest запрос на добавление в стоп-лист
type SetSuppressionRequest struct {
	// Список email или доменов
	Items []SuppressionItem `json:"items"`
}

// SuppressionItem элемент для добавления
type SuppressionItem struct {
	// Email или домен
	Value string `json:"value"`

	// Тип: 'email', 'domain'
	Type string `json:"type"`

	// Причина
	Reason string `json:"reason,omitempty"`
}

// SetSuppressionResponse ответ на добавление
type SetSuppressionResponse struct {
	Result SetSuppressionResult `json:"result"`
}

// SetSuppressionResult результат добавления
type SetSuppressionResult struct {
	// Добавлено записей
	Added int `json:"added"`

	// Пропущено записей
	Skipped int `json:"skipped"`
}

// SuppressionListResponse ответ со списком стоп-листа
type SuppressionListResponse struct {
	Result []Suppression `json:"result"`
}
