package UseCases

import (
	domain "diffme.dev/diffme-api/modules/events"
)

type eventUseCase struct {
	eventRepo domain.EventRepository
}

func NewEventUseCase(eventRepo domain.EventRepository) domain.EventUseCases {
	return &eventUseCase{
		eventRepo,
	}
}
