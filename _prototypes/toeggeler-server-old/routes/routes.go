package routes

import (
	"database/sql"
	"log"

	"github.com/steinm91/toeggeler/data"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	db   *sql.DB
	repo *data.SQLiteRepository
}

func New(db *sql.DB, repo *data.SQLiteRepository) Controller {
	return Controller{
		db:   db,
		repo: repo,
	}
}

func (controller Controller) HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Println("Could not generate hash and salt from password")
		return "", err
	}
	return string(hashedPassword), nil
}

func (controller Controller) CompareHashAndPassword(hash string, password string) bool {
	byteHash := []byte(hash)
	bytePassword := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
