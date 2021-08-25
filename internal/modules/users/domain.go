package users

import (
	"github.com/wI2L/jsondiff"
	"time"
)

type Diff jsondiff.Operation

type UserAuthProvider struct {
	Provider       string `json:"provider"`
	ProviderUserId string `json:"provider_user_id""`
}

type CreateUserParams struct {
	Name        string            `json:"name"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	PhoneNumber string            `json:"phone_number"`
	Email       string            `json:"email"`
	ProfileUrl  string            `json:"profile_url"`
	Password    string            `json:"password"`
	Auth        *UserAuthProvider `json:"auth"`
}

type User struct {
	Id          string           `json:"id"`
	Name        string           `json:"name"`
	FirstName   string           `json:"first_name"`
	LastName    string           `json:"last_name"`
	PhoneNumber string           `json:"phone_number"`
	Email       string           `json:"email"`
	ProfileUrl  string           `json:"profile_url"`
	Auth        UserAuthProvider `json:"auth"`
	UpdatedAt   time.Time        `json:"updated_at"`
	CreatedAt   time.Time        `json:"created_at"`
}

type UserRepository interface {
	FindById(id string) (*User, error)
	Update(userId string, user User) (*User, error)
	Create(user CreateUserParams) (*User, error)
}

type UserUseCases interface {
	GetUserById(userId string) (*User, error)
	CreateUser(user CreateUserParams) (*User, error)
}
