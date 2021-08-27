package use_cases

import (
	"diffme.dev/diffme-api/api/protos"
	"diffme.dev/diffme-api/internal/core/interfaces"
	"diffme.dev/diffme-api/internal/modules/changes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/wI2L/jsondiff"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (a *ChangeUseCases) CreateChange(
	eventName string,
	currentSnapshot protos.Snapshot,
	previousData map[string]interface{},
	currentData map[string]interface{},
) ([]domain.Change, error) {

	changeSetID := "change_" + uuid.New().String()

	patch, err := jsondiff.Compare(previousData, currentData)

	if err != nil {
		return nil, err
	}

	var Metadata map[string]interface{}

	err = json.Unmarshal([]byte(currentSnapshot.Metadata), &Metadata)

	newChanges := make([]domain.Change, len(patch))

	for i, op := range patch {
		change := domain.Change{
			Id:          bson.NewObjectId().Hex(),
			Label:       &currentSnapshot.Label,
			EventName:   &eventName,
			Editor:      currentSnapshot.Editor,
			Metadata:    Metadata,
			SnapshotId:  currentSnapshot.Id,
			ReferenceId: currentSnapshot.ReferenceId,
			ChangeSetId: changeSetID,
			Diff: domain.ChangeDiff{
				From:     interfaces.StringPointer(op.From),
				Path:     interfaces.StringPointer(op.Path),
				Value:    op.Value,
				OldValue: op.OldValue,
				Type:     op.Type,
			},
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		}

		newChanges[i] = change
	}

	fmt.Printf("Changes (%d): %s\n\n", len(newChanges), newChanges)

	changes, err := a.changeRepo.CreateMultiple(newChanges)

	if err != nil {
		println(err)
		return nil, err
	}

	// fire off event to index with elastic search...
	for _, change := range changes {
		_, err := a.IndexSearchableChange(change)

		if err != nil {
			println(err)
		}
	}

	return newChanges, nil
}
