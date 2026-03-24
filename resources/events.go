package resources

import (
	"context"
	"fmt"

	"github.com/basili4-1982/unisender-go-api/models"
)

type Events struct {
	client ClientGetter
	BaseResource
}

func NewEventsResource(client ClientGetter) *Events {
	return &Events{client: client}
}

func (e Events) Create(ctx context.Context, req *models.CreateEventDumpRequest) (*models.EventDump, error) {
	res := models.EventDump{}

	cfg := e.getConfig()
	if cfg == nil {
		panic("config is nil")
	}

	err := e.ExecutePost(ctx, "event-dump/create.json", req, &res, &RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"X-API-KEY":    cfg.APIKey,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute post: %w", err)
	}

	return &res, nil
}

func (e Events) List(ctx context.Context) (*models.EventDump, error) {
	res := models.EventDump{}

	cfg := e.getConfig()
	if cfg == nil {
		panic("config is nil")
	}

	err := e.ExecutePost(ctx, "event-dump/list.json", nil, &res, &RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"X-API-KEY":    cfg.APIKey,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute post: %w", err)
	}

	return &res, nil
}

func (e Events) Delete(ctx context.Context, dumpID string) error {
	res := ""

	cfg := e.getConfig()
	if cfg == nil {
		panic("config is nil")
	}

	err := e.ExecutePost(ctx, "event-dump/delete.json", map[string]string{
		"dump_id": dumpID,
	}, &res, &RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"X-API-KEY":    cfg.APIKey,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to execute post: %w", err)
	}

	return nil
}
