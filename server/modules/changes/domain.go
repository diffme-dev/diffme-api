package domain

import (
	"diffme.dev/diffme-api/api/protos"
	"diffme.dev/diffme-api/server/core/interfaces"
	"time"
)

type ChangeDiff interfaces.DiffOperation

type Change struct {
	Id          string                 `json:"id"`
	Label       *string                `json:"label"`
	EventName   *string                `json:"event_name"`
	ChangeSetId string                 `json:"change_set_id"`
	ReferenceId string                 `json:"reference_id"`
	SnapshotId  string                 `json:"current_snapshot_id"`
	Editor      string                 `json:"editor"`
	Metadata    map[string]interface{} `json:"metadata"`
	Diff        ChangeDiff             `json:"diff"`
	UpdatedAt   time.Time              `json:"updated_at"`
	CreatedAt   time.Time              `json:"created_at"`
}

type SearchChange struct {
	Id          string                 `json:"id"`
	ChangeSetId string                 `json:"change_set_id"`
	SnapshotId  string                 `json:"snapshot_id"`
	ReferenceId string                 `json:"reference_id"`
	Editor      string                 `json:"editor"`
	Metadata    map[string]interface{} `json:"metadata"`
	Diff        ChangeDiff             `json:"diff"`
	UpdatedAt   time.Time              `json:"updated_at"`
	CreatedAt   time.Time              `json:"created_at"`
}

type SearchRequest struct {
	ReferenceIds *[]string `json:"reference_ids"`
	Editor       *string   `json:"editor,omitempty"`
	Field        *string   `json:"field,omitempty"`
	Value        *string   `json:"value,omitempty"`
}

type QueryChangesRequest struct {
	Limit  *int    `json:"limit",omitempty`
	Sort   *string `json:"sort",omitempty`
	Before *string `json:"before",omitempty`
	After  *string `json:"after",omitempty`
}

type ChangeRepository interface {
	FindById(id string) (Change, error)
	Find(query QueryChangesRequest) ([]Change, error)
	FindByReferenceId(referenceId string) ([]Change, error)
	Create(change Change) (Change, error)
	CreateMultiple(change []Change) ([]Change, error)
}

type SearchChangeRepository interface {
	Query(match SearchRequest) ([]SearchChange, error)
	Create(change SearchChange) (SearchChange, error)
}

type ChangeUseCases interface {
	GetChanges(query QueryChangesRequest) ([]Change, error)
	CreateChange(
		eventName string,
		currentSnapshot protos.Snapshot,
		previousData map[string]interface{},
		currentData map[string]interface{},
	) ([]Change, error)
	SearchChange(query SearchRequest) ([]SearchChange, error)
	IndexSearchableChange(change Change) (SearchChange, error)
	FetchChangeForReferenceId(referenceId string) ([]Change, error)
}
