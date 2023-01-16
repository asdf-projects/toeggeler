package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/steinm91/toeggeler/toeggeler-server/eval"
	"github.com/steinm91/toeggeler/toeggeler-server/models"
)

type StatsController struct {
	GameService *models.GameService
	UserService *models.UserService
	EvalEngine  *eval.EvalEngine
}

type Stats struct {
	PlayerId int64 `json:"playerId"`
	Rating   int   `json:"rating"`
	Wins     int   `json:"wins"`
	Losses   int   `json:"losses"`
	Goals    int   `json:"goals"`
	Foetelis int   `json:"foetelis"`
	OwnGoals int   `json:"ownGoals"`
}

func (ctrl StatsController) GetStats(c *gin.Context) {
	stats := getFakeStats()
	c.JSON(http.StatusOK, stats)
}

func (ctrl StatsController) GetStatsForPlayer(c *gin.Context) {
	stats := getFakeStats()

	id, _ := strconv.Atoi(c.Param(("id")))

	if (id - 1) > len(stats) {
		c.String(http.StatusNotFound, "Player not found")
		return
	}

	statsForPlayer := stats[id-1]
	c.JSON(http.StatusOK, statsForPlayer)
}

func getFakeStats() []Stats {
	p1 := Stats{PlayerId: 1, Rating: 344, Wins: 15, Losses: 2, Goals: 34, Foetelis: 13, OwnGoals: 2}
	p2 := Stats{PlayerId: 2, Rating: 494, Wins: 4, Losses: 11, Goals: 12, Foetelis: 0, OwnGoals: 3}
	p3 := Stats{PlayerId: 3, Rating: 248, Wins: 7, Losses: 7, Goals: 11, Foetelis: 4, OwnGoals: 0}
	p4 := Stats{PlayerId: 4, Rating: 465, Wins: 10, Losses: 5, Goals: 4, Foetelis: 1, OwnGoals: 0}

	return []Stats{p1, p2, p3, p4}
}
