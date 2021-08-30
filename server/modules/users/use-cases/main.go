package use_cases

import (
	UserDomain "diffme.dev/diffme-api/server/modules/users"
)

type UserUseCases struct {
	userRepo UserDomain.UserRepository
}

func NewUserUseCases(userRepo UserDomain.UserRepository) UserDomain.UserUseCases {
	return &UserUseCases{
		userRepo: userRepo,
	}
}
