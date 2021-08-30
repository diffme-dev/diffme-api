package infra

import (
	"context"
	config2 "diffme.dev/diffme-api/config"
	"diffme.dev/diffme-api/server/core/interfaces"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"strings"
)

type FirebaseProvider struct {
	client *firebase.App
}

func NewFirebaseProvider() interfaces.AuthProvider {
	//`{"private_key": ""}`
	config := config2.GetConfig()
	privateKey := strings.Replace(config.FirebasePrivateKey, "\n", "", 0)

	json := fmt.Sprintf(
		`{"private_key": "%s", "client_email": "%s", "type": "%s" }`,
		privateKey,
		config.FirebaseClientEmail,
		"service_account",
	)

	opt := option.WithCredentialsJSON([]byte(json))

	firebaseConfig := &firebase.Config{
		ProjectID: config2.GetConfig().FirebaseProjectId,
	}

	app, err := firebase.NewApp(context.Background(), firebaseConfig, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return &FirebaseProvider{
		client: app,
	}
}

func (p *FirebaseProvider) toAuthUser(firebaseUser *auth.UserRecord) *interfaces.AuthUser {
	return &interfaces.AuthUser{
		Name:           firebaseUser.DisplayName,
		ProfileUrl:     firebaseUser.PhotoURL,
		PhoneNumber:    firebaseUser.PhoneNumber,
		Email:          firebaseUser.Email,
		ProviderUserId: firebaseUser.UID,
		Provider:       "firebase",
	}
}

func (p *FirebaseProvider) GetName() string {
	return "firebase"
}

func (p *FirebaseProvider) GetAuthToken(uid string) (string, error) {
	c, err := p.client.Auth(context.Background())

	if err != nil {
		return "", err
	}

	tok, err := c.CustomToken(context.Background(), uid)

	if err != nil {
		return "", err
	}

	return tok, nil
}

func (p *FirebaseProvider) VerifyToken(token string) (*string, error) {
	c, err := p.client.Auth(context.Background())

	if err != nil {
		return nil, err
	}

	tok, err := c.VerifyIDToken(context.Background(), token)

	if err != nil {
		return nil, err
	}

	return &tok.UID, nil
}

func (p *FirebaseProvider) Create(createUser interfaces.CreateUserParams) (interfaces.AuthUser, error) {
	c, err := p.client.Auth(context.Background())

	if err != nil {
		return interfaces.AuthUser{}, err
	}

	params := (&auth.UserToCreate{}).
		Email(createUser.Email).
		EmailVerified(false).
		Password(createUser.Password).
		DisplayName(createUser.Name).
		Disabled(false)

	if createUser.ProfileUrl != "" {
		params.PhotoURL(createUser.ProfileUrl)
	}

	if createUser.PhoneNumber != "" {
		params.PhoneNumber(createUser.PhoneNumber)
	}

	firebaseUser, err := c.CreateUser(context.Background(), params)

	if err != nil {
		return interfaces.AuthUser{}, err
	}

	return *p.toAuthUser(firebaseUser), nil
}

// FIXME: atm this can only be done clientside which is kinda a bummer and ties
// the project pretty tightly to firebase...
func (p *FirebaseProvider) CheckPasswordForEmail(email string, password string) (*interfaces.AuthUser, error) {

	c, err := p.client.Auth(context.Background())

	if err != nil {
		return nil, err
	}

	_, err = c.GetUserByEmail(context.Background(), email)

	if err != nil {
		return nil, err
	}

	return nil, err
}

func (p *FirebaseProvider) FindByEmail(email string) (*interfaces.AuthUser, error) {
	c, err := p.client.Auth(context.Background())

	if err != nil {
		return nil, err
	}

	firebaseUser, err := c.GetUserByEmail(context.Background(), email)

	fmt.Printf("FB Error: %s\n", err)

	if err != nil {
		return nil, err
	}

	return p.toAuthUser(firebaseUser), nil
}

func (p *FirebaseProvider) FindOrCreate(email string, createUser interfaces.CreateUserParams) (interfaces.AuthUser, error) {

	// Note: errors if the user does not exist, so we
	// don't handle the error case bc it just creates a user in that case
	firebaseUser, _ := p.FindByEmail(email)

	fmt.Printf("firebase user: %+v", firebaseUser)

	if firebaseUser != nil {
		return *firebaseUser, nil
	}

	newUser, err := p.Create(createUser)

	fmt.Printf("firebase user: %+v", newUser)

	if err != nil {
		return interfaces.AuthUser{}, err
	}

	return newUser, nil
}
