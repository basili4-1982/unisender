package resources

import (
	"context"
	"fmt"

	"github.com/basili4-1982/unisender/models"
)

type Domains struct {
	client ClientGetter
	BaseResource
}

func NewDomainsResource(client ClientGetter) *Domains {
	return &Domains{client: client}
}

func (d Domains) GetDNS(ctx context.Context, req *models.DomainsRequest) (*models.GetDNSDomainsResponse, error) {
	resp := models.GetDNSDomainsResponse{}

	cfg := d.getConfig()
	if cfg == nil {
		panic("config is nil")
	}

	err := d.ExecutePost(ctx, "domain/get-dns-records.json", req, &resp, &RequestOptions{
		Headers: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
			"X-API-KEY":    cfg.APIKey,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error getting dns records: %w", err)
	}

	return &resp, nil
}

func (d Domains) ValidateVerification(ctx context.Context, req *models.DomainsRequest) (*models.DomainResponse, error) {
	resp := models.DomainResponse{}

	cfg := d.getConfig()
	if cfg == nil {
		panic("config is nil")
	}

	err := d.ExecutePost(ctx, "domain/validate-verification-record.json", req, &resp, &RequestOptions{
		Headers: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
			"X-API-KEY":    cfg.APIKey,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error validate verification record: %w", err)
	}

	return &resp, nil
}

func (d Domains) ValidateDkim(ctx context.Context, req *models.DomainsRequest) (*models.DomainResponse, error) {
	resp := models.DomainResponse{}

	cfg := d.getConfig()
	if cfg == nil {
		panic("config is nil")
	}

	err := d.ExecutePost(ctx, "domain/validate-dkim.json", req, &resp, &RequestOptions{
		Headers: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
			"X-API-KEY":    cfg.APIKey,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error validate verification record: %w", err)
	}

	return &resp, nil
}

func (d Domains) List(ctx context.Context, req *models.ListDomainsRequest) (*models.ListDomainsResponse, error) {
	resp := models.ListDomainsResponse{}

	cfg := d.getConfig()
	if cfg == nil {
		panic("config is nil")
	}

	err := d.ExecutePost(ctx, "domain/list.json", req, &resp, &RequestOptions{
		Headers: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
			"X-API-KEY":    cfg.APIKey,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error validate verification record: %w", err)
	}

	return &resp, nil
}

func (d Domains) Delete(ctx context.Context, req *models.DomainsRequest) (*models.DomainResponse, error) {
	resp := models.DomainResponse{}

	cfg := d.getConfig()
	if cfg == nil {
		panic("config is nil")
	}

	err := d.ExecutePost(ctx, "domain/delete.json", req, &resp, &RequestOptions{
		Headers: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
			"X-API-KEY":    cfg.APIKey,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error validate verification record: %w", err)
	}

	return &resp, nil
}
