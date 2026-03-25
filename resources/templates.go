package resources

import (
	"context"
	"fmt"

	"github.com/basili4-1982/unisender/errors"
	"github.com/basili4-1982/unisender/interfaces"
	"github.com/basili4-1982/unisender/models"
)

// TemplatesResource реализация для работы с шаблонами
type TemplatesResource struct {
	BaseResource
}

// NewTemplatesResource создает новый ресурс шаблонов
func NewTemplatesResource(client ClientGetter) *TemplatesResource {
	return &TemplatesResource{
		BaseResource: NewBaseResource(client),
	}
}

// Set создает или обновляет шаблон
func (r *TemplatesResource) Set(ctx context.Context, req *models.SetTemplateRequest) (*models.Template, error) {
	// Валидация запроса
	if err := req.Validate(); err != nil {
		return nil, &errors.ValidationError{
			Field:   "request",
			Message: err.Error(),
		}
	}

	var result models.TemplateResponse

	err := r.ExecutePost(ctx, "/template/set", req, &result, nil)
	if err != nil {
		return nil, err
	}

	return &result.Result, nil
}

// Get получает шаблон по ID
func (r *TemplatesResource) Get(ctx context.Context, templateID string) (*models.Template, error) {
	if templateID == "" {
		return nil, &errors.ValidationError{
			Field:   "templateID",
			Message: "template ID is required",
		}
	}

	var result models.TemplateResponse
	path := fmt.Sprintf("/template/get/%s", templateID)

	err := r.executeGet(ctx, path, &result, nil)
	if err != nil {
		return nil, err
	}

	return &result.Result, nil
}

// List возвращает список шаблонов
func (r *TemplatesResource) List(ctx context.Context, opts *models.ListOptions) ([]models.Template, error) {
	var result models.TemplatesListResponse

	err := r.executeGet(ctx, "/template/list", &result, &RequestOptions{
		QueryParams: opts,
	})
	if err != nil {
		return nil, err
	}

	return result.Result, nil
}

// Delete удаляет шаблон
func (r *TemplatesResource) Delete(ctx context.Context, templateID string) error {
	if templateID == "" {
		return &errors.ValidationError{
			Field:   "templateID",
			Message: "template ID is required",
		}
	}

	path := fmt.Sprintf("/template/delete/%s", templateID)

	var res string

	err := r.ExecutePost(ctx, path, nil, &res, nil)
	if err != nil {
		return err
	}

	return nil
}

// Убеждаемся, что TemplatesResource реализует интерфейс
var _ interfaces.TemplatesResourceInterface = (*TemplatesResource)(nil)
