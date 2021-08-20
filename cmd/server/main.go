package server

import (
	"diffme.dev/diffme-api/internal/core/interfaces"
	"diffme.dev/diffme-api/internal/core/middleware"
	ChangeDomain "diffme.dev/diffme-api/internal/modules/changes"
	EventHTTP "diffme.dev/diffme-api/internal/modules/changes/surfaces/http"
	SnapshotDomain "diffme.dev/diffme-api/internal/modules/snapshots"
	SnapshotHTTP "diffme.dev/diffme-api/internal/modules/snapshots/surfaces/http"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type ServerDependencies struct {
	changeUseCases   ChangeDomain.ChangeUseCases
	snapshotRepo     SnapshotDomain.SnapshotRepo
	snapshotUseCases SnapshotDomain.SnapshotUseCases
	searchChangeRepo ChangeDomain.SearchChangeRepository
	producer         *kafka.Producer
	authProvider     interfaces.AuthProvider
}

func NewServerDependencies(
	changeUseCases ChangeDomain.ChangeUseCases,
	snapshotRepo SnapshotDomain.SnapshotRepo,
	snapshotUseCases SnapshotDomain.SnapshotUseCases,
	searchChangeRepo ChangeDomain.SearchChangeRepository,
	producer *kafka.Producer,
	authProvider interfaces.AuthProvider,
) ServerDependencies {
	return ServerDependencies{
		changeUseCases:   changeUseCases,
		snapshotRepo:     snapshotRepo,
		snapshotUseCases: snapshotUseCases,
		searchChangeRepo: searchChangeRepo,
		producer:         producer,
		authProvider:     authProvider,
	}
}

func StartServer(deps ServerDependencies) {
	println("[starting server]")

	// Fiber instance
	app := fiber.New()

	// logger
	app.Use(logger.New())

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// compress
	app.Use(compress.New())

	// cache
	app.Use(cache.New())

	// request ID
	app.Use(requestid.New())

	app.Use(middleware.AuthMiddleware(deps.authProvider))

	v1 := app.Group("/v1")

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

func addRoutes(app fiber.Router, deps ServerDependencies) {

	EventHTTP.ChangeRoutes(app, deps.changeUseCases)
	SnapshotHTTP.SnapshotRoutes(app, deps.snapshotRepo, deps.snapshotUseCases)
}
