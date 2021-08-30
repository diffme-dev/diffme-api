package http

import (
	"diffme.dev/diffme-api/server/core/middleware"
	domain "diffme.dev/diffme-api/server/modules/organizations"
	"github.com/gofiber/fiber/v2"
)

func OrgRoutes(f fiber.Router, orgUseCases domain.OrganizationUseCases) {
	eventRoutes := f.Group("/organizations")

	handler := &OrgController{
		orgUseCases: orgUseCases,
	}

	eventRoutes.Get("/active", middleware.AuthRequired, handler.GetMyActiveOrganization)
	eventRoutes.Post("/", middleware.AuthRequired, handler.CreateOrganization)
}
