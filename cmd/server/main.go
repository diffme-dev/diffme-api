package server

import (
	ChangeDomain "diffme.dev/diffme-api/internal/modules/changes"
	EventHTTP "diffme.dev/diffme-api/internal/modules/changes/surfaces/http"
	SnapshotDomain "diffme.dev/diffme-api/internal/modules/snapshots"
	SnapshotHTTP "diffme.dev/diffme-api/internal/modules/snapshots/surfaces/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type ServerDependencies struct {
	changeUseCases   ChangeDomain.ChangeUseCases
	snapshotRepo     SnapshotDomain.SnapshotRepo
	snapshotUseCases SnapshotDomain.SnapshotUseCases
	searchChangeRepo ChangeDomain.SearchChangeRepository
}

func NewServerDependencies(
	changeUseCases ChangeDomain.ChangeUseCases,
	snapshotRepo SnapshotDomain.SnapshotRepo,
	snapshotUseCases SnapshotDomain.SnapshotUseCases,
	searchChangeRepo ChangeDomain.SearchChangeRepository,
) ServerDependencies {
	return ServerDependencies{
		changeUseCases:   changeUseCases,
		snapshotRepo:     snapshotRepo,
		snapshotUseCases: snapshotUseCases,
		searchChangeRepo: searchChangeRepo,
	}
}

func StartServer(deps ServerDependencies) {
	println("[starting server]")

	// Fiber instance
	app := fiber.New()
	app.Use(logger.New())
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

func addRoutes(route fiber.Router, deps ServerDependencies) {
	EventHTTP.ChangeRoutes(route, deps.changeUseCases)
	SnapshotHTTP.SnapshotRoutes(route, deps.snapshotRepo, deps.snapshotUseCases)
}
