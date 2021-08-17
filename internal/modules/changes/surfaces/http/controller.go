package http

import (
	domain "diffme.dev/diffme-api/internal/modules/changes"
	"github.com/gofiber/fiber/v2"
)

type SomeStruct struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

type ChangeController struct {
	changeUseCases domain.ChangeUseCases
}

func (e *ChangeController) GetChanges(c *fiber.Ctx) error {

	e.changeUseCases.SearchChange("hi")

	return c.JSON(&SomeStruct{})
}

func (e *ChangeController) SearchChanges(c *fiber.Ctx) error {

	// editor
	// field name

	e.changeUseCases.SearchChange("hi")

	return c.JSON(&SomeStruct{})
}

func (e *ChangeController) GetChangeByReferenceID(c *fiber.Ctx) error {

	e.changeUseCases.SearchChange("hi")

	return c.JSON(&SomeStruct{})
}
