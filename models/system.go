package models

// SystemInfo системная информация
type SystemInfo struct {
	// Версия API
	Version string `json:"version"`

	// Информация о сервисе
	Service ServiceInfo `json:"service"`

	// Ограничения
	Limits SystemLimits `json:"limits"`
}

// ServiceInfo информация о сервисе
type ServiceInfo struct {
	// Название
	Name string `json:"name"`

	// Статус
	Status string `json:"status"`

	// Актуальная версия
	Version string `json:"version"`
}

// SystemLimits системные ограничения
type SystemLimits struct {
	// Максимум получателей в одном запросе
	MaxRecipients int `json:"max_recipients"`

	// Максимум вложений
	MaxAttachments int `json:"max_attachments"`

	// Максимальный размер вложения (в байтах)
	MaxAttachmentSize int64 `json:"max_attachment_size"`

	// Максимум писем в день для проекта
	DefaultDailyLimit int `json:"default_daily_limit"`
}

// SystemInfoResponse ответ с системной информацией
type SystemInfoResponse struct {
	Result SystemInfo `json:"result"`
}
