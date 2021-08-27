package main

import (
	"diffme.dev/diffme-api/cmd/server"
	"diffme.dev/diffme-api/cmd/workers"
	"diffme.dev/diffme-api/config"
	"diffme.dev/diffme-api/internal/core/infra"
	"diffme.dev/diffme-api/internal/modules/authentication/use-cases"
	ChangeElastic "diffme.dev/diffme-api/internal/modules/changes/infra/elasticsearch"
	ChangeMongo "diffme.dev/diffme-api/internal/modules/changes/infra/mongo"
	ChangeUseCases "diffme.dev/diffme-api/internal/modules/changes/use-cases"
	OrgUseCases "diffme.dev/diffme-api/internal/modules/organizations/UseCases"
	OrgMongo "diffme.dev/diffme-api/internal/modules/organizations/infra/mongo"
	SnapshotMongo "diffme.dev/diffme-api/internal/modules/snapshots/infra/mongo"
	SnapshotUseCases "diffme.dev/diffme-api/internal/modules/snapshots/use-cases"
	UserMongo "diffme.dev/diffme-api/internal/modules/users/infra/mongo"
	UserUseCases "diffme.dev/diffme-api/internal/modules/users/use-cases"
	"diffme.dev/diffme-api/internal/shared/compression"
	"log"
	"sync"
)

func main() {
	// loads the config and builds a singleton
	config.GetConfig()

	//fmt.Printf("Config %+v", c)

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
	orgRepo := OrgMongo.NewMongoOrganizationRepo(mongoClient)

	// TODO:
	println(userRepo)

	// use cases
	changeUseCases := ChangeUseCases.NewChangeUseCase(changeRepo, searchChangeRepo)
	snapshotUseCases := SnapshotUseCases.NewSnapshotUseCases(snapshotRepo, lz4Compression)
	orgUseCases := OrgUseCases.NewOrganizationUseCases(orgRepo)
	userUseCases := UserUseCases.NewUserUseCases(userRepo)
	authUseCases := use_cases.NewAuthenticationUseCases(userUseCases, authProvider)

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
			orgUseCases,
			snapshotRepo,
			snapshotUseCases,
			searchChangeRepo,
			userRepo,
			userUseCases,
			authUseCases,
			kafkaProducer,
			authProvider,
		)

		server.StartServer(serverDeps)

		wg.Done()
	}()

	wg.Wait()

	defer mongoClient.Session.Close()

}
