package main

import (
	"diffme.dev/diffme-api/internal/loaders"
	EventDomain "diffme.dev/diffme-api/modules/events"
	EventUseCases "diffme.dev/diffme-api/modules/events/UseCases"
	EventPostgres "diffme.dev/diffme-api/modules/events/infra/postgres"
	EventHTTP "diffme.dev/diffme-api/modules/events/surfaces/http"
	SnapshotDomain "diffme.dev/diffme-api/modules/snapshots"
	SnapshotUseCases "diffme.dev/diffme-api/modules/snapshots/UseCases"
	SnapshotMongo "diffme.dev/diffme-api/modules/snapshots/infra/mongo"
	SnapshotHTTP "diffme.dev/diffme-api/modules/snapshots/surfaces/http"
	"github.com/go-bongo/bongo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

var (
	mongoConfig = &bongo.Config{
		ConnectionString: "localhost",
		Database:         "diffme",
	}
)

type dependencies struct {
	eventUseCases    EventDomain.EventUseCases
	snapshotRepo     SnapshotDomain.SnapshotRepo
	snapshotUseCases SnapshotDomain.SnapshotUseCases
}

func main() {
	// Fiber instance
	app := fiber.New()
	app.Use(logger.New())
	v1 := app.Group("/v1")

	// how do I share this connection"
	mongoDB, err := bongo.Connect(mongoConfig)

	if err != nil {
		log.Fatal(err)
	}

	db := loaders.NewPostgresConnection()

	eventRepo := EventPostgres.NewPostgresEventRepository(db)
	snapshotRepo := SnapshotMongo.NewMongoSnapshotRepo(mongoDB)
	eventUseCases := EventUseCases.NewEventUseCase(eventRepo)
	snapshotUseCases := SnapshotUseCases.NewSnapshotUseCases(snapshotRepo)

	deps := dependencies{
		eventUseCases:    eventUseCases,
		snapshotRepo:     snapshotRepo,
		snapshotUseCases: snapshotUseCases,
	}

	addRoutes(v1, deps)

	// start server
	app.Listen(":3001")
}

func addRoutes(route fiber.Router, deps dependencies) {
	EventHTTP.Add(route, deps.eventUseCases)
	SnapshotHTTP.SnapshotController(route, deps.snapshotRepo, deps.snapshotUseCases)
}
