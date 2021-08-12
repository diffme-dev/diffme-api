package http

import (
	domain "diffme.dev/diffme-api/modules/events"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	EventUseCase domain.EventUseCases
}

func Add(f fiber.Router, eventUseCases domain.EventUseCases) {
	eventRoutes := f.Group("/events")

	handler := &EventHandler{
		eventUseCases,
	}

	eventRoutes.Get("/", handler.GetEvents)
	eventRoutes.Get("/search", handler.SearchEvents)
	eventRoutes.Get("/:id", handler.GetEventById)
	eventRoutes.Post("/", handler.CreateEvent)
}
