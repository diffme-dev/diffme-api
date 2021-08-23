package http

import (
	"diffme.dev/diffme-api/internal/core"
	domain "diffme.dev/diffme-api/internal/modules/users"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userRepo     domain.UserRepository
	userUseCases domain.UserUseCases
}

func (e *UserController) GetMyUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	data, err := e.userRepo.FindById(userId)

	if err != nil {
		return err
	}

	return c.JSON(data)
}

func (e *UserController) CreateUser(c *fiber.Ctx) error {

	userParams := new(domain.User)

	if err := c.BodyParser(userParams); err != nil {
		return err
	}

	errors := core.ValidateStruct(userParams)

	if len(errors) > 0 {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid json.")
	}

	// TODO: bunch of firebase / auth things need to go here...
	// TODO: figure out the right abstraction

	snapshot, err := e.userUseCases.CreateUser(*userParams)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(snapshot)
}
