package http

import (
	"diffme.dev/diffme-api/internal/modules/changes"
	"github.com/gofiber/fiber/v2"
)

func ChangeRoutes(f fiber.Router, eventUseCases domain.ChangeUseCases) {
	eventRoutes := f.Group("/changes")

	handler := &ChangeController{
		eventUseCases,
	}

	eventRoutes.Get("/", handler.GetEvents)
	eventRoutes.Get("/search", handler.SearchEvents)
	eventRoutes.Get("/:id", handler.GetEventById)
	eventRoutes.Post("/", handler.CreateEvent)
}
