package http

import (
	domain "diffme.dev/diffme-api/server/modules/api-keys"
	"github.com/gofiber/fiber/v2"
)

func ApiKeyRoutes(f fiber.Router, apiKeyRepo domain.ApiKeyRepository) {
	users := f.Group("/api-keys")

	// If don't pass in members it is like calling stuff on nil/null
	controller := &ApiKeyController{
		apiKeyRepo: apiKeyRepo,
	}

	users.Get("/:id", controller.GetApiKeyById)
	users.Post("/", controller.CreateApiKey)
}
