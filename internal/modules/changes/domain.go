package domain

import (
	"diffme.dev/diffme-api/api/protos"
	"github.com/wI2L/jsondiff"
	"time"
)

type Diff jsondiff.Operation

type Change struct {
	Id          string                 `json:"id"`
	ChangeSetId string                 `json:"change_set_id"`
	ReferenceId string                 `json:"reference_id"`
	SnapshotId  string                 `json:"current_snapshot_id"`
	Editor      string                 `json:"editor"`
	Metadata    map[string]interface{} `json:"metadata"`
	Diff        Diff                   `json:"diff"`
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
	Diff        Diff                   `json:"diff"`
	UpdatedAt   time.Time              `json:"updated_at"`
	CreatedAt   time.Time              `json:"created_at"`
}

type SearchRequest struct {
	Editor *string `json:"editor",omitempty`
	Field  *string `json:"field",omitempty`
	Value  *string `json:"value",omitempty`
}

type ChangeRepository interface {
	FindByID(id string) (Change, error)
	Create(change Change) (Change, error)
	CreateMultiple(change []Change) ([]Change, error)
}

type SearchChangeRepository interface {
	Query(match SearchRequest) ([]SearchChange, error)
	Create(change SearchChange) (SearchChange, error)
}

type ChangeUseCases interface {
	CreateChange(
		currentSnapshot protos.Snapshot,
		previousData map[string]interface{},
		currentData map[string]interface{},
	) ([]Change, error)
	SearchChange(query SearchRequest) ([]SearchChange, error)
	IndexSearchableChange(change Change) (SearchChange, error)
	FetchChangeForReferenceId(referenceId string) ([]Change, error)
}
