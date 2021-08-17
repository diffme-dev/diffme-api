package services

import (
	"diffme.dev/diffme-api/api/protos"
	"diffme.dev/diffme-api/cmd/workers"
	"diffme.dev/diffme-api/internal/core/infra"
	"diffme.dev/diffme-api/internal/modules/snapshots"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/hibiken/asynq"
)

func SnapshotCreated(previous domain.Snapshot, next domain.Snapshot) {

	client := infra.NewAsynqClient()

	nextData, _ := json.Marshal(next.Data)

	person := &protos.SnapshotCreatedEvent{
		Editor:      next.Editor,
		Data:        string(nextData),
		ReferenceId: next.ReferenceId,
		Metadata:    "some-metadata",
	}

	data, err := proto.Marshal(person)

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
