package http

import (
	"diffme.dev/diffme-api/internal/core"
	domain "diffme.dev/diffme-api/internal/modules/changes"
	"github.com/gofiber/fiber/v2"
	"log"
)

type SomeStruct struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

type ChangeController struct {
	changeUseCases domain.ChangeUseCases
}

func (e *ChangeController) SearchChanges(c *fiber.Ctx) error {

	search := new(domain.SearchRequest)

	if err := c.QueryParser(search); err != nil {
		return err
	}

	errors := core.ValidateStruct(search)

	if len(errors) > 0 {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid json.")
	}

	log.Printf("Query %+v", &search)

	searchChanges, err := e.changeUseCases.SearchChange(*search)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return c.JSON(searchChanges)
}

func (e *ChangeController) GetChangeByReferenceID(c *fiber.Ctx) error {

	referenceId := c.Params("id")

	changes, err := e.changeUseCases.FetchChangeForReferenceId(referenceId)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return c.JSON(changes)
}
