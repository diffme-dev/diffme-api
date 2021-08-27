package use_cases

import (
	domain "diffme.dev/diffme-api/internal/modules/users"
)

func (u *UserUseCases) CreateUser(userParams domain.CreateUserParams) (*domain.User, error) {
	return u.userRepo.Create(userParams)
}
