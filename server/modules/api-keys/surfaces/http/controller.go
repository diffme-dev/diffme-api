package http

import (
	"diffme.dev/diffme-api/server/modules/api-keys"
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-uuid"
)

type ApiKeyController struct {
	apiKeyRepo     api_keys.ApiKeyRepository
	apiKeyUseCases api_keys.ApiKeyUseCases
}

func (e *ApiKeyController) CreateApiKey(c *fiber.Ctx) error {

	uuid, _ := uuid.GenerateUUID()

	data, err := e.apiKeyRepo.Create(api_keys.ApiKey{
		ApiKey:         "sk_" + uuid,
		Label:          "TODO:",
		OrganizationId: "TODO:",
		ACL:            "TODO:", // TODO: need to make a policy type or something. maybe link ACL to a db record
	})

	if err != nil {
		return err
	}

	return c.JSON(data)
}

func (e *ApiKeyController) GetApiKeyById(c *fiber.Ctx) error {
	apiKeyId := c.Params("id")

	data, err := e.apiKeyRepo.FindById(apiKeyId)

	if err != nil {
		return err
	}

	return c.JSON(data)
}
