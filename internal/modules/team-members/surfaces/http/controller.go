package http

import (
	"github.com/gofiber/fiber/v2"
)

type TeamMembersController struct {
	teamUseCases
}

// TODO: fix this route
func (e *TeamMembersController) InviteTeamMember(c *fiber.Ctx) error {
	userId := c.Params("id")

	data, err := e.userRepo.FindById(userId)

	if err != nil {
		return err
	}

	return c.JSON(data)
}

func (e *TeamMembersController) GetMyTeamMember(c *fiber.Ctx) error {
	userId := c.Params("id")

	data, err := e.userRepo.FindById(userId)

	if err != nil {
		return err
	}

	return c.JSON(data)
}
