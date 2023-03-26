package auth

import (
	"errors"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
)

var (
	errInvalidToken = errors.New("invalid token")
	errParseToken   = errors.New("failed to parse token")
)

type CustomClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func EncodeToken(claims CustomClaims) (string, error) {
	claims.StandardClaims = jwt.StandardClaims{
		Issuer: "server",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		log.Printf("[auth.EncodeToken] failed to encode token for username %s: %s", claims.Username, err.Error())
		return "", err
	}

	return signed, err
}

func VerifyToken(header string) (CustomClaims, error) {
	token, err := jwt.ParseWithClaims(header, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Printf("[auth.VerifyToken] failed to verify token: %s\n", err.Error())
		return CustomClaims{}, err
	}
	if !token.Valid {
		return CustomClaims{}, errInvalidToken
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		log.Printf("[auth.VerifyToken] %s", err.Error())
		return CustomClaims{}, errParseToken
	}

	return *claims, nil
}
