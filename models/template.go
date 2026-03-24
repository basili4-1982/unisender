package models

// Template структура шаблона
type Template struct {
	// ID шаблона
	ID string `json:"id"`

	// Название шаблона
	Name string `json:"name"`

	// Тип шаблона: 'design' или 'code'
	Type string `json:"type"`

	// Тема письма
	Subject string `json:"subject"`

	// Отправитель
	FromEmail string `json:"from_email"`
	FromName  string `json:"from_name"`

	// Тело шаблона
	Body TemplateBody `json:"body"`

	// Метаданные
	Metadata Metadata `json:"metadata"`
}

// TemplateBody тело шаблона
type TemplateBody struct {
	HTML      string `json:"html,omitempty"`
	Plaintext string `json:"plaintext,omitempty"`
	AMP       string `json:"amp,omitempty"`
}

// SetTemplateRequest запрос на создание/обновление шаблона
type SetTemplateRequest struct {
	// ID шаблона (для обновления)
	ID string `json:"id,omitempty"`

	// Название шаблона
	Name string `json:"name"`

	// Тип шаблона: 'design' или 'code'
	Type string `json:"type"`

	// Тема письма
	Subject string `json:"subject"`

	// Отправитель
	FromEmail string `json:"from_email"`
	FromName  string `json:"from_name,omitempty"`

	// Тело шаблона
	Body TemplateBody `json:"body"`

	// Язык шаблона по умолчанию
	DefaultLanguage string `json:"default_language,omitempty"`
}

// TemplateResponse ответ с шаблоном
type TemplateResponse struct {
	Result Template `json:"result"`
}

// TemplatesListResponse ответ со списком шаблонов
type TemplatesListResponse struct {
	Result []Template `json:"result"`
}
