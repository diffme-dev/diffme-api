package http

import (
	"diffme.dev/diffme-api/internal/modules/snapshots"
	"github.com/RichardKnop/machinery/v1"
	valid "github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"log"
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

	isValid, err := valid.ValidateStruct(snapshotParams)

	if !isValid {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	log.Printf("snapshot: %+v", snapshotParams)

	snapshot, err := e.snapshotUseCases.CreateSnapshot(*snapshotParams)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(snapshot)
}
