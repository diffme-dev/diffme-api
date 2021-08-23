package http

import (
	"diffme.dev/diffme-api/internal/core"
	Errors "diffme.dev/diffme-api/internal/core/errors"
	domain "diffme.dev/diffme-api/internal/modules/organizations"
	"github.com/gofiber/fiber/v2"
)

type OrgController struct {
	orgRepo     domain.OrganizationRepository
	orgUseCases domain.OrganizationUseCases
}

func (e *OrgController) GetMyActiveOrganization(c *fiber.Ctx) error {
	userId := c.Params("id")

	data, err := e.orgRepo.FindById(userId)

	if err != nil {
		apiErr := Errors.ApiError{
			Message: err.Error(),
			StatusCode: fiber.StatusBadRequest,
		}

		return c.Status(apiErr.StatusCode).JSON(apiErr)
	}

	return c.JSON(data)
}

func (e *OrgController) CreateOrganization(c *fiber.Ctx) error {

	userParams := new(domain.Organization)

	if err := c.BodyParser(userParams); err != nil {
		return err
	}

	errors := core.ValidateStruct(userParams)

	if len(errors) > 0 {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid json.")
	}

	snapshot, err := e.orgUseCases.Create(*userParams)

	if err != nil {
		apiErr := Errors.ApiError{
			Message: err.Error(),
			StatusCode: fiber.StatusBadRequest,
		}

		return c.Status(apiErr.StatusCode).JSON(apiErr)
	}

	return c.JSON(snapshot)
}
