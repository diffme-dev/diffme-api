package interfaces

type CreateUserParams struct {
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	ProfileUrl  string
}

type AuthUser struct {
	Name           string
	Email          string
	PhoneNumber    string
	ProfileUrl     string
	Provider       string
	ProviderUserId string
}

type AuthProvider interface {
	GetName() string
	GetAuthToken(uid string) (string, error)
	VerifyToken(token string) (*string, error)
	CheckPasswordForEmail(email string, password string) (*AuthUser, error)
	Create(user CreateUserParams) (AuthUser, error)
	FindOrCreate(email string, user CreateUserParams) (AuthUser, error)
	FindByEmail(email string) (*AuthUser, error)
}
