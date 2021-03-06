package http

import (
	"diffme.dev/diffme-api/server/core/middleware"
	"diffme.dev/diffme-api/server/modules/snapshots"
	"github.com/gofiber/fiber/v2"
)

func SnapshotRoutes(f fiber.Router, snapshotRepo domain.SnapshotRepo, snapshotUseCases domain.SnapshotUseCases) {
	snapshots := f.Group("/snapshots")

	// If don't pass in members it is like calling stuff on nil/null
	controller := &SnapshotController{
		snapshotRepo:     snapshotRepo,
		snapshotUseCases: snapshotUseCases,
	}

	snapshots.Get("/reference/:reference_id", middleware.AuthRequired, controller.GetLatestSnapshotForReference)
	//snapshots.Get("/:id", controller.GetSnapshotByID)
	snapshots.Post("/", middleware.AuthRequired, controller.CreateSnapshot)
}
