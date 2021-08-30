package use_cases

import (
	UserDomain "diffme.dev/diffme-api/server/modules/users"
)

// TODO: fill this in atm this project is kinda tightly coupled to firebase...
func (e *AuthUseCases) EmailLogin(email string, password string) (*UserDomain.User, error) {
	return nil, nil
}
