package main

import (
	"diffme.dev/diffme-api/cmd/server"
	"diffme.dev/diffme-api/cmd/workers"
	"diffme.dev/diffme-api/internal/core/infra"
	ChangeUseCases "diffme.dev/diffme-api/internal/modules/changes/UseCases"
	ChangeElastic "diffme.dev/diffme-api/internal/modules/changes/infra/elasticsearch"
	ChangeMongo "diffme.dev/diffme-api/internal/modules/changes/infra/mongo"
	SnapshotUseCases "diffme.dev/diffme-api/internal/modules/snapshots/UseCases"
	SnapshotMongo "diffme.dev/diffme-api/internal/modules/snapshots/infra/mongo"
	"diffme.dev/diffme-api/internal/shared/compression"
	"log"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(2)

	// infra connections
	mongoClient, err := infra.NewBongoConnection()
	//redisClient, err := Infra.NewRedisClient()
	elasticClient, err := infra.NewElasticSearch()
	lz4Compression := compression.NewLZ4Compression()

	if err != nil {
		log.Fatal(err)
	}

	searchChangeRepo := ChangeElastic.NewElasticSearchChangeRepo(elasticClient)
	changeRepo := ChangeMongo.NewMongoChangeRepo(mongoClient)
	snapshotRepo := SnapshotMongo.NewMongoSnapshotRepo(mongoClient)
	changeUseCases := ChangeUseCases.NewChangeUseCase(changeRepo, searchChangeRepo)
	snapshotUseCases := SnapshotUseCases.NewSnapshotUseCases(snapshotRepo, lz4Compression)

	go func() {
		workerDeps := workers.NewWorkerDependencies(
			changeUseCases,
			snapshotRepo,
			snapshotUseCases,
			searchChangeRepo,
		)

		workers.StartWorkers(workerDeps)

		wg.Done()
	}()

	go func() {
		serverDeps := server.NewServerDependencies(
			changeUseCases,
			snapshotRepo,
			snapshotUseCases,
			searchChangeRepo,
		)

		server.StartServer(serverDeps)

		wg.Done()
	}()

	wg.Wait()

	defer mongoClient.Session.Close()

}
