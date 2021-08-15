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
	EventUseCase domain.ChangeUseCases
}

func (e *ChangeController) GetEvents(c *fiber.Ctx) error {
	data := SomeStruct{
		Name: "Grame",
		Age:  20,
	}

	return c.JSON(data)
}

func (e *ChangeController) SearchEvents(c *fiber.Ctx) error {
	data := SomeStruct{
		Name: "Grame",
		Age:  20,
	}

	return c.JSON(data)
}

func (e *ChangeController) GetEventById(c *fiber.Ctx) error {
	data := SomeStruct{
		Name: "Grame",
		Age:  20,
	}

	return c.JSON(data)
}

func (e *ChangeController) CreateEvent(c *fiber.Ctx) error {
	data := SomeStruct{
		Name: "Grame",
		Age:  20,
	}

	return c.JSON(data)
}
