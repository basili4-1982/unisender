package unisender

import (
	"encoding/json"
	"fmt"

	"resty.dev/v3"

	"github.com/basili4-1982/unisender-go-api/config"
	"github.com/basili4-1982/unisender-go-api/errors"
	"github.com/basili4-1982/unisender-go-api/interfaces"
	"github.com/basili4-1982/unisender-go-api/resources"
)

// Client реализация клиента Unisender
type Client struct {
	config *config.Config
	resty  *resty.Client

	emails      interfaces.EmailsResourceInterface
	templates   interfaces.TemplatesResourceInterface
	webhooks    interfaces.WebhooksResourceInterface
	projects    interfaces.ProjectsResourceInterface
	suppression interfaces.SuppressionResourceInterface
	tags        interfaces.TagsResourceInterface
	domains     interfaces.DomainsResourceInterface
	events      interfaces.EventsResourceInterface
	system      interfaces.SystemResourceInterface
}

// NewClient создает новый клиент Unisender
func NewClient(apiKey string, opts ...Option) (interfaces.ClientInterface, error) {
	cfg := DefaultConfig()
	cfg.APIKey = apiKey

	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.APIKey == "" {
		return nil, &errors.ValidationError{
			Field:   "apiKey",
			Message: "API key is required",
		}
	}

	// Настройка resty клиента
	restyClient := resty.New()
	restyClient.SetBaseURL(cfg.BaseURL)
	restyClient.SetHeader("X-API-KEY", cfg.APIKey)
	restyClient.SetHeader("Content-Type", "application/json")
	restyClient.SetHeader("Accept", "application/json")
	restyClient.SetHeader("User-Agent", cfg.UserAgent)
	restyClient.SetTimeout(cfg.Timeout)
	restyClient.SetDebug(cfg.Debug)

	// Настройка retry
	if cfg.RetryCount > 0 {
		restyClient.
			SetRetryCount(cfg.RetryCount).
			SetRetryWaitTime(cfg.RetryWaitTime).
			AddRetryConditions(func(r *resty.Response, err error) bool {
				if err != nil {
					return true
				}
				// Retry on 429 Too Many Requests and 5xx errors
				return r.StatusCode() == 429 || r.StatusCode() >= 500
			})
	}

	restyClient.OnSuccess(func(client *resty.Client, resp *resty.Response) {
		if resp.IsError() {
			client.SetError(handleErrorResponse(resp))
		}
	})

	client := &Client{
		config: cfg,
		resty:  restyClient,
	}

	client.emails = resources.NewEmailsResource(client)
	client.templates = resources.NewTemplatesResource(client)
	client.webhooks = resources.NewWebhooksResource(client)
	client.projects = resources.NewProjectsResource(client)
	client.suppression = resources.NewSuppressionResource(client)
	client.tags = resources.NewTagsResource(client)
	client.domains = resources.NewDomainsResource(client)
	client.events = resources.NewEventsResource(client)
	client.system = resources.NewSystemResource(client)

	return client, nil
}

// Emails возвращает ресурс для работы с email
func (c *Client) Emails() interfaces.EmailsResourceInterface {
	return c.emails
}

// Templates возвращает ресурс для работы с шаблонами
func (c *Client) Templates() interfaces.TemplatesResourceInterface {
	return c.templates
}

// Webhooks возвращает ресурс для работы с вебхуками
func (c *Client) Webhooks() interfaces.WebhooksResourceInterface {
	return c.webhooks
}

// Projects возвращает ресурс для работы с проектами
func (c *Client) Projects() interfaces.ProjectsResourceInterface {
	return c.projects
}

// Suppression возвращает ресурс для работы со стоп-листом
func (c *Client) Suppression() interfaces.SuppressionResourceInterface {
	return c.suppression
}

// Tags возвращает ресурс для работы с тегами
func (c *Client) Tags() interfaces.TagsResourceInterface {
	return c.tags
}

// Domains возвращает ресурс для работы с доменами
func (c *Client) Domains() interfaces.DomainsResourceInterface {
	return c.domains
}

// Events возвращает ресурс для работы с выгрузкой событий
func (c *Client) Events() interfaces.EventsResourceInterface {
	return c.events
}

// System возвращает ресурс для работы с системной информацией
func (c *Client) System() interfaces.SystemResourceInterface {
	return c.system
}

// GetConfig возвращает конфигурацию клиента
func (c *Client) GetConfig() *config.Config {
	if c.config == nil {
		c.config = DefaultConfig()
	}

	return c.config
}

func (c *Client) GetResty() *resty.Client {
	return c.resty
}

// Close закрывает клиент
func (c *Client) Close() error {
	err := c.resty.Close()
	if err != nil {
		return fmt.Errorf("close resty client: %w", err)
	}

	return nil
}

// handleErrorResponse обрабатывает ошибки API
func handleErrorResponse(resp *resty.Response) error {
	var apiErr errors.APIError
	apiErr.StatusCode = resp.StatusCode()

	err := json.Unmarshal([]byte(resp.String()), &apiErr)
	if err != nil {
		// Если не удалось распарсить, возвращаем общую ошибку
		return &errors.APIError{
			StatusCode: resp.StatusCode(),
			Message:    resp.String(),
		}
	}

	// Специфичная обработка для rate limit
	if resp.StatusCode() == 429 {
		retryAfter := resp.Header().Get("Retry-After")
		var retrySeconds int
		if retryAfter != "" {
			_, _ = fmt.Sscanf(retryAfter, "%d", &retrySeconds)
		}
		return &errors.RateLimitError{
			APIError:   &apiErr,
			RetryAfter: retrySeconds,
		}
	}

	return &apiErr
}
