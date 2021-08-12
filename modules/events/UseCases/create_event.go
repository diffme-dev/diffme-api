package UseCases

import (
	domain "diffme.dev/diffme-api/modules/events"
)

type UseCase struct {
	eventRepo domain.EventRepository
}

func (a *eventUseCase) CreateEvent(id string) (domain.Event, error) {

	// Get the author's id
	event := domain.Event{}

	return event, nil
}
