package resources

import (
	"context"

	"github.com/basili4-1982/unisender-go-api/models"
)

type WebhooksResource struct {
	client ClientGetter
}

func NewWebhooksResource(client ClientGetter) *WebhooksResource {
	return &WebhooksResource{client: client}
}

func (w WebhooksResource) Set(ctx context.Context, req *models.SetWebhookRequest) (*models.Webhook, error) {
	//TODO implement me
	panic("implement me")
}

func (w WebhooksResource) Get(ctx context.Context, webhookID string) (*models.Webhook, error) {
	//TODO implement me
	panic("implement me")
}

func (w WebhooksResource) List(ctx context.Context) ([]*models.Webhook, error) {
	//TODO implement me
	panic("implement me")
}

func (w WebhooksResource) Delete(ctx context.Context, webhookID string) error {
	//TODO implement me
	panic("implement me")
}
