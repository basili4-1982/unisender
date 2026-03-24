package resources

import (
	"context"

	"github.com/basili4-1982/unisender-go-api/models"
)

type Projects struct {
	client ClientGetter
}

func NewProjectsResource(client ClientGetter) *Projects {
	return &Projects{client: client}
}

func (p Projects) Create(ctx context.Context, req *models.CreateProjectRequest) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p Projects) Update(ctx context.Context, projectID string, req *models.UpdateProjectRequest) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p Projects) List(ctx context.Context) ([]*models.Project, error) {
	//TODO implement me
	panic("implement me")
}
