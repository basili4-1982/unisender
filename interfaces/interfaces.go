package interfaces

import (
	"context"

	"github.com/basili4-1982/unisender/config"
	"github.com/basili4-1982/unisender/models"
)

// EmailsResourceInterface интерфейс для работы с email
type EmailsResourceInterface interface {
	Send(ctx context.Context, req *models.SendEmailRequest) (*models.SendEmailResponse, error)
	Subscribe(ctx context.Context, req *models.SubscribeRequest) (*models.SubscribeResponse, error)
}

// TemplatesResourceInterface интерфейс для работы с шаблонами
type TemplatesResourceInterface interface {
	Set(ctx context.Context, req *models.SetTemplateRequest) (*models.Template, error)
	Get(ctx context.Context, templateID string) (*models.Template, error)
	List(ctx context.Context, opts *models.ListOptions) ([]models.Template, error)
	Delete(ctx context.Context, templateID string) error
}

// WebhooksResourceInterface интерфейс для работы с вебхуками
type WebhooksResourceInterface interface {
	Set(ctx context.Context, req *models.SetWebhookRequest) (*models.Webhook, error)
	Get(ctx context.Context, webhookID string) (*models.Webhook, error)
	List(ctx context.Context) ([]*models.Webhook, error)
	Delete(ctx context.Context, webhookID string) error
}

// ProjectsResourceInterface интерфейс для работы с проектами
type ProjectsResourceInterface interface {
	Create(ctx context.Context, req *models.CreateProjectRequest) (*models.Project, error)
	Update(ctx context.Context, projectID string, req *models.UpdateProjectRequest) (*models.Project, error)
	List(ctx context.Context) ([]*models.Project, error)
}

// SuppressionResourceInterface интерфейс для работы со стоп-листом
type SuppressionResourceInterface interface {
	Set(ctx context.Context, req *models.SetSuppressionRequest) (*models.SetSuppressionResponse, error)
	List(ctx context.Context, opts *models.ListOptions) ([]*models.Suppression, error)
}

// TagsResourceInterface интерфейс для работы с тегами
type TagsResourceInterface interface {
	List(ctx context.Context) ([]*models.Tag, error)
	Delete(ctx context.Context, tagID string) error
}

// DomainsResourceInterface интерфейс для работы с доменами
type DomainsResourceInterface interface {
	Delete(ctx context.Context, req *models.DomainsRequest) (*models.DomainResponse, error)
	GetDNS(ctx context.Context, req *models.DomainsRequest) (*models.GetDNSDomainsResponse, error)
	ValidateVerification(ctx context.Context, req *models.DomainsRequest) (*models.DomainResponse, error)
	ValidateDkim(ctx context.Context, req *models.DomainsRequest) (*models.DomainResponse, error)
	List(ctx context.Context, req *models.ListDomainsRequest) (*models.ListDomainsResponse, error)
}

// EventsResourceInterface интерфейс для работы с выгрузкой событий
type EventsResourceInterface interface {
	Create(ctx context.Context, req *models.CreateEventDumpRequest) (*models.EventDump, error)
	List(ctx context.Context) (*models.EventDump, error)
	Delete(ctx context.Context, dumpID string) error
}

// SystemResourceInterface интерфейс для системной информации
type SystemResourceInterface interface {
	Info(ctx context.Context) (*models.SystemInfo, error)
}

// HTTPClientInterface интерфейс HTTP клиента для возможности мокинга
type HTTPClientInterface interface {
	Get(ctx context.Context, path string, result interface{}, opts ...RequestOption) error
	Post(ctx context.Context, path string, body interface{}, result interface{}, opts ...RequestOption) error
	Put(ctx context.Context, path string, body interface{}, result interface{}, opts ...RequestOption) error
	Delete(ctx context.Context, path string, opts ...RequestOption) error
}

// RequestOption опции для запроса
type RequestOption func(*RequestConfig)

// RequestConfig конфигурация запроса
type RequestConfig struct {
	QueryParams interface{}
	Headers     map[string]string
}

// ClientInterface основной интерфейс клиента
type ClientInterface interface {
	// Ресурсы
	Emails() EmailsResourceInterface
	Templates() TemplatesResourceInterface
	Webhooks() WebhooksResourceInterface
	Projects() ProjectsResourceInterface
	Suppression() SuppressionResourceInterface
	Tags() TagsResourceInterface
	Domains() DomainsResourceInterface
	Events() EventsResourceInterface
	System() SystemResourceInterface

	// Вспомогательные методы
	GetConfig() *config.Config
	Close() error
}
