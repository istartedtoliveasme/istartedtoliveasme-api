package jwt_helper

import (
	"api/constants"
	errorHelper "api/helpers/error-helper"
	"fmt"
	"github.com/golang-jwt/jwt"
)

type JWTSign interface {
	SignClaim() (string, error)
}

type IJWTError interface {
	errorHelper.CustomError
}

type JWTError struct {
	Message string
	Err     error
}

func (JWTErr JWTError) Error() string {
	return JWTErr.Message
}

func (JWTErr JWTError) Unwrap() error {
	return JWTErr.Err
}

type JWTClaim jwt.MapClaims

// SignClaim use os.Getenv("JWT_SECRET") to get the secret key
func (claims JWTClaim) SignClaim(secret []byte) (string, error) {
	// TODO :: use SigningString method that has a parameter of key
	return jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims(claims)).SigningString()
}

type JWTToken string

func (jwtToken JWTToken) ParseClaim() (*jwt.Token, error) {
	return jwt.Parse(string(jwtToken), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(constants.FailedParseClaim+": %v", token.Header[HeaderAlg])
		}

		return jwtToken, nil
	})
}

const (
	HeaderAlg = "alg"
)
