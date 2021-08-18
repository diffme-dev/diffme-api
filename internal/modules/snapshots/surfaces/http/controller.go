package http

import (
	"diffme.dev/diffme-api/internal/core"
	"diffme.dev/diffme-api/internal/modules/snapshots"
	"github.com/gofiber/fiber/v2"
)

type SnapshotController struct {
	snapshotRepo     domain.SnapshotRepo
	snapshotUseCases domain.SnapshotUseCases
}

func (e *SnapshotController) GetSnapshotByID(c *fiber.Ctx) error {
	snapshotID := c.Params("id")

	data, err := e.snapshotRepo.FindByID(snapshotID)

	if err != nil {
		return err
	}

	return c.JSON(data)
}

func (e *SnapshotController) CreateSnapshot(c *fiber.Ctx) error {

	snapshotParams := new(domain.CreateSnapshotParams)

	if err := c.BodyParser(snapshotParams); err != nil {
		return err
	}

	errors := core.ValidateStruct(snapshotParams)

	if len(errors) > 0 {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid json.")
	}

	snapshot, err := e.snapshotUseCases.CreateSnapshot(*snapshotParams)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(snapshot)
}
