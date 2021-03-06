package http

import (
	"diffme.dev/diffme-api/server/core"
	errors2 "diffme.dev/diffme-api/server/core/errors"
	"diffme.dev/diffme-api/server/core/interfaces"
	domain "diffme.dev/diffme-api/server/modules/users"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

type UserController struct {
	userRepo     domain.UserRepository
	userUseCases domain.UserUseCases
	authProvider interfaces.AuthProvider
}

func (e *UserController) GetMyUser(c *fiber.Ctx) error {

	ctx := c.UserContext()

	user := ctx.Value("user").(*domain.User)

	return c.JSON(user)
}

func (e *UserController) CreateUser(c *fiber.Ctx) error {

	userParams := new(domain.CreateUserParams)

	if err := c.BodyParser(userParams); err != nil {
		return err
	}

	errors := core.ValidateStruct(userParams)

	if len(errors) > 0 {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid json.")
	}

	log.Printf("auth provider %+v", e.authProvider)

	userAuth, err := e.authProvider.FindOrCreate(userParams.Email, interfaces.CreateUserParams{
		Name:        userParams.Name,
		Email:       userParams.Email,
		Password:    userParams.Password,
		PhoneNumber: userParams.PhoneNumber,
	})

	if err != nil {
		fmt.Printf("Error: %+v", err)
		return errors2.NewApiError(c, err, fiber.StatusBadRequest, struct{}{})
	}

	// update the auth on the user to be the firebase auth
	userParams.Auth = &domain.UserAuthProvider{
		Provider:       userAuth.Provider,
		ProviderUserId: userAuth.ProviderUserId,
	}

	user, err := e.userUseCases.CreateUser(*userParams)

	if err != nil {
		return errors2.NewApiError(c, err, fiber.StatusBadRequest, struct{}{})
	}

	return c.JSON(user)
}
