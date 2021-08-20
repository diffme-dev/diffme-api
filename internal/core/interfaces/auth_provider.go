package interfaces

type AuthProvider interface {
	GetName() string
	GetAuthToken(uid string) (string, error)
	VerifyToken(token string) (string, error)
}
