package workers

import (
	Infra "diffme.dev/diffme-api/internal/core/infra"
	ChangeDomain "diffme.dev/diffme-api/internal/modules/changes"
	ChangeAsynq "diffme.dev/diffme-api/internal/modules/changes/surfaces/asynq"
	SnapshotDomain "diffme.dev/diffme-api/internal/modules/snapshots"
	"log"
)

var (
	SnapshotCreated = "SnapshotCreated"
	ChangeCreated   = "ChangeCreated"
)

type WorkerDependencies struct {
	changeUseCases   ChangeDomain.ChangeUseCases
	snapshotRepo     SnapshotDomain.SnapshotRepo
	snapshotUseCases SnapshotDomain.SnapshotUseCases
	searchChangeRepo ChangeDomain.SearchChangeRepository
}

func NewWorkerDependencies(
	changeUseCases ChangeDomain.ChangeUseCases,
	snapshotRepo SnapshotDomain.SnapshotRepo,
	snapshotUseCases SnapshotDomain.SnapshotUseCases,
	searchChangeRepo ChangeDomain.SearchChangeRepository,
) WorkerDependencies {
	return WorkerDependencies{
		changeUseCases:   changeUseCases,
		snapshotRepo:     snapshotRepo,
		snapshotUseCases: snapshotUseCases,
		searchChangeRepo: searchChangeRepo,
	}
}

func StartWorkers(deps WorkerDependencies) {
	println("[starting workers]")

	server, mux := Infra.NewAsynqServer()

	changeAsynqSurface := ChangeAsynq.NewChangeAsnqSurface(deps.changeUseCases)

	mux.HandleFunc(SnapshotCreated, changeAsynqSurface.CreateChangeHandler)
	mux.HandleFunc(ChangeCreated, changeAsynqSurface.CreateSearchableChangeHandler)

	if err := server.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

}
