package http

import (
	domain "diffme.dev/diffme-api/internal/modules/organizations"
	"github.com/gofiber/fiber/v2"
)

func OrgRoutes(f fiber.Router, orgUseCases domain.OrganizationUseCases) {
	eventRoutes := f.Group("/organizations")

	handler := &OrgController{
		orgUseCases: orgUseCases,
	}

	eventRoutes.Get("/active", handler.GetMyActiveOrganization)
	eventRoutes.Post("/", handler.CreateOrganization)
}
