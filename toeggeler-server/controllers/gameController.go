package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"github.com/steinm91/toeggeler/toeggeler-server/models"
)

type GameController struct {
	GameService *models.GameService
}

/*
func (gameCtrl GameController) GetGamesPlayed(c *gin.Context) {
	playerId := c.Query("player")
	if playerId == "" {
		c.String(http.StatusBadRequest, "Please provide a player id using the query parameter 'player'")
		return
	}

	id, _ := strconv.ParseInt(playerId, 10, 64)
	games, _ := gameCtrl.GameService.GetGamesPlayedForPlayer(id)

	c.JSON(http.StatusOK, games)
}
*/

func (gameCtrl GameController) SubmitGame(c *gin.Context) {
	var err error
	var gameId ksuid.KSUID
	var gameEvents []models.GameEvent

	if err := c.BindJSON(&gameEvents); err != nil {
		c.String(http.StatusBadRequest, "Invalid payload")
	}

	gameId, err = ksuid.NewRandom()
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not submit game data")
		return
	}

	for i, event := range gameEvents {
		err = event.IsValid()
		gameEvents[i].GameId = gameId.String()
		if err != nil {
			break
		}
	}

	if err != nil {
		c.String(http.StatusBadRequest, "Invalid event provided")
		return
	}

	game, err := gameCtrl.GameService.InsertGameEvents(gameId.String(), &gameEvents)
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not submit game data")
		return
	}

	c.JSON(http.StatusOK, game)
}

func (gameCtrl GameController) ClearGames(c *gin.Context) {
	err := gameCtrl.GameService.ClearGamesTable()

	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Could not clear games table")
	}
}

func (gameCtrl GameController) ClearPlayerStats(c *gin.Context) {
	err := gameCtrl.GameService.ClearGamePlayerStats()

	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Could not clear player stats table")
	}
}
