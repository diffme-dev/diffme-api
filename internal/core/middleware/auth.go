package middleware

import (
	"diffme.dev/diffme-api/internal/core/interfaces"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Context struct {
	ctx   *fiber.Ctx
	token string
}

func AuthMiddleware(
	authProvider interfaces.AuthProvider,
) fiber.Handler {
	return func(c *fiber.Ctx) error {

		rcc := &Context{
			ctx:   c,
			token: "",
		}

		header := c.Request().Header.Peek("Authorization")

		idToken := strings.TrimSpace(strings.Replace(string(header), "Bearer", "", 1))

		uid, err := authProvider.VerifyToken(idToken)

		if err != nil {
			return c.JSON(err.Error())
		}

		println(uid)

		// Set context to token
		rcc.token = ""

		return c.Next()
	}

}

func RequireAuth() bool {
	return false
}
