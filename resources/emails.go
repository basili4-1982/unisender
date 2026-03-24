package resources

import (
	"context"
	"fmt"

	"github.com/basili4-1982/unisender/errors"
	"github.com/basili4-1982/unisender/interfaces"
	"github.com/basili4-1982/unisender/models"
)

// EmailsResource реализация для работы с email
type EmailsResource struct {
	BaseResource
}

// NewEmailsResource создает новый email ресурс
func NewEmailsResource(client ClientGetter) *EmailsResource {
	return &EmailsResource{
		BaseResource: NewBaseResource(client),
	}
}

// Send отправляет email
func (r *EmailsResource) Send(ctx context.Context, req *models.SendEmailRequest) (*models.SendEmailResponse, error) {
	// Валидация запроса
	if err := req.Validate(); err != nil {
		return nil, &errors.ValidationError{
			Field:   "request",
			Message: err.Error(),
		}
	}

	var result models.SendEmailResponse

	cfg := r.getConfig()
	if cfg == nil {
		panic("config is nil")
	}

	err := r.ExecutePost(ctx,
		"/email/send.json",
		req,
		&result, &RequestOptions{
			Headers: map[string]string{
				"Content-Type": "application/json",
				"Accept":       "application/json",
				"X-API-KEY":    cfg.APIKey,
			},
		})
	if err != nil {
		return nil, fmt.Errorf("send_email: %w", err)
	}

	return &result, nil
}

// Subscribe подтверждает подписку
func (r *EmailsResource) Subscribe(ctx context.Context, req *models.SubscribeRequest) (*models.SubscribeResponse, error) {
	// Валидация
	if req.FromEmail == "" {
		return nil, &errors.ValidationError{
			Field:   "from_email",
			Message: "from_email is required",
		}
	}

	if req.FromName == "" {
		return nil, &errors.ValidationError{
			Field:   "from_name",
			Message: "from_name is required",
		}
	}

	if req.ToEmail == "" {
		return nil, &errors.ValidationError{
			Field:   "to_email",
			Message: "to_email is required",
		}
	}

	if err := models.ValidateEmail(req.FromEmail); err != nil {
		return nil, &errors.ValidationError{
			Field:   "email",
			Message: err.Error(),
		}
	}

	cfg := r.getConfig()
	if cfg == nil {
		panic("config is nil")
	}

	var result models.SubscribeResponse

	err := r.ExecutePost(ctx, "/email/subscribe.json", req, &result, &RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"X-API-KEY":    cfg.APIKey,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("subscribe: %w", err)
	}

	return &result, nil
}

// Убеждаемся, что EmailsResource реализует интерфейс
var _ interfaces.EmailsResourceInterface = (*EmailsResource)(nil)
