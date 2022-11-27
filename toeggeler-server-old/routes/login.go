package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steinm91/toeggeler/security"
)

type LoginRequest struct {
	Player   string `json:"player"`
	Password string `json:"password"`
}

func (controller Controller) Login(c *gin.Context) {
	var loginRequest LoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		return
	}

	player, err := controller.repo.GetPlayerByName(loginRequest.Player)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unexpected error"})
	}

	if security.CompareHashAndPassword(player.Password, loginRequest.Password) {
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "wrong credentials"})
	}
}
