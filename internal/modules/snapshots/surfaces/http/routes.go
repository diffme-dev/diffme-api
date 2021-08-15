package http

import (
	"diffme.dev/diffme-api/internal/modules/snapshots"
	"github.com/RichardKnop/machinery/v1"
	"github.com/gofiber/fiber/v2"
)

func SnapshotRoutes(f fiber.Router, snapshotRepo domain.SnapshotRepo, snapshotUseCases domain.SnapshotUseCases, taskserver *machinery.Server) {
	snapshots := f.Group("/snapshots")

	// If don't pass in members it is like calling stuff on nil/null
	controller := &SnapshotController{
		snapshotRepo:     snapshotRepo,
		snapshotUseCases: snapshotUseCases,
		taskserver:       *taskserver,
	}

	snapshots.Get("/:id", controller.GetSnapshotByID)
	snapshots.Post("/", controller.CreateSnapshot)
}
