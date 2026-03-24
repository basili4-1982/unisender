package resources

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
	"resty.dev/v3"

	"github.com/basili4-1982/unisender/config"
)

// ClientGetter интерфейс для получения HTTP клиента и конфигурации
type ClientGetter interface {
	GetResty() *resty.Client
	GetConfig() *config.Config
}

// BaseResource базовый ресурс с общей функциональностью
type BaseResource struct {
	client ClientGetter
}

// NewBaseResource создает базовый ресурс
func NewBaseResource(client ClientGetter) BaseResource {
	return BaseResource{client: client}
}

// getResty возвращает resty клиент
func (r *BaseResource) getResty() *resty.Client {
	return r.client.GetResty()
}

// getConfig возвращает конфигурацию
func (r *BaseResource) getConfig() *config.Config {
	return r.client.GetConfig()
}

// RequestOptions опции для запроса
type RequestOptions struct {
	QueryParams interface{}
	Headers     map[string]string
}

// executeGet выполняет GET запрос
func (r *BaseResource) executeGet(ctx context.Context, path string, result interface{}, opts *RequestOptions) error {
	req := r.getResty().R().SetContext(ctx)

	if opts != nil {
		if opts.QueryParams != nil {
			// Конвертируем struct в query parameters
			values, err := query.Values(opts.QueryParams)
			if err == nil {
				req.SetQueryParamsFromValues(values)
			} else {
				// Если не удалось сконвертировать, пробуем как map
				if paramsMap, ok := opts.QueryParams.(map[string]string); ok {
					req.SetQueryParams(paramsMap)
				}
			}
		}
		if opts.Headers != nil {
			req.SetHeaders(opts.Headers)
		}
	}

	resp, err := req.SetResult(result).Get(path)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("API error: %s", resp.String())
	}

	return nil
}

// ExecutePost выполняет POST запрос
func (r *BaseResource) ExecutePost(ctx context.Context, path string, body interface{}, result interface{}, opts *RequestOptions) error {
	req := r.getResty().R().SetContext(ctx).SetBody(body)

	if opts != nil && opts.Headers != nil {
		req.SetHeaders(opts.Headers)
	}

	resp, err := req.SetResult(result).Post(path)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("API error: %s", resp.String())
	}

	return nil
}

// executePut выполняет PUT запрос
func (r *BaseResource) executePut(ctx context.Context, path string, body interface{}, result interface{}, opts *RequestOptions) error {
	req := r.getResty().R().SetContext(ctx).SetBody(body)

	if opts != nil && opts.Headers != nil {
		req.SetHeaders(opts.Headers)
	}

	resp, err := req.SetResult(result).Put(path)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("API error: %s", resp.String())
	}

	return nil
}

// executeDelete выполняет DELETE запрос
func (r *BaseResource) executeDelete(ctx context.Context, path string, opts *RequestOptions) error {
	req := r.getResty().R().SetContext(ctx)

	if opts != nil && opts.Headers != nil {
		req.SetHeaders(opts.Headers)
	}

	resp, err := req.Delete(path)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("API error: %s", resp.String())
	}

	return nil
}
