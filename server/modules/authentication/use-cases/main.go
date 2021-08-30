package use_cases

import (
	"diffme.dev/diffme-api/server/core/interfaces"
	"diffme.dev/diffme-api/server/modules/authentication"
	UserDomain "diffme.dev/diffme-api/server/modules/users"
)

type AuthUseCases struct {
	userUseCases UserDomain.UserUseCases
	authProvider interfaces.AuthProvider
}

func NewAuthenticationUseCases(
	userUseCases UserDomain.UserUseCases,
	authProvider interfaces.AuthProvider,
) authentication.UseCases {
	return &AuthUseCases{
		userUseCases: userUseCases,
		authProvider: authProvider,
	}
}
