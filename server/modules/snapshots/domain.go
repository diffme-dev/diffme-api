package domain

import (
	"time"
)

type Snapshot struct {
	Id          string                 `json:"id"`
	Label       *string                `json:"label"`
	EventName   *string                `json:"event_name"`
	ReferenceId string                 `json:"reference_id"`
	Data        map[string]interface{} `json:"data" omitempty`
	Editor      string                 `json:"editor" omitempty`
	Metadata    map[string]interface{} `json:"metadata" omitempty`
	UpdatedAt   time.Time              `json:"updated_at" omitempty`
	CreatedAt   time.Time              `json:"created_at" omitempty`
}

type CreateSnapshotParams struct {
	Id        string                 `json:"id" validate:"required"`
	Label     *string                `json:"label" validate:"-"`
	EventName *string                `json:"event_name" validate:"-"`
	Data      map[string]interface{} `json:"data" validate:"required"`
	Editor    string                 `json:"editor" validate:"required"`
	Metadata  map[string]interface{} `json:"metadata" validate:"-"`
	CreatedAt time.Time              `json:"created_at"  validate:"required"`
}

type SnapshotUseCases interface {
	CreateSnapshot(params CreateSnapshotParams) (*Snapshot, error)
	GetSnapshotByID(id string) (*Snapshot, error)
	FindMostRecentReference(referenceId string, before *time.Time) (*Snapshot, error)
}

type SnapshotRepo interface {
	FindByID(id string) (*Snapshot, error)
	FindByReferenceID(referenceId string) (*Snapshot, error)
	FindMostRecentByReference(referenceId string, before *time.Time) (*Snapshot, error)
	FindForReference(referenceId string) ([]Snapshot, error)
	Create(params CreateSnapshotParams) (*Snapshot, error)
}
