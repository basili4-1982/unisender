package models

// Project структура проекта
type Project struct {
	// ID проекта
	ID string `json:"id"`

	// Название проекта
	Name string `json:"name"`

	// Статус
	Status string `json:"status"`

	// Домены
	Domains []string `json:"domains,omitempty"`

	// Настройки
	Settings ProjectSettings `json:"settings"`

	// Метаданные
	Metadata Metadata `json:"metadata"`
}

// ProjectSettings настройки проекта
type ProjectSettings struct {
	// Лимиты
	Limits ProjectLimits `json:"limits"`

	// Язык по умолчанию
	DefaultLanguage string `json:"default_language,omitempty"`

	// Часовой пояс
	Timezone string `json:"timezone,omitempty"`
}

// ProjectLimits лимиты проекта
type ProjectLimits struct {
	// Максимум писем в день
	DailyLimit int `json:"daily_limit,omitempty"`

	// Максимум писем в месяц
	MonthlyLimit int `json:"monthly_limit,omitempty"`
}

// CreateProjectRequest запрос на создание проекта
type CreateProjectRequest struct {
	// Название проекта
	Name string `json:"name"`

	// Язык по умолчанию
	DefaultLanguage string `json:"default_language,omitempty"`

	// Часовой пояс
	Timezone string `json:"timezone,omitempty"`
}

// UpdateProjectRequest запрос на обновление проекта
type UpdateProjectRequest struct {
	// Название проекта
	Name string `json:"name,omitempty"`

	// Язык по умолчанию
	DefaultLanguage string `json:"default_language,omitempty"`

	// Часовой пояс
	Timezone string `json:"timezone,omitempty"`
}

// ProjectResponse ответ с проектом
type ProjectResponse struct {
	Result Project `json:"result"`
}

// ProjectsListResponse ответ со списком проектов
type ProjectsListResponse struct {
	Result []Project `json:"result"`
}
