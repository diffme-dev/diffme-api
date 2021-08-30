package middleware

import (
	"context"
	"diffme.dev/diffme-api/server/core"
	"diffme.dev/diffme-api/server/core/errors"
	"diffme.dev/diffme-api/server/core/interfaces"
	ApiKeyDomain "diffme.dev/diffme-api/server/modules/api-keys"
	UserDomain "diffme.dev/diffme-api/server/modules/users"
	errors2 "errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
)

type EntityType string

var (
	ApiKey EntityType = "api_key"
	User   EntityType = "user"
)

type JWTResponse struct {
	Type  EntityType `json:"type"`
	Value string     `json:"value"`
}

func getAuthedContext(
	isAuthed bool,
	entityType *EntityType,
	user *UserDomain.User,
	apiKey *ApiKeyDomain.ApiKey,
) context.Context {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "is_authenticated", isAuthed)
	ctx = context.WithValue(ctx, "user", user)
	ctx = context.WithValue(ctx, "entity_type", entityType)
	ctx = context.WithValue(ctx, "api_key", apiKey)

	return ctx
}

func verifyJWT(
	authProvider interfaces.AuthProvider,
	jwtToken string,
) *JWTResponse {
	uid, err := authProvider.VerifyToken(jwtToken)

	if err == nil && uid != nil {
		return &JWTResponse{
			Type:  User,
			Value: *uid,
		}
	}

	value, err := core.ParseToken(jwtToken)

	if err == nil && value != nil {
		return &JWTResponse{
			Type:  ApiKey,
			Value: string(value),
		}
	}

	return nil
}

func AuthMiddleware(
	authProvider interfaces.AuthProvider,
	userRepo UserDomain.UserRepository,
	apiKeyRepo ApiKeyDomain.ApiKeyRepository,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()

		header := c.Request().Header.Peek("Authorization")
		idToken := strings.TrimSpace(strings.Replace(string(header), "Bearer", "", 1))

		//fmt.Printf("Token: %s\n", idToken)

		if idToken != "" {
			jwtResponse := verifyJWT(authProvider, idToken)

			log.Printf("jwt response: %+v\n", *jwtResponse)

			if jwtResponse == nil {
				ctx = getAuthedContext(false, nil, nil, nil)
			} else {

				switch jwtResponse.Type {
				case ApiKey:
					{
						var entityType = ApiKey
						apiKey, _ := apiKeyRepo.FindById(jwtResponse.Value)
						ctx = getAuthedContext(apiKey != nil, &entityType, nil, apiKey)
					}
				case User:
					{
						var entityType = User
						user, _ := userRepo.FindByAuthProviderId(jwtResponse.Value)
						ctx = getAuthedContext(user != nil, &entityType, user, nil)
					}
				}
			}
		} else {
			ctx = getAuthedContext(false, nil, nil, nil)
		}

		fmt.Printf("\nUser: %+v\n", ctx.Value("user"))

		c.SetUserContext(ctx)

		return c.Next()
	}

}

func AuthRequired(c *fiber.Ctx) error {
	ctx := c.UserContext()

	isAuthed := ctx.Value("is_authenticated").(bool)
	user := ctx.Value("user").(*UserDomain.User) // TODO: is this bad to case?

	if !isAuthed {
		return c.JSON(errors.NewApiError(c, errors2.New("not authenticated"), 401, struct{}{}))
	}

	if user == nil {
		return errors.NewApiError(c, errors2.New("no user found"), 401, struct{}{})
	}

	return c.Next()
}
