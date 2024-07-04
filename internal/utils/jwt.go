package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint) (string, error) {
	var claims = jwt.MapClaims{}
	claims["id"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	var process = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := process.SignedString([]byte("passkeyJWT"))
	if err != nil {
		return "", err
	}
	return result, nil
}

func DecodeToken(token *jwt.Token) (uint, error) {
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return 0, errors.New("invalid token claims")
    }

    userIDFloat, ok := claims["id"].(float64)
    if !ok {
        return 0, errors.New("invalid user ID type")
    }

    userID := uint(userIDFloat)
    return userID, nil
}