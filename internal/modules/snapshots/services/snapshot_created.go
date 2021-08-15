package services

import (
	"diffme.dev/diffme-api/internal/modules/snapshots"
	"encoding/base64"
	"encoding/json"
	Machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
)

type SnapshotCreatedEvent struct {
	previous domain.Snapshot
	next     domain.Snapshot
}

func SnapshotCreated(taskserver *Machinery.Server, previous domain.Snapshot, next domain.Snapshot) {

	event := SnapshotCreatedEvent{
		previous: previous,
		next:     next,
	}

	encodedJSON, err := json.Marshal(event)

	if err != nil {
		println("decode json failed")
	}

	encodedRequest := base64.StdEncoding.EncodeToString(encodedJSON)

	signature := &tasks.Signature{
		Name: "SnapshotCreated",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: encodedRequest,
			},
		},
	}

	_, err = taskserver.SendTask(signature)

	if err != nil {
		println(err)
		// failed to send the task
		// do something with the error
	}
}
