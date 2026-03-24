package models

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateEmail проверяет корректность email
func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email is required")
	}

	// Простая валидация email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format: %s", email)
	}

	return nil
}

// ValidateSendEmailRequest валидирует запрос на отправку
func (r *SendEmailRequest) Validate() error {
	if r.Message.Subject == "" {
		return fmt.Errorf("subject is required")
	}

	if r.Message.FromEmail == "" {
		return fmt.Errorf("from_email is required")
	}

	if err := ValidateEmail(r.Message.FromEmail); err != nil {
		return fmt.Errorf("invalid from_email: %w", err)
	}

	if len(r.Message.Recipients) == 0 {
		return fmt.Errorf("at least one recipient is required")
	}

	for i, recipient := range r.Message.Recipients {
		if err := ValidateEmail(recipient.Email); err != nil {
			return fmt.Errorf("recipient[%d]: %w", i, err)
		}
	}

	// Проверка: либо шаблон, либо тело письма
	if r.Message.TemplateID == "" && r.Message.Body.HTML == "" && r.Message.Body.Plaintext == "" {
		return fmt.Errorf("either template_id or body is required")
	}

	return nil
}

// ValidateSetTemplateRequest валидирует запрос на создание шаблона
func (r *SetTemplateRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}

	if r.Subject == "" {
		return fmt.Errorf("subject is required")
	}

	if r.FromEmail == "" {
		return fmt.Errorf("from_email is required")
	}

	if err := ValidateEmail(r.FromEmail); err != nil {
		return fmt.Errorf("invalid from_email: %w", err)
	}

	if r.Type != "design" && r.Type != "code" {
		return fmt.Errorf("type must be 'design' or 'code'")
	}

	if r.Body.HTML == "" && r.Body.Plaintext == "" {
		return fmt.Errorf("either html or plaintext body is required")
	}

	return nil
}

// ValidateSetWebhookRequest валидирует запрос на создание вебхука
func (r *SetWebhookRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}

	if r.URL == "" {
		return fmt.Errorf("url is required")
	}

	if !strings.HasPrefix(r.URL, "http://") && !strings.HasPrefix(r.URL, "https://") {
		return fmt.Errorf("url must start with http:// or https://")
	}

	if len(r.Events) == 0 {
		return fmt.Errorf("at least one event is required")
	}

	validEvents := map[string]bool{
		WebhookEventEmailSent:         true,
		WebhookEventEmailDelivered:    true,
		WebhookEventEmailOpened:       true,
		WebhookEventEmailClicked:      true,
		WebhookEventEmailUnsubscribed: true,
		WebhookEventEmailBounced:      true,
		WebhookEventEmailComplained:   true,
	}

	for _, event := range r.Events {
		if !validEvents[event.Type] {
			return fmt.Errorf("invalid event type: %s", event.Type)
		}
	}

	return nil
}

// ValidateCreateProjectRequest валидирует запрос на создание проекта
func (r *CreateProjectRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}

	if len(r.Name) > 100 {
		return fmt.Errorf("name must be less than 100 characters")
	}

	return nil
}

// ValidateSetSuppressionRequest валидирует запрос на добавление в стоп-лист
func (r *SetSuppressionRequest) Validate() error {
	if len(r.Items) == 0 {
		return fmt.Errorf("at least one item is required")
	}

	for i, item := range r.Items {
		if item.Value == "" {
			return fmt.Errorf("item[%d]: value is required", i)
		}

		if item.Type != "email" && item.Type != "domain" {
			return fmt.Errorf("item[%d]: type must be 'email' or 'domain'", i)
		}

		if item.Type == "email" {
			if err := ValidateEmail(item.Value); err != nil {
				return fmt.Errorf("item[%d]: %w", i, err)
			}
		}
	}

	return nil
}

// ValidateCreateEventDumpRequest валидирует запрос на создание выгрузки
func (r *CreateEventDumpRequest) Validate() error {
	if len(r.EventTypes) == 0 {
		return fmt.Errorf("at least one event type is required")
	}

	if r.DateFrom.IsZero() {
		return fmt.Errorf("date_from is required")
	}

	if r.DateTo.IsZero() {
		return fmt.Errorf("date_to is required")
	}

	if r.DateFrom.After(r.DateTo) {
		return fmt.Errorf("date_from must be before date_to")
	}

	if r.Format != "" && r.Format != "json" && r.Format != "csv" {
		return fmt.Errorf("format must be 'json' or 'csv'")
	}

	return nil
}
