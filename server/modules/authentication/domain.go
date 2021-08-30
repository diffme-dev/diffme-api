package authentication

import (
	UserDomain "diffme.dev/diffme-api/server/modules/users"
)

type UseCases interface {
	EmailLogin(email string, password string) (*UserDomain.User, error)
}
