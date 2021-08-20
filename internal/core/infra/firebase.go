package infra

import (
	"context"
	"diffme.dev/diffme-api/internal/core/interfaces"
	firebase "firebase.google.com/go/v4"
	"log"
)

type FirebaseProvider struct {
	client *firebase.App
}

func NewFirebaseProvider() interfaces.AuthProvider {
	config := &firebase.Config{
		ProjectID: "my-project-id",
	}

	app, err := firebase.NewApp(context.Background(), config)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return &FirebaseProvider{
		client: app,
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

func (p *FirebaseProvider) VerifyToken(token string) (string, error) {
	c, err := p.client.Auth(context.Background())

	if err != nil {
		return "", err
	}

	tok, err := c.VerifyIDToken(context.Background(), token)

	if err != nil {
		return "", err
	}

	return tok.UID, nil
}
