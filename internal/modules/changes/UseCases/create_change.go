package UseCases

import (
	"diffme.dev/diffme-api/internal/modules/changes"
	"diffme.dev/diffme-api/internal/modules/changes/services"
	"encoding/json"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/google/uuid"
	"log"
)

type JSONPatch struct{}

type Diff []interface{}

type Result struct {
	Diffs []Diff
}

func (r *Result) UnmarshalJSON(p []byte) error {
	var tmp []Diff

	if err := json.Unmarshal(p, &tmp); err != nil {
		return err
	}

	r.Diffs = tmp
	return nil
}

func (a *ChangeUseCases) CreateChange(oldSnapshot []byte, newSnapshot []byte) ([]domain.Change, error) {

	changeSetID := "change_" + uuid.New().String()

	println(changeSetID)

	patch, err := jsonpatch.CreateMergePatch(oldSnapshot, newSnapshot)

	if err != nil {
		return nil, err
	}

	log.Printf("patch %+v", patch)

	// TODO:
	jsonData := []byte("hi")

	var r Result

	if err := json.Unmarshal(jsonData, &r); err != nil {
		log.Fatal(err)
	}

	changes := make([]domain.Change, len(r.Diffs))

	for _, diff := range r.Diffs {
		// Get the author's id
		diffByte, _ := json.Marshal(diff)

		change := domain.Change{
			ChangeSetID: changeSetID,
			Diffs:       diffByte,
		}

		changes = append(changes, change)
	}

	// TODO: save chages

	changes, err = a.changeRepo.CreateMultiple(changes)

	if err != nil {
		return nil, err
	}

	// fire off event to index with elastic search...
	for _, change := range changes {
		services.ChangeCreated(change)
	}

	return changes, nil
}
