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

	var nextData map[string]interface{}
	err = json.Unmarshal([]byte(payload.GetData()), &nextData)

	if err != nil {
		fmt.Printf("\nError: %s", err)
	}
	log.Printf("\nPayload: %s", payload.GetData())
	log.Printf("\nDecoded: %s", nextData)

	if err != nil {
		println(err)
	}
	//
	//log.Printf("\nPrevious: %+v", payload.previous)
	//log.Printf("\nNext: %+v", payload.next)
	//
	//previousBytes, err := json.Marshal(payload.previous.Data)
	//currentBytes, err := json.Marshal(payload.next.Data)
	//
	//changes, err := e.changeUseCases.CreateChange(previousBytes, currentBytes)
	//
	//if err != nil {
	//	return err
	//}
	//
	//log.Printf("Change %s", changes)

	return nil
}
