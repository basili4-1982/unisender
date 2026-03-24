package models

// CreateEventDumpRequest запрос на создание выгрузки событий
type CreateEventDumpRequest struct {
	StartTime   string       `json:"start_time"`
	EndTime     string       `json:"end_time"`
	Limit       int          `json:"limit,omitempty"`
	AllProjects bool         `json:"all_projects,omitempty"`
	Filter      *EventFilter `json:"filter,omitempty"`
	DumpFields  []string     `json:"dump_fields,omitempty"`
	Aggregate   string       `json:"aggregate,omitempty"`
	Delimiter   string       `json:"delimiter,omitempty"`
	Format      string       `json:"format,omitempty"`
}

// EventFilter фильтр для выгрузки событий
type EventFilter struct {
	JobID          string `json:"job_id,omitempty"`
	Status         string `json:"status,omitempty"`
	DeliveryStatus string `json:"delivery_status,omitempty"`
	Email          string `json:"email,omitempty"`
	EmailFrom      string `json:"email_from,omitempty"`
	Domain         string `json:"domain,omitempty"`
	CampaignID     string `json:"campaign_id,omitempty"`
}

// EventDump ответ с выгрузкой событий
type EventDump struct {
	Status     string          `json:"status"`
	EventDumps []EventDumpItem `json:"event_dumps"`
}

type EventDumpItem struct {
	DumpId     string `json:"dump_id"`
	DumpStatus string `json:"dump_status"`
	Files      []struct {
		Url  string `json:"url"`
		Size int    `json:"size"`
	} `json:"files"`
}

// EventDumpResponse ответ на создание выгрузки
type EventDumpResponse struct {
	Result EventDump `json:"result"`
}

// EventDumpsListResponse ответ со списком выгрузок
type EventDumpsListResponse struct {
	Result []EventDump `json:"result"`
}
