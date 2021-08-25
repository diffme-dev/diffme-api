package http

import (
	"diffme.dev/diffme-api/internal/core/interfaces"
	domain "diffme.dev/diffme-api/internal/modules/users"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(f fiber.Router, userRepo domain.UserRepository, userUseCases domain.UserUseCases, authProvider interfaces.AuthProvider) {
	users := f.Group("/users")

	// If don't pass in members it is like calling stuff on nil/null
	controller := &UserController{
		userRepo:     userRepo,
		authProvider: authProvider,
		userUseCases: userUseCases,
	}

	// user routes
	users.Get("/me", controller.GetMyUser)
	//snapshots.Get("/:id", controller.GetSnapshotByID)
	users.Post("/", controller.CreateUser)
}
