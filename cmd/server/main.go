package server

import (
	Infra "diffme.dev/diffme-api/internal/infra"
	ChangeDomain "diffme.dev/diffme-api/internal/modules/changes"
	ChangeUseCases "diffme.dev/diffme-api/internal/modules/changes/UseCases"
	ChangeElastic "diffme.dev/diffme-api/internal/modules/changes/infra/elasticsearch"
	ChangeMongo "diffme.dev/diffme-api/internal/modules/changes/infra/mongo"
	EventHTTP "diffme.dev/diffme-api/internal/modules/changes/surfaces/http"
	SnapshotDomain "diffme.dev/diffme-api/internal/modules/snapshots"
	SnapshotUseCases "diffme.dev/diffme-api/internal/modules/snapshots/UseCases"
	SnapshotMongo "diffme.dev/diffme-api/internal/modules/snapshots/infra/mongo"
	SnapshotHTTP "diffme.dev/diffme-api/internal/modules/snapshots/surfaces/http"
	"diffme.dev/diffme-api/internal/shared/compression"
	"github.com/RichardKnop/machinery/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

type dependencies struct {
	changeUseCases ChangeDomain.ChangeUseCases
	// TODO: remove
	snapshotRepo     SnapshotDomain.SnapshotRepo
	snapshotUseCases SnapshotDomain.SnapshotUseCases
	taskserver       machinery.Server
}

func StartServer() {
	println("[starting server]")

	// Fiber instance
	app := fiber.New()
	app.Use(logger.New())
	v1 := app.Group("/v1")

	// infra connections
	mongoClient, err := Infra.NewMongoConnection()
	redisClient, err := Infra.NewRedisClient()
	machineryClient, err := Infra.NewMachinery()
	elasticClient, err := Infra.NewElasticSearch()

	if err != nil {
		log.Fatal(err)
	}

	println(redisClient)
	println(machineryClient)

	lz4Compression := compression.NewLZ4Compression()

	searchChangeRepo := ChangeElastic.NewElasticSearchChangeRepo(elasticClient)
	changeRepo := ChangeMongo.NewMongoChangeRepo(mongoClient)
	snapshotRepo := SnapshotMongo.NewMongoSnapshotRepo(mongoClient)
	changeUseCases := ChangeUseCases.NewChangeUseCase(changeRepo)
	snapshotUseCases := SnapshotUseCases.NewSnapshotUseCases(snapshotRepo, machineryClient, lz4Compression)

	// TODO: use this
	println(searchChangeRepo)

	deps := dependencies{
		changeUseCases:   changeUseCases,
		snapshotRepo:     snapshotRepo,
		snapshotUseCases: snapshotUseCases,
		taskserver:       *machineryClient,
	}

	addRoutes(v1, deps)

	println("server listening on :3001")

	// start server
	app.Listen(":3001")
	//defer func() {
	//	if err = mongoClient.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
}

func addRoutes(route fiber.Router, deps dependencies) {
	EventHTTP.ChangeRoutes(route, deps.changeUseCases)
	SnapshotHTTP.SnapshotRoutes(route, deps.snapshotRepo, deps.snapshotUseCases, &deps.taskserver)
}
