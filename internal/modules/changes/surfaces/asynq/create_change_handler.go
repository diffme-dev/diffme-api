package asynq

import (
	"context"
	domain "diffme.dev/diffme-api/internal/modules/snapshots"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

type CreateChangePayload struct {
	previous domain.Snapshot
	next     domain.Snapshot
}

func (e *ChangeAsynqSurface) CreateChangeHandler(ctx context.Context, t *asynq.Task) error {
	var payload CreateChangePayload

	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Previous: %+v", payload.previous.Data)
	log.Printf("Next: %+v", payload.next.Data)

	previousBytes, err := json.Marshal(payload.previous.Data)
	currentBytes, err := json.Marshal(payload.next.Data)

	changes, err := e.changeUseCases.CreateChange(previousBytes, currentBytes)

	if err != nil {
		return err
	}

	log.Printf("Change %s", changes)

	return nil
}
