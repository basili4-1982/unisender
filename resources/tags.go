package resources

import (
	"context"

	"github.com/basili4-1982/unisender-go-api/models"
)

type Tags struct {
	client ClientGetter
}

func NewTagsResource(client ClientGetter) *Tags {
	return &Tags{client: client}
}

func (t Tags) List(ctx context.Context) ([]*models.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (t Tags) Delete(ctx context.Context, tagID string) error {
	//TODO implement me
	panic("implement me")
}
