package cmd

import (
	Infra "diffme.dev/diffme-api/internal/infra"
	EventDomain "diffme.dev/diffme-api/modules/events"
	EventUseCases "diffme.dev/diffme-api/modules/events/UseCases"
	EventPostgres "diffme.dev/diffme-api/modules/events/infra/postgres"
	EventHTTP "diffme.dev/diffme-api/modules/events/surfaces/http"
	SnapshotDomain "diffme.dev/diffme-api/modules/snapshots"
	SnapshotUseCases "diffme.dev/diffme-api/modules/snapshots/UseCases"
	SnapshotMongo "diffme.dev/diffme-api/modules/snapshots/infra/mongo"
	SnapshotHTTP "diffme.dev/diffme-api/modules/snapshots/surfaces/http"
	"github.com/RichardKnop/machinery/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

type dependencies struct {
	eventUseCases    EventDomain.EventUseCases
	snapshotRepo     SnapshotDomain.SnapshotRepo
	snapshotUseCases SnapshotDomain.SnapshotUseCases
	taskserver       machinery.Server
}

func StartServer() {
	// Fiber instance
	app := fiber.New()
	app.Use(logger.New())
	v1 := app.Group("/v1")

	// infra connections
	mongoClient, err := Infra.NewMongoConnection()
	redisClient, err := Infra.NewRedisClient()
	postgresClient := Infra.NewPostgresConnection()
	machineryClient, err := Infra.NewMachinery()

	if err != nil {
		log.Fatal(err)
	}

	println(redisClient)
	println(machineryClient)

	eventRepo := EventPostgres.NewPostgresEventRepository(postgresClient)
	snapshotRepo := SnapshotMongo.NewMongoSnapshotRepo(mongoClient)
	eventUseCases := EventUseCases.NewEventUseCase(eventRepo)
	snapshotUseCases := SnapshotUseCases.NewSnapshotUseCases(snapshotRepo)

	deps := dependencies{
		eventUseCases:    eventUseCases,
		snapshotRepo:     snapshotRepo,
		snapshotUseCases: snapshotUseCases,
		taskserver:       *machineryClient,
	}

	addRoutes(v1, deps)

	// start server
	app.Listen(":3001")

	//defer func() {
	//	if err = mongoClient.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
}

func addRoutes(route fiber.Router, deps dependencies) {
	EventHTTP.Add(route, deps.eventUseCases)
	SnapshotHTTP.SnapshotRoutes(route, deps.snapshotRepo, deps.snapshotUseCases, &deps.taskserver)
}
