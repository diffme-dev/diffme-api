package http

import (
	"diffme.dev/diffme-api/server/core"
	errors2 "diffme.dev/diffme-api/server/core/errors"
	domain "diffme.dev/diffme-api/server/modules/authentication"
	"github.com/gofiber/fiber/v2"
)

type AuthenticationController struct {
	authUseCases domain.UseCases
}

type LoginParams struct {
	Type     string `json:"type"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (e *AuthenticationController) Login(c *fiber.Ctx) error {

	bodyParams := new(LoginParams)

	if err := c.BodyParser(bodyParams); err != nil {
		return err
	}

	errors := core.ValidateStruct(bodyParams)

	if len(errors) > 0 {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid json.")
	}

	switch bodyParams.Type {
	case "email":
		{
			user, err := e.authUseCases.EmailLogin(bodyParams.Email, bodyParams.Password)

			if err != nil {
				return errors2.NewApiError(c, err, fiber.StatusBadRequest, struct{}{})
			}

			return c.Status(200).JSON(*user)
		}
	default:
		{
			return errors2.NewApiError(
				c,
				fiber.NewError(fiber.StatusBadRequest, "Missing valid login type."),
				fiber.StatusBadRequest, struct{}{},
			)
		}

	}

}
