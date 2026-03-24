package models

// SendEmailRequest запрос на отправку email
type SendEmailRequest struct {
	Message EmailMessage `json:"message"`
}

// EmailMessage структура письма
type EmailMessage struct {
	// Получатели
	Recipients []Recipient `json:"recipients"`

	// Шаблон
	TemplateID     string `json:"template_id,omitempty"`
	TemplateEngine string `json:"template_engine,omitempty"` // velocity, twig, handlebars, simple

	// Теги
	Tags []string `json:"tags,omitempty"`

	// Настройки отписки (0 или 1)
	SkipUnsubscribe int `json:"skip_unsubscribe,omitempty"`

	// Язык
	GlobalLanguage string `json:"global_language,omitempty"`

	// Глобальные подстановки
	GlobalSubstitutions map[string]string `json:"global_substitutions,omitempty"`

	// Глобальные метаданные
	GlobalMetadata map[string]string `json:"global_metadata,omitempty"`

	// Тело письма (если без шаблона)
	Body EmailBody `json:"body,omitempty"`

	// Отправитель
	Subject   string `json:"subject"`
	FromEmail string `json:"from_email"`
	FromName  string `json:"from_name,omitempty"`

	// Ответы
	ReplyTo     string `json:"reply_to,omitempty"`
	ReplyToName string `json:"reply_to_name,omitempty"`

	// Отслеживание (0 - выкл, 1 - вкл)
	TrackLinks int `json:"track_links,omitempty"`
	TrackRead  int `json:"track_read,omitempty"`

	// Настройки обхода проверок (0 или 1)
	BypassGlobal       int `json:"bypass_global,omitempty"`
	BypassUnavailable  int `json:"bypass_unavailable,omitempty"`
	BypassUnsubscribed int `json:"bypass_unsubscribed,omitempty"`
	BypassComplained   int `json:"bypass_complained,omitempty"`

	// Дополнительные заголовки
	Headers map[string]string `json:"headers,omitempty"`

	// Вложения
	Attachments       []Attachment `json:"attachments,omitempty"`
	InlineAttachments []Attachment `json:"inline_attachments,omitempty"`

	// Опции (дополнительные параметры)
	Options map[string]interface{} `json:"options,omitempty"`

	// Идемпотентность
	IdempotenceKey string `json:"idempotence_key,omitempty"`
}

// Recipient получатель письма
type Recipient struct {
	// Обязательные поля
	Email string `json:"email"`

	// Подстановки для конкретного получателя
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// Метаданные для конкретного получателя
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// EmailBody тело письма
type EmailBody struct {
	HTML      string `json:"html,omitempty"`
	Plaintext string `json:"plaintext,omitempty"`
	AMP       string `json:"amp,omitempty"`
}

// Attachment вложение
type Attachment struct {
	// MIME тип: text/plain, image/gif, application/pdf и т.д.
	Type string `json:"type"`

	// Имя файла
	Name string `json:"name"`

	// Контент в base64
	Content string `json:"content"`
}

// SendEmailResponse расширенный ответ на отправку email
type SendEmailResponse struct {
	Status       string            `json:"status"`
	JobID        string            `json:"job_id"`
	Emails       []string          `json:"emails"`
	FailedEmails map[string]string `json:"failed_emails,omitempty"`

	// Дополнительные поля
	TotalRecipients int `json:"total_recipients,omitempty"`
	SuccessCount    int `json:"success_count,omitempty"`
	FailedCount     int `json:"failed_count,omitempty"`

	// Детальная статистика
	Statistics *SendStatistics `json:"statistics,omitempty"`
}

// SendStatistics статистика отправки
type SendStatistics struct {
	Queued       int `json:"queued"`
	Sent         int `json:"sent"`
	Delivered    int `json:"delivered"`
	Opened       int `json:"opened"`
	Clicked      int `json:"clicked"`
	Bounced      int `json:"bounced"`
	Complained   int `json:"complained"`
	Unsubscribed int `json:"unsubscribed"`
}

type SubscribeRequest struct {
	FromEmail string `json:"from_email"`
	FromName  string `json:"from_name"`
	ToEmail   string `json:"to_email"`
}

type SubscribeResponse struct {
	Status       string            `json:"status"`
	JobID        string            `json:"job_id,omitempty"`
	Emails       []string          `json:"emails,omitempty"`
	FailedEmails map[string]string `json:"failed_emails,omitempty"`
}
