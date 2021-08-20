package users

import (
	"github.com/wI2L/jsondiff"
	"time"
)

type Diff jsondiff.Operation

type UserAuthProvider struct {
	Provider string `json:"provider"`
	Uid      string `json:"uid`
}

type User struct {
	Id        string           `json:"id"`
	Name      string           `json:"name"`
	Auth      UserAuthProvider `json:"auth"`
	UpdatedAt time.Time        `json:"updated_at"`
	CreatedAt time.Time        `json:"created_at"`
}

type UserRepository interface {
	FindById(id string) (User, error)
	Update(userId string, user User) (User, error)
	Create(user User) (User, error)
}

type UserUseCases interface {
	GetUserById(userId string) (User, error)
	CreateUser(user User) (User, error)
}
