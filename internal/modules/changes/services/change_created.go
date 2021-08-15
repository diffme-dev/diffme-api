package services

import (
	"diffme.dev/diffme-api/cmd/workers"
	"diffme.dev/diffme-api/internal/core/infra"
	domain "diffme.dev/diffme-api/internal/modules/changes"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
)

type ChangeCreatedPayload struct {
	change domain.Change
}

func ChangeCreated(change domain.Change) {

	client := infra.NewAsynqClient()

	event := ChangeCreatedPayload{
		change: change,
	}

	log.Printf("change created: %+v", event)

	payload, err := json.Marshal(event)

	if err != nil {
		println("decode json failed")
	}

	log.Printf("Event: %+v", payload)

	task := asynq.NewTask(workers.ChangeCreated, payload)

	client.Enqueue(task)

	if err != nil {
		println(err)
		// failed to send the task
		// do something with the error
	}
}
