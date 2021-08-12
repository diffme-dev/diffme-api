package http

import (
	domain "diffme.dev/diffme-api/modules/snapshots"
	valid "github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (e *snapshotController) GetSnapshotByID(c *fiber.Ctx) error {
	snapshotID := c.Params("id")

	println("snapshot ID: ", snapshotID)

	data, err := e.snapshotRepo.FindByID(snapshotID)

	if err != nil {
		return err
	}

	return c.JSON(data)
}

func (e *snapshotController) CreateSnapshot(c *fiber.Ctx) error {

	snapshotParams := new(domain.CreateSnapshotParams)

	if err := c.BodyParser(snapshotParams); err != nil {
		return err
	}

	_, err := valid.ValidateStruct(snapshotParams)

	if err != nil {
		return err
	}

	snapshot, err := e.snapshotRepo.CreateSnapshot(*snapshotParams)

	// TODO: fire off previous and the current to kafka to process and store
	// the diffs
	log.Printf("Snapshot: %+v", snapshot)

	return c.JSON(snapshot)
}
