package domain

import (
	"time"
)

type Snapshot struct {
	ID          string      `json:"id" omitempty`
	ReferenceID string      `json:"reference_id" omitempty`
	Data        interface{} `json:"data" omitempty`
	Editor      string      `json:"editor" omitempty`
	Metadata    interface{} `json:"metadata" omitempty`
	UpdatedAt   time.Time   `json:"updated_at" omitempty`
	CreatedAt   time.Time   `json:"created_at" omitempty`
}

type CreateSnapshotParams struct {
	ID        string      `json:"id" valid:"required"`
	Data      interface{} `json:"data" valid:"-"`
	Editor    string      `json:"editor" valid:"required"`
	Metadata  interface{} `json:"metadata" valid:"-"`
	CreatedAt time.Time   `json:"created_at"  valid:"required"`
}

type SnapshotUseCases interface {
	CreateSnapshot(params CreateSnapshotParams) (Snapshot, error)
	GetSnapshotByID(id string) (Snapshot, error)
	FindMostRecentReference(referenceID string) (Snapshot, error)
}

type SnapshotRepo interface {
	FindByID(id string) (Snapshot, error)
	FindByReferenceID(referenceID string) (Snapshot, error)
	FindMostRecentByReference(referenceID string) (Snapshot, error)
	FindForReference(referenceID string) ([]Snapshot, error)
	Create(params CreateSnapshotParams) (Snapshot, error)
}
