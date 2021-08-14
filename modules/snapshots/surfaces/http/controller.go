package http

import (
	domain "diffme.dev/diffme-api/modules/snapshots"
	"diffme.dev/diffme-api/modules/snapshots/services"
	"github.com/RichardKnop/machinery/v1"
	valid "github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type SnapshotController struct {
	snapshotRepo     domain.SnapshotRepo
	snapshotUseCases domain.SnapshotUseCases
	taskserver       machinery.Server
}

func (e *SnapshotController) GetSnapshotByID(c *fiber.Ctx) error {
	snapshotID := c.Params("id")

	println("snapshot ID: ", snapshotID)

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

	_, err := valid.ValidateStruct(snapshotParams)

	if err != nil {
		return err
	}

	lastSnapshot, err := e.snapshotRepo.FindMostRecentByReference(snapshotParams.ReferenceID)

	snapshot, err := e.snapshotRepo.CreateSnapshot(*snapshotParams)

	services.SnapshotCreated(&e.taskserver, lastSnapshot, snapshot)

	return c.JSON(snapshot)
}
