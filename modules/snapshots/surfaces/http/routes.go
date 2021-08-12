package http

import (
	domain "diffme.dev/diffme-api/modules/snapshots"
	"github.com/gofiber/fiber/v2"
)

type snapshotController struct {
	snapshotRepo     domain.SnapshotRepo
	snapshotUseCases domain.SnapshotUseCases
}

func SnapshotController(f fiber.Router, snapshotRepo domain.SnapshotRepo, snapshotUseCases domain.SnapshotUseCases) {
	snapshots := f.Group("/snapshots")

	controller := &snapshotController{
		snapshotRepo,
		snapshotUseCases,
	}

	snapshots.Get("/:id", controller.GetSnapshotByID)
	snapshots.Post("/", controller.CreateSnapshot)
}
