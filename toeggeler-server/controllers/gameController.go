package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"github.com/steinm91/toeggeler/toeggeler-server/eval"
	"github.com/steinm91/toeggeler/toeggeler-server/models"
)

type GameController struct {
	GameService *models.GameService
	EvalEngine  *eval.EvalEngine
}

type Score struct {
	Team1 int `json:"team1"`
	Team2 int `json:"team2"`
}

type Team struct {
	Offense int64 `json:"offense"`
	Defense int64 `json:"defense"`
}

type Game struct {
	GameId    string `json:"gameId"`
	GameStart int64  `json:"gameStart"`
	GameEnd   int64  `json:"gameEnd"`
	Team1     Team   `json:"team1"`
	Team2     Team   `json:"team2"`
	Score     Score  `json:"score"`
}

// SubmitGame godoc
// @Summary      Submit a game
// @Description  Submit a game by listing every event.
// @Description  Available events: <ul><li>GAME_START</li><li>GOAL</li><li>OWN_GOAL</li><li>FOETELI</li><li>GAME_END</li></ul>
// @Tags		 Games
// @Security	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        user body []models.GameEvent true "Game events"
// @Success      200  {object}	Game
// @Failure		 404
// @Router       /games [post]
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

func (gameCtrl GameController) GetGamesPlayed(c *gin.Context) {
	c.JSON(http.StatusOK, gameCtrl.EvalEngine.GetGames())
}

/*
func (gameCtrl GameController) GetGamesPlayed(c *gin.Context) {
	gamesPlayed, err := gameCtrl.GameService.GetGamesPlayed()

	if err != nil {
		c.String(http.StatusInternalServerError, "Could not retrieve games")
	}

	games := []Game{}
	for _, g := range *gamesPlayed {
		game := Game{
			GameId:    g.GameId,
			GameStart: g.GameStart,
			GameEnd:   g.GameEnd,
			Team1: Team{
				Offense: g.Offense1,
				Defense: g.Defense1,
			},
			Team2: Team{
				Offense: g.Offense2,
				Defense: g.Defense2,
			},
			Score: Score{
				Team1: g.Team1Goals,
				Team2: g.Team2Goals,
			},
		}

		games = append(games, game)
	}

	c.JSON(http.StatusOK, games)
}
*/

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
