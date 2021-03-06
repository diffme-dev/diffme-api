package main

import (
	"diffme.dev/diffme-api/cmd/server"
	"diffme.dev/diffme-api/cmd/workers"
	"diffme.dev/diffme-api/config"
	infra2 "diffme.dev/diffme-api/server/core/infra"
	ApiKeyMongo "diffme.dev/diffme-api/server/modules/api-keys/infra/mongo"
	"diffme.dev/diffme-api/server/modules/authentication/use-cases"
	ChangeElastic "diffme.dev/diffme-api/server/modules/changes/infra/elasticsearch"
	ChangeMongo "diffme.dev/diffme-api/server/modules/changes/infra/mongo"
	ChangeUseCases "diffme.dev/diffme-api/server/modules/changes/use-cases"
	OrgMongo "diffme.dev/diffme-api/server/modules/organizations/infra/mongo"
	OrgUseCases "diffme.dev/diffme-api/server/modules/organizations/use-cases"
	SnapshotMongo "diffme.dev/diffme-api/server/modules/snapshots/infra/mongo"
	SnapshotUseCases "diffme.dev/diffme-api/server/modules/snapshots/use-cases"
	UserMongo "diffme.dev/diffme-api/server/modules/users/infra/mongo"
	UserUseCases "diffme.dev/diffme-api/server/modules/users/use-cases"
	"diffme.dev/diffme-api/server/shared/compression"
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
	kafkaConsumer := infra2.NewKafkaConsumer()
	kafkaProducer := infra2.NewKafkaProducer()

	mongoClient, err := infra2.NewBongoConnection()
	_, err = infra2.NewRedisClient() // TODO:
	elasticClient, err := infra2.NewElasticSearch()
	lz4Compression := compression.NewLZ4Compression()

	if err != nil {
		log.Fatal(err)
	}

	// auth provider
	authProvider := infra2.NewFirebaseProvider()

	// repositories
	searchChangeRepo := ChangeElastic.NewElasticSearchChangeRepo(elasticClient)
	changeRepo := ChangeMongo.NewMongoChangeRepo(mongoClient)
	snapshotRepo := SnapshotMongo.NewMongoSnapshotRepo(mongoClient)
	userRepo := UserMongo.NewMongoUserRepo(mongoClient)
	orgRepo := OrgMongo.NewMongoOrganizationRepo(mongoClient)
	apiKeyRepo := ApiKeyMongo.NewMongoApiKeyRepo(mongoClient)

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
			apiKeyRepo,
		)

		server.StartServer(serverDeps)

		wg.Done()
	}()

	wg.Wait()

	defer mongoClient.Session.Close()

}
