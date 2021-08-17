package services

import (
	"diffme.dev/diffme-api/cmd/workers"
	"diffme.dev/diffme-api/internal/core/infra"
	"diffme.dev/diffme-api/internal/modules/snapshots"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
)

type SnapshotCreatedEvent struct {
	previous domain.Snapshot
	next     domain.Snapshot
}

func SnapshotCreated(previous domain.Snapshot, next domain.Snapshot) {

	client := infra.NewAsynqClient()

	event := SnapshotCreatedEvent{
		previous: previous,
		next:     next,
	}

	payload, err := json.Marshal(event)

	if err != nil {
		println("decode json failed")
	}

	log.Printf("Event: %s", payload)

	task := asynq.NewTask(workers.SnapshotCreated, payload)

	_, err = client.Enqueue(task)

	if err != nil {
		println(err)
		// failed to send the task
		// do something with the error
	}
}
