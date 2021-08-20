package UseCases

import (
	domain "diffme.dev/diffme-api/internal/modules/users"
)

func (u *UserUseCases) CreateUser(userParams domain.User) (domain.User, error) {
	return u.userRepo.Create(userParams)
}
