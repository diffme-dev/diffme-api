package asynq

import (
	"context"
	"diffme.dev/diffme-api/api/protos"
	domain "diffme.dev/diffme-api/internal/modules/snapshots"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hibiken/asynq"
	"log"
)

type CreateChangePayload struct {
	previous domain.Snapshot `json:"previous"`
	next     domain.Snapshot `json:"next"`
}

func (e *ChangeAsynqSurface) CreateChangeHandler(ctx context.Context, t *asynq.Task) error {
	var payload protos.SnapshotCreatedEvent

	err := proto.Unmarshal(t.Payload(), &payload)

	previous := payload.GetPrevious()
	current := payload.GetCurrent()

	var previousData map[string]interface{}
	var currentData map[string]interface{}

	err = json.Unmarshal([]byte(previous.Data), &previousData)
	err = json.Unmarshal([]byte(current.Data), &currentData)

	if err != nil {
		fmt.Printf("\nError: %s", err)
	}

	log.Printf("\nPrevious: %s", previousData)
	log.Printf("\nCurrent: %s", currentData)

	if err != nil {
		println(err)
		return err
	}

	changes, err := e.changeUseCases.CreateChange(*current, previousData, currentData)

	if err != nil {
		return err
	}

	log.Printf("Change %s", changes)

	return nil
}
