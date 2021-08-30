package http

import (
	domain "diffme.dev/diffme-api/server/modules/team-members"
	"github.com/gofiber/fiber/v2"
)

type TeamMembersController struct {
	teamUseCases domain.TeamMemberUseCases
}

func (e *TeamMembersController) InviteTeamMember(c *fiber.Ctx) error {

	data, err := e.teamUseCases.InviteTeamMember(domain.TeamMember{
		// TODO:
	})

	if err != nil {
		return err
	}

	return c.JSON(data)
}

func (e *TeamMembersController) GetMyTeamMember(c *fiber.Ctx) error {

	// TODO: fill this stuff in
	//c.JSON(data)
	return nil
}
