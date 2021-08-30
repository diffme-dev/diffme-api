package services

import (
	"diffme.dev/diffme-api/api/protos"
	"diffme.dev/diffme-api/cmd/workers"
	"diffme.dev/diffme-api/server/core/infra"
	"diffme.dev/diffme-api/server/modules/snapshots"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/hibiken/asynq"
)

func SnapshotCreated(previous domain.Snapshot, current domain.Snapshot) {

	client := infra.NewAsynqClient()

	previousData, _ := json.Marshal(previous.Data)
	previousMetadata, _ := json.Marshal(previous.Metadata)
	currentData, _ := json.Marshal(current.Data)
	currentMetadata, _ := json.Marshal(current.Metadata)

	event := &protos.SnapshotCreatedEvent{
		EventName: *current.EventName,
		Previous: &protos.Snapshot{
			Id:          previous.Id,
			Label:       *previous.Label,
			Editor:      previous.Editor,
			Data:        string(previousData),
			ReferenceId: previous.ReferenceId,
			Metadata:    string(previousMetadata),
		},
		Current: &protos.Snapshot{
			Id:          current.Id,
			Label:       *current.Label,
			Editor:      current.Editor,
			Data:        string(currentData),
			ReferenceId: current.ReferenceId,
			Metadata:    string(currentMetadata),
		},
	}

	data, err := proto.Marshal(event)

	if data != nil {
		println(data)
	}

	task := asynq.NewTask(workers.SnapshotCreated, data)

	_, err = client.Enqueue(task)

	if err != nil {
		println(err)
		// failed to send the task
		// do something with the error
	}
}
