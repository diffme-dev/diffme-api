package main

import (
	"diffme.dev/diffme-api/cmd/server"
	"diffme.dev/diffme-api/cmd/workers"
	"diffme.dev/diffme-api/config"
	"diffme.dev/diffme-api/internal/core/infra"
	ChangeUseCases "diffme.dev/diffme-api/internal/modules/changes/UseCases"
	ChangeElastic "diffme.dev/diffme-api/internal/modules/changes/infra/elasticsearch"
	ChangeMongo "diffme.dev/diffme-api/internal/modules/changes/infra/mongo"
	SnapshotUseCases "diffme.dev/diffme-api/internal/modules/snapshots/UseCases"
	SnapshotMongo "diffme.dev/diffme-api/internal/modules/snapshots/infra/mongo"
	UserMongo "diffme.dev/diffme-api/internal/modules/users/infra/mongo"
	"diffme.dev/diffme-api/internal/shared/compression"
	"fmt"
	"log"
	"sync"
)

func main() {
	// loads the config and builds a singleton
	c := config.GetConfig()

	fmt.Printf("Config %+v", c)

	wg := new(sync.WaitGroup)

	wg.Add(2)

	// infra connections
	kafkaConsumer := infra.NewKafkaConsumer()
	kafkaProducer := infra.NewKafkaProducer()

	mongoClient, err := infra.NewBongoConnection()
	_, err = infra.NewRedisClient() // TODO:
	elasticClient, err := infra.NewElasticSearch()
	lz4Compression := compression.NewLZ4Compression()

	if err != nil {
		log.Fatal(err)
	}

	// auth provider
	authProvider := infra.NewFirebaseProvider()

	// repositories
	searchChangeRepo := ChangeElastic.NewElasticSearchChangeRepo(elasticClient)
	changeRepo := ChangeMongo.NewMongoChangeRepo(mongoClient)
	snapshotRepo := SnapshotMongo.NewMongoSnapshotRepo(mongoClient)
	userRepo := UserMongo.NewMongoUserRepo(mongoClient)

	// TODO:
	println(userRepo)

	// use cases
	changeUseCases := ChangeUseCases.NewChangeUseCase(changeRepo, searchChangeRepo)
	snapshotUseCases := SnapshotUseCases.NewSnapshotUseCases(snapshotRepo, lz4Compression)

	go func() {
		workerDeps := workers.NewWorkerDependencies(
			changeUseCases,
			snapshotRepo,
			snapshotUseCases,
			searchChangeRepo,
			kafkaConsumer,
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
			kafkaProducer,
			authProvider,
		)

		server.StartServer(serverDeps)

		wg.Done()
	}()

	wg.Wait()

	defer mongoClient.Session.Close()

}
