package domain

import (
	"time"
)

type Snapshot struct {
	ID          string    `json:"id"`
	ReferenceID string    `json:"id"`
	Data        []byte    `json:"data"`
	Editor      string    `json:"editor"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateSnapshotParams struct {
	ID          string    `json:"id" validate:"required"`
	ReferenceID string    `json:"reference_id" validate:"required"`
	Data        []byte    `json:"data" validate:"required"`
	Editor      string    `json:"editor" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type SnapshotUseCases interface {
	CreateSnapshot() (Snapshot, error)
	GetSnapshotByID(id string) (Snapshot, error)
}

type SnapshotRepo interface {
	FindByID(id string) (Snapshot, error)
	FindByReferenceID(referenceID string) (Snapshot, error)
	FindMostRecentByReference(referenceID string) (Snapshot, error)
	FindSnapshotsForReference(referenceID string) ([]Snapshot, error)
	CreateSnapshot(params CreateSnapshotParams) (Snapshot, error)
}
