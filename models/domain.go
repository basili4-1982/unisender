package models

// DeleteDomainRequest запрос на удаление домена
type DeleteDomainRequest struct {
	// Домен для удаления
	Domain string `json:"domain"`

	// ID проекта (опционально)
	ProjectID string `json:"project_id,omitempty"`
}

type DomainsRequest struct {
	Domain string `json:"domain"`
}

type ListDomainsRequest struct {
	Domain string `json:"domain"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type GetDNSDomainsResponse struct {
	Status             string `json:"status"`
	Domain             string `json:"domain"`
	VerificationRecord string `json:"verification-record"`
	Dkim               string `json:"dkim"`
}

type DomainResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ListDomainsResponse struct {
	Status  string       `json:"status"`
	Domains []DomainItem `json:"domains"`
}

type DomainItem struct {
	Domain             string `json:"domain"`
	VerificationRecord struct {
		Value  string `json:"value"`
		Status string `json:"status"`
	} `json:"verification-record"`
	Dkim struct {
		Key    string `json:"key"`
		Status string `json:"status"`
	} `json:"dkim"`
}
