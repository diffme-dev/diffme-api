package UseCases

import (
	domain "diffme.dev/diffme-api/internal/modules/users"
)

func (u *UserUseCases) GetUserById(userId string) (*domain.User, error) {
	return u.userRepo.FindById(userId)
}
