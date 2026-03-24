package resources

import (
	"context"

	"github.com/basili4-1982/unisender/models"
)

type Suppression struct {
	client ClientGetter
}

func NewSuppressionResource(client ClientGetter) *Suppression {
	return &Suppression{client: client}
}

func (s Suppression) Set(ctx context.Context, req *models.SetSuppressionRequest) (*models.SetSuppressionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Suppression) List(ctx context.Context, opts *models.ListOptions) ([]*models.Suppression, error) {
	//TODO implement me
	panic("implement me")
}
