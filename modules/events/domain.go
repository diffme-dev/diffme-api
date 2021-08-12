package domain

import (
	"context"
	"time"
)

type Event struct {
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type EventUseCases interface {
	CreateEvent(id string) (Event, error)
}

type EventRepository interface {
	GetByID(ctx context.Context, id string) (Event, error)
}
