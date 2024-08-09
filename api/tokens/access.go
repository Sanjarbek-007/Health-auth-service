package tokens

import (
	"auth-service/config"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func GenerateAccessToken(id, role string) (string, error) {
	token := *jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["role"] = role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()

	newToken, err := token.SignedString([]byte(config.Load().Server.ACCESS_TOKEN_KEY))

	if err != nil {
		log.Println(err)
		return "", errors.Wrap(err, "failed to generate access token")
	}

	return newToken, nil
}
