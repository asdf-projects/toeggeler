package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steinm91/toeggeler/data"
	"github.com/steinm91/toeggeler/mock"
)

func (controller Controller) GetGame(c *gin.Context) {
	game := mock.CreateGame()
	c.JSON(http.StatusOK, game)
}

func (controller Controller) ParseGameEvent(c *gin.Context) {
	var gameEvent data.GameEvent

	if err := c.BindJSON(&gameEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "all good"})
}
