package UseCases

import (
	"diffme.dev/diffme-api/internal/modules/changes"
	SnapshotDomain "diffme.dev/diffme-api/internal/modules/snapshots"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/wI2L/jsondiff"
	"log"
	"time"
)

func (a *ChangeUseCases) CreateChange(oldSnapshot []byte, newSnapshot []byte) ([]domain.Change, error) {

	changeSetID := "change_" + uuid.New().String()

	var previous SnapshotDomain.Snapshot
	var next SnapshotDomain.Snapshot

	err := json.Unmarshal(oldSnapshot, previous)
	err = json.Unmarshal(newSnapshot, next)

	//next.Data = map[string]interface{}{"name": "hello man"}

	fmt.Printf("\nPrevious Data %s", previous)
	fmt.Printf("\nNext Data %s", next.Data)

	patch, err := jsondiff.Compare(previous.Data, next.Data)

	if err != nil {
		return nil, err
	}

	fmt.Printf("\nJSON Patch: %s", patch)

	changes := make([]domain.Change, len(patch))

	for _, op := range patch {
		println("------------\n")
		log.Printf("OPERATION: %s", op)

		change := domain.Change{
			Id:          next.Id,
			Editor:      next.Editor,
			Metadata:    next.Metadata,
			SnapshotId:  next.Id,
			ReferenceId: next.ReferenceId,
			ChangeSetId: changeSetID,
			Diff:        domain.Diff(op),
			UpdatedAt:   time.Now(),
			CreatedAt:   time.Now(),
		}

		fmt.Printf("new change %s", change)

		changes = append(changes, change)
	}

	changes, err = a.changeRepo.CreateMultiple(changes)

	if err != nil {
		println(err)
		return nil, err
	}

	// fire off event to index with elastic search...
	for _, change := range changes {
		_, err = a.IndexSearchableChange(change)

		if err != nil {
			println(err)
		}
		// TODO: event driven architecture
		//services.ChangeCreated(change)
	}

	return changes, nil
}
