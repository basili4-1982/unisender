package models

// Webhook структура вебхука
type Webhook struct {
	// ID вебхука
	ID string `json:"id"`

	// Название вебхука
	Name string `json:"name"`

	// URL для вызова
	URL string `json:"url"`

	// События
	Events []WebhookEvent `json:"events"`

	// Активен ли
	IsActive bool `json:"is_active"`

	// Метаданные
	Metadata Metadata `json:"metadata"`
}

// WebhookEvent событие вебхука
type WebhookEvent struct {
	// Тип события
	Type string `json:"type"`

	// Настройки события
	Settings map[string]interface{} `json:"settings,omitempty"`
}

// SetWebhookRequest запрос на создание/обновление вебхука
type SetWebhookRequest struct {
	// ID вебхука (для обновления)
	ID string `json:"id,omitempty"`

	// Название вебхука
	Name string `json:"name"`

	// URL для вызова
	URL string `json:"url"`

	// События
	Events []WebhookEventInput `json:"events"`

	// Активен ли
	IsActive *bool `json:"is_active,omitempty"`
}

// WebhookEventInput входные данные события
type WebhookEventInput struct {
	// Тип события
	Type string `json:"type"`

	// Настройки события
	Settings map[string]interface{} `json:"settings,omitempty"`
}

// WebhookResponse ответ с вебхуком
type WebhookResponse struct {
	Result Webhook `json:"result"`
}

// WebhooksListResponse ответ со списком вебхуков
type WebhooksListResponse struct {
	Result []Webhook `json:"result"`
}

// WebhookEventTypes типы событий вебхуков
const (
	WebhookEventEmailSent         = "email_sent"
	WebhookEventEmailDelivered    = "email_delivered"
	WebhookEventEmailOpened       = "email_opened"
	WebhookEventEmailClicked      = "email_clicked"
	WebhookEventEmailUnsubscribed = "email_unsubscribed"
	WebhookEventEmailBounced      = "email_bounced"
	WebhookEventEmailComplained   = "email_complained"
)
