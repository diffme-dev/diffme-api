package http

import (
	"github.com/gofiber/fiber/v2"
)

func TeamMemberRoutes(f fiber.Router) {
	users := f.Group("/team-members")

	// If don't pass in members it is like calling stuff on nil/null
	controller := &TeamMembersController{}

	users.Post("/invite", controller.InviteTeamMember)
}
