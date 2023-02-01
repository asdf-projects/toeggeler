package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/steinm91/toeggeler/toeggeler-server/models"
	"golang.org/x/crypto/bcrypt"
)

type SecurityController struct {
	SecretKey   string
	UserService *models.UserService
}

type AuthenRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthenResponse struct {
	Token string `json:"token"`
}

// Authenticate godoc
// @Summary      Authenticate by username and password
// @Description  Authenticate (Passwords for all available users is "1234")
// @Tags         Authenticate
// @Accept       json
// @Produce      json
// @Param        Credentials body AuthenRequest true  "User credentials"
// @Success      200 {object} AuthenResponse
// @Router       /authenticate [post]
func (s SecurityController) Authenticate(c *gin.Context) {
	var authenRequest AuthenRequest

	if err := c.BindJSON(&authenRequest); err != nil {
		c.String(http.StatusBadRequest, "Invalid payload")
		return
	}

	hashedPassword, err := s.UserService.GetUserPassword(authenRequest.Username)
	if err != nil {
		log.Println(err)
		c.String(http.StatusUnauthorized, "Could not authenticate")
		return
	}

	if compareHashAndPassword(*hashedPassword, authenRequest.Password) {
		token, err := s.generateJWT(authenRequest.Username)
		if err != nil {
			c.String(http.StatusInternalServerError, "Could not generate JWT token")
			return
		}

		authenResponse := AuthenResponse{Token: token}

		c.JSON(http.StatusOK, authenResponse)
	} else {
		log.Println("blalba")
		c.String(http.StatusUnauthorized, "Could not authenticate")
	}
}

func compareHashAndPassword(hash string, password string) bool {
	byteHash := []byte(hash)
	bytePassword := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (s SecurityController) generateJWT(username string) (string, error) {
	secretKey := []byte(s.SecretKey)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		log.Printf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
