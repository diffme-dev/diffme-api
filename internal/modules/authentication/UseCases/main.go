package UseCases

import (
	"diffme.dev/diffme-api/internal/core/interfaces"
	"diffme.dev/diffme-api/internal/modules/authentication"
	UserDomain "diffme.dev/diffme-api/internal/modules/users"
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
