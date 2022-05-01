package jwt_helper

import (
	"api/constants"
	errorHelper "api/helpers/error-helper"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
)

type JWTSign interface {
	SignClaim() (string, error)
}

type JWTError struct {
	Message string
	Err error
};

func (JWTErr JWTError) Error() string {
	return JWTErr.Message
}

func (JWTErr JWTError) Unwrap() error {
	return JWTErr.Err
}

type JWTClaim jwt.MapClaims

// SignClaim use os.Getenv("JWT_SECRET") to get the secret key
func (claims JWTClaim) SignClaim(secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims(claims))

	return token.SignedString(secret)
}

type JWTToken string

func (jwtToken string) ParseClaim() (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token jwt.Token) (interface{}, error) {

	})

	return token, err
}