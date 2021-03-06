package use_cases

import (
	domain "diffme.dev/diffme-api/server/modules/users"
)

func (u *UserUseCases) GetUserById(userId string) (*domain.User, error) {
	return u.userRepo.FindById(userId)
}
