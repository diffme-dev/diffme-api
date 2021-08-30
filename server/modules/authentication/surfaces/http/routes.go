package http

import (
	domain "diffme.dev/diffme-api/server/modules/authentication"
	"github.com/gofiber/fiber/v2"
)

func AuthenticationRoutes(f fiber.Router, authUseCases domain.UseCases) {
	auth := f.Group("/auth")

	// If don't pass in members it is like calling stuff on nil/null
	controller := &AuthenticationController{
		authUseCases: authUseCases,
	}

	auth.Post("/login", controller.Login)
}
