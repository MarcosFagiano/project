package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("superduper secure password prr")
var RefreshSecret = []byte("superduper secure password prr 2")

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func CheckPassword(password string, hash string) bool {
	return HashPassword(password) == hash
}

func GenerateJWT(email string, admin bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"admin": admin,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString(jwtKey)
}
func ParseJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
