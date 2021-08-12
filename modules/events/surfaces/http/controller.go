package http

import "github.com/gofiber/fiber/v2"

type SomeStruct struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func (e *EventHandler) GetEvents(c *fiber.Ctx) error {
	data := SomeStruct{
		Name: "Grame",
		Age:  20,
	}

	return c.JSON(data)
}

func (e *EventHandler) SearchEvents(c *fiber.Ctx) error {
	data := SomeStruct{
		Name: "Grame",
		Age:  20,
	}

	return c.JSON(data)
}

func (e *EventHandler) GetEventById(c *fiber.Ctx) error {
	data := SomeStruct{
		Name: "Grame",
		Age:  20,
	}

	return c.JSON(data)
}

func (e *EventHandler) CreateEvent(c *fiber.Ctx) error {
	data := SomeStruct{
		Name: "Grame",
		Age:  20,
	}

	return c.JSON(data)
}
