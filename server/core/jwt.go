package core

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Value string `json:"value"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

var hmacSampleSecret = []byte("some bytes") // TODO:
var issuer = "diffme"

// one hour
var expirationMilliseconds int64 = 60 * 60 * 1000

func GenerateToken(value string, role string) (string, error) {
	// Create the Claims
	claims := Claims{
		value,
		role,
		jwt.StandardClaims{
			ExpiresAt: expirationMilliseconds,
			Issuer:    issuer,
		},
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}

func ParseToken(tokenString string) ([]byte, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return []byte(claims.Value), nil
	} else {
		return nil, fmt.Errorf("token is not valid")
	}
}
