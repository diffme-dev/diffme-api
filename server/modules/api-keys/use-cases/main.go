package use_cases

import (
	ApiKeyDomain "diffme.dev/diffme-api/server/modules/api-keys"
)

type ApiKeyUseCases struct {
	apiKeyRepo ApiKeyDomain.ApiKeyRepository
}

func NewApiKeyUseCases(apiKeyRepo ApiKeyDomain.ApiKeyRepository) ApiKeyDomain.ApiKeyUseCases {
	return &ApiKeyUseCases{
		apiKeyRepo: apiKeyRepo,
	}
}
