package workers

import (
	Infra "diffme.dev/diffme-api/internal/core/infra"
	ChangeUseCases "diffme.dev/diffme-api/internal/modules/changes/UseCases"
	"diffme.dev/diffme-api/internal/modules/changes/infra/elasticsearch"
	ChangeMongo "diffme.dev/diffme-api/internal/modules/changes/infra/mongo"
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

	mongoClient, _ := Infra.NewMongoConnection()
	elasticClient, _ := Infra.NewElasticSearch()
	changeRepo := ChangeMongo.NewMongoChangeRepo(mongoClient)
	searchChangeRepo := elasticsearch.NewElasticSearchChangeRepo(elasticClient)

	changeUseCases := ChangeUseCases.NewChangeUseCase(changeRepo, searchChangeRepo)
	changeAsynqSurface := ChangeAsynq.NewChangeAsnqSurface(changeUseCases)

	mux.HandleFunc(SnapshotCreated, changeAsynqSurface.CreateChangeHandler)
	mux.HandleFunc(ChangeCreated, changeAsynqSurface.CreateSearchableChangeHandler)

	if err := server.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

}
