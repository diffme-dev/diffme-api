package http

import (
	domain "diffme.dev/diffme-api/internal/modules/changes"
	"github.com/gofiber/fiber/v2"
)

func ChangeRoutes(f fiber.Router, changeUseCases domain.ChangeUseCases) {
	eventRoutes := f.Group("/changes")

	handler := &ChangeController{
		changeUseCases: changeUseCases,
	}

	eventRoutes.Get("/", handler.GetChanges)
	eventRoutes.Get("/search", handler.SearchChanges)
	eventRoutes.Get("/references/:id", handler.GetChangeByReferenceID)
}
