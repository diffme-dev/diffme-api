package http

import (
	domain "diffme.dev/diffme-api/internal/modules/users"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(f fiber.Router, userRepo domain.UserRepository, userUseCases domain.UserUseCases) {
	users := f.Group("/users")

	// If don't pass in members it is like calling stuff on nil/null
	controller := &UserController{
		userRepo: userRepo,
	}

	users.Get("/me", controller.GetMyUser)
	//snapshots.Get("/:id", controller.GetSnapshotByID)
	users.Post("/", controller.CreateUser)
}
