package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

const salt = "proyecto2025"

// HashPasswordSHA256 genera el hash SHA256 + sal
func HashPasswordSHA256(password string) string {
	hash := sha256.New()
	hash.Write([]byte(salt + password))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

// CheckPasswordSHA256 compara un password con su hash
func CheckPasswordSHA256(password, hashed string) bool {
	return HashPasswordSHA256(password) == hashed
}
