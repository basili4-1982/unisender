package resources

import (
	"context"

	"github.com/basili4-1982/unisender-go-api/models"
)

type System struct {
	client ClientGetter
}

func NewSystemResource(client ClientGetter) *System {
	return &System{client: client}
}

func (s System) Info(ctx context.Context) (*models.SystemInfo, error) {
	//TODO implement me
	panic("implement me")
}
