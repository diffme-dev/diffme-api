package domain

import (
	"time"
)

type Change struct {
	ID                 string    `json:"id"`
	ChangeSetID        string    `json:"change_set_id"`
	ReferenceID        string    `json:"reference_id"`
	PreviousSnapshotID string    `json:"previous_snapshot_id"`
	CurrentSnapshotID  string    `json:"current_snapshot_id"`
	Editor             string    `json:"id"`
	Metadata           []byte    `json:"metadata"`
	Diffs              []byte    `json:"diffs"`
	UpdatedAt          time.Time `json:"updated_at"`
	CreatedAt          time.Time `json:"created_at"`
}

type ChangeUseCases interface {
	CreateChange(oldSnapshot []byte, newSnapshot []byte) ([]*Change, error)
}

type ChangeRepository interface {
	FindByID(id string) (Change, error)
	Create(change Change) (Change, error)
}

type SearchChangeRepository interface {
	Query(query string) (Change, error)
	Create(change Change) (Change, error)
}
