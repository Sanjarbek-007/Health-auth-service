package tokens

import (
	"auth-service/config"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func GenerateRefreshToken(userID string) (string, error) {
	token := *jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	newToken, err := token.SignedString([]byte(config.Load().Server.REFRESH_TOKEN_KEY))
	if err != nil {
		log.Println(err)
		return "", errors.Wrap(err, "failed to generate refresh token")
	}

	return newToken, nil
}

func ValidateRefreshToken(tokenStr string) (bool, error) {
	_, err := ExtractRefreshClaims(tokenStr)
	if err != nil {
		return false, errors.Wrap(err, "validation failure")
	}

	return true, nil
}

func ExtractRefreshClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Load().Server.REFRESH_TOKEN_KEY), nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse refresh token")
	}

	if !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func ExtractRefreshUserID(tokenStr string) (string, error) {
	claims, err := ExtractRefreshClaims(tokenStr)
	if err != nil {
		return "", errors.Wrap(err, "failed to extract refresh token's claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("invalid token claims: user id not found")
	}

	return userID, nil
}
