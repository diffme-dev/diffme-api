package workers

import (
	Infra "diffme.dev/diffme-api/internal/infra"
	ChangeAsynq "diffme.dev/diffme-api/internal/modules/changes/surfaces/asynq"
	"log"
)

var (
	SnapshotCreated = "SnapshotCreated"
	ChangeCreated   = "ChangeCreated"
)

func StartWorkers() {
	println("[starting workers]")

	server, mux := Infra.NewAsynqServer()

	mux.HandleFunc(SnapshotCreated, ChangeAsynq.CreateChangeHandler)
	mux.HandleFunc(ChangeCreated, ChangeAsynq.CreateSearchableChangeHandler)

	if err := server.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

}
