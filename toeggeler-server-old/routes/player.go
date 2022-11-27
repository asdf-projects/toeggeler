package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steinm91/toeggeler/data"
	"github.com/steinm91/toeggeler/security"
)

type NewPlayerRequest struct {
	Name     string `json:"name" binding:"required"`
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (controller Controller) CreatePlayer(c *gin.Context) {
	var newPlayer NewPlayerRequest

	if err := c.BindJSON(&newPlayer); err != nil {
		return
	}

	hashedPassword, err := security.HashPassword(newPlayer.Password)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	player := data.Player{
		Name:     newPlayer.Name,
		Mail:     newPlayer.Mail,
		Password: hashedPassword,
	}

	_, err = controller.repo.CreatePlayer(player)
	if err != nil {
		if errors.Is(err, data.ErrDuplicate) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "player already exists"})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unexpected error"})
	}
}

func (controller Controller) GetAllPlayers(c *gin.Context) {
	players, err := controller.repo.GetAllPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unexpected error"})
		return
	}
	c.JSON(http.StatusOK, players)
}

func (controller Controller) GetPlayerByName(c *gin.Context) {
	name := c.Param("name")

	player, err := controller.repo.GetPlayerByName(name)
	if err != nil {
		if errors.Is(err, data.ErrNotExists) {
			c.JSON(http.StatusNotFound, gin.H{"message": "player not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unexpected error"})
	}
	c.IndentedJSON(http.StatusOK, player)
}
