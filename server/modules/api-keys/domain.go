package api_keys

import (
	"time"
)

type ApiKey struct {
	Id             string    `json:"id"`
	OrganizationId string    `json:"organization_id"`
	Label          string    `json:"label"`
	ApiKey         string    `json:"api_key"`
	ACL            string    `json:"acl"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreateApiKeyParams struct {
	OrganizationId string `json:"organization_id"`
	Label          string `json:"label"`
	ApiKey         string `json:"api_key"`
	ACL            string `json:"acl"`
}

type ApiKeyRepository interface {
	FindById(id string) (*ApiKey, error)
	Create(apiKey ApiKey) (*ApiKey, error)
}

type ApiKeyUseCases interface {
	CreateApiKey(apiKey CreateApiKeyParams) (*ApiKey, error)
}
