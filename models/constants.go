package models

// Типы событий
const (
	EventTypeEmailSent         = "email_sent"
	EventTypeEmailDelivered    = "email_delivered"
	EventTypeEmailOpened       = "email_opened"
	EventTypeEmailClicked      = "email_clicked"
	EventTypeEmailUnsubscribed = "email_unsubscribed"
	EventTypeEmailBounced      = "email_bounced"
	EventTypeEmailComplained   = "email_complained"
	EventTypeEmailSpam         = "email_spam"
)

// Статусы выгрузки событий
const (
	EventDumpStatusPending    = "pending"
	EventDumpStatusProcessing = "processing"
	EventDumpStatusCompleted  = "completed"
	EventDumpStatusFailed     = "failed"
)

// Форматы выгрузки
const (
	ExportFormatJSON = "json"
	ExportFormatCSV  = "csv"
)

// Типы шаблонов
const (
	TemplateTypeDesign = "design"
	TemplateTypeCode   = "code"
)

// Типы движков шаблонов
const (
	TemplateEngineTwig       = "twig"
	TemplateEngineHandlebars = "handlebars"
	TemplateEngineSimple     = "simple"
)

// Типы стоп-листа
const (
	SuppressionTypeEmail  = "email"
	SuppressionTypeDomain = "domain"
)

// Статусы проекта
const (
	ProjectStatusActive    = "active"
	ProjectStatusSuspended = "suspended"
	ProjectStatusDeleted   = "deleted"
)

// Коды ошибок API
const (
	ErrorCodeInvalidAPIKey     = "invalid_api_key"
	ErrorCodeRateLimitExceeded = "rate_limit_exceeded"
	ErrorCodeValidationError   = "validation_error"
	ErrorCodeNotFound          = "not_found"
	ErrorCodeInternalError     = "internal_error"
	ErrorCodeTemplateNotFound  = "template_not_found"
	ErrorCodeProjectNotFound   = "project_not_found"
	ErrorCodeWebhookNotFound   = "webhook_not_found"
	ErrorCodeTagNotFound       = "tag_not_found"
)
