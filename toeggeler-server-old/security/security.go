package security

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Println("Could not generate hash and salt from password")
		return "", err
	}
	return string(hashedPassword), nil
}

func CompareHashAndPassword(hash string, password string) bool {
	byteHash := []byte(hash)
	bytePassword := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
