package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := validateToken(c, secretKey)

		if err != nil {
			c.String(http.StatusUnauthorized, "Authentication required")
			log.Println(err)
			c.Abort()
			return
		}

		c.Next()
	}
}

func validateToken(c *gin.Context, secretKey string) error {
	token, err := extractToken(c, secretKey)

	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)

		log.Print("Token verified for user " + claims["username"].(string))
		return nil
	}

	return errors.New("Invalid token provided")
}

func extractToken(c *gin.Context, secretKey string) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)

	log.Printf("Token: %s", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return []byte(secretKey), nil
	})

	return token, err
}

func getTokenFromRequest(c *gin.Context) string {
	log.Println(c.Request.Header)
	bearerToken := c.Request.Header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}

	return ""
}
